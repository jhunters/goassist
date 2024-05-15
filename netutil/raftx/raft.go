// raftx is a utility package for easy to use raft library.
package raftx

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"github.com/Jille/raft-grpc-leader-rpc/leaderhealth"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/Jille/raftadmin"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	RAFT_DATA_DIR = "data"

	Logs_File = "logs.dat"

	Stable_File = "stable.dat"

	LOCAL_HOST = "localhost"
)

// Node 实例节点信息
type Node struct {
	Id   string // 节点id
	Addr string // 节点地址
}

// RaftX 实例
type RaftX struct {
	// Raft 配置信息
	config *raft.Config

	// 本地IP地址
	localip string

	// 本地端口号
	localport string

	// 数据保存目录, 包括wal日志， snap快照， raft配置文件
	dataDir string

	// 集群其它节点列表
	peers []*Node

	// 标记RaftX是否已启动的原子布尔值
	started atomic.Bool

	// 是否启用健康服务
	healthService bool

	// 健康服务名称
	healthServiceName string

	// 是否启用Raft管理功能
	raftadmin bool

	// 是否启用反射服务
	reflectionService bool

	// 有限状态机（FSM）
	fsm raft.FSM

	// Raft 实例指针
	r *raft.Raft

	// 是否为引导Raft集群的标记
	raftBootstrap bool
}

// NewRaftX 函数用于创建一个新的RaftX实例。
//
// 参数：
// node: 指向Node类型的指针，表示节点对象。
// raftBootstrap: bool类型，表示是否开启Raft的引导模式。
// fsm: raft.FSM类型，表示状态机对象。
//
// 返回值：
// *RaftX: 指向RaftX类型的指针，表示创建的RaftX实例。
// error: 表示在创建RaftX实例时可能出现的错误。
func NewRaftX(node *Node, raftBootstrap bool, fsm raft.FSM) (*RaftX, error) {
	c := raft.DefaultConfig()
	return NewRaftXWithConfig(c, node, raftBootstrap, fsm)
}

// NewRaftXWithConfig 根据提供的 Raft 配置、节点信息、是否引导 Raft 集群和有限状态机（FSM）创建一个新的 RaftX 实例
//
// 参数：
// c *raft.Config - Raft 配置信息
// node *Node - 节点信息
// raftBootstrap bool - 是否引导 Raft 集群. 如果已存在 Raft 集群配置目录，则不会生效该设置，使用配置目录中的设置
// fsm raft.FSM - raft.FSM 接口实现，支持BatchFSM接口
//
// 返回值：
// *RaftX - 创建的 RaftX 实例
// error - 如果创建过程中出现错误，则返回错误信息
func NewRaftXWithConfig(c *raft.Config, node *Node, raftBootstrap bool, fsm raft.FSM) (*RaftX, error) {
	if c == nil {
		return nil, fmt.Errorf("raft config is nil")
	}

	if node == nil {
		return nil, fmt.Errorf("node is nil")
	}

	if fsm == nil {
		return nil, fmt.Errorf("raft fsm is nil")
	}

	r := &RaftX{}
	r.config = c
	r.fsm = fsm
	r.raftBootstrap = raftBootstrap
	c.LocalID = raft.ServerID(node.Id)

	host, port, err := net.SplitHostPort(node.Addr)
	if err != nil {
		return nil, err
	}

	if host == "" || host == LOCAL_HOST {
		host, err = os.Hostname()
		if err != nil {
			return nil, err
		}
	}

	r.localip = host
	r.localport = port
	r.dataDir = RAFT_DATA_DIR

	r.peers = make([]*Node, 0)

	return r, nil

}

func (r *RaftX) SetDataDir(dir string) error {
	if r == nil || dir == "" {
		return fmt.Errorf("raftx is nil or data dir is empty")
	}
	r.dataDir = dir
	return nil
}

func (r *RaftX) AddPeer(node *Node) error {
	if r.started.Load() {
		return fmt.Errorf("can't add peer while raft is started")
	}
	if node == nil {
		return fmt.Errorf("node is nil")
	}
	r.peers = append(r.peers, node)
	return nil
}

func (r *RaftX) AddPeers(nodes []*Node) error {
	for _, node := range nodes {
		err := r.AddPeer(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RaftX) SetHealthService(name string) error {
	if r.started.Load() {
		return fmt.Errorf("can't set health service while raft is started")
	}
	r.healthService = true
	r.healthServiceName = name
	return nil
}

func (r *RaftX) EnableRaftAdmin() {
	r.raftadmin = true
}

func (r *RaftX) EnableRefelctionService() {
	r.reflectionService = true
}

// Start 启动 RaftX 实例，并返回一个 raft.Raft 实例
func (r *RaftX) Start(s *grpc.Server, fn func(*raft.Raft)) error {
	if r == nil || r.config == nil {
		return fmt.Errorf("raftx is nil or config is nil")
	}

	if s == nil {
		return fmt.Errorf("grpc server is nil")
	}

	if !r.started.CompareAndSwap(false, true) {
		return fmt.Errorf("raftx is already started")
	}

	// 数据存储目录
	baseDir := filepath.Join(r.dataDir, string(r.config.LocalID))
	os.MkdirAll(baseDir, os.ModeDir)

	// 使用 boltdb进行wal日志存储
	// 1 存储 配置信息 rafe.Configuration 包含集群信息 LogType=LogConfiguration
	//
	wal, err := boltdb.NewBoltStore(filepath.Join(baseDir, "logs.dat"))
	if err != nil {
		return fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, Logs_File), err)
	}

	// 使用 boltdb进行 stable 存储 内容包含 最后的投票节点信息LastVoteCand，CurrentTerm LastVoteTerm
	sdb, err := boltdb.NewBoltStore(filepath.Join(baseDir, "stable.dat"))
	if err != nil {
		return fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, Stable_File), err)
	}

	// snap 数据目录
	// mss := raft.NewInmemSnapshotStore() // 仅使用内存的snapshot
	fss, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		return fmt.Errorf(`raft.NewFileSnapshotStore(%q, ...): %v`, baseDir, err)
	}

	// 创建 grpc 传输 transport.Manager 对象， 实现了 raft 接口同步协议实现
	// insecure.NewCredentials()  <= grpc.WithInsecure()
	// grpc.WithTransportCredentials(insecure.NewCredentials())
	tm := transport.New(raft.ServerAddress(r.GetAddress()), []grpc.DialOption{grpc.WithInsecure()})

	// fsm 需要实现日志同步的接口
	r.r, err = raft.NewRaft(r.config, r.fsm, wal, sdb, fss, tm.Transport())
	if err != nil {
		return fmt.Errorf("raft.NewRaft: %v", err)
	}

	// 使用grpc作为 raft 通讯服务
	tm.Register(s)

	if r.healthService {
		// 注册健康检查服务
		services := []string{r.healthServiceName}
		// 创建  建康心跳服务， 会发布 CacheManager 和 quis.RaftLeader 服务， 由grpc在负载时检测
		leaderhealth.Setup(r.r, s, services)
	}

	if r.raftadmin {
		// 注册 raft管理服务， 工具可用 go get github.com/Jille/raftadmin 安装后，提供命令行功能方式管理
		raftadmin.Register(s, r.r)
	}

	if r.reflectionService {
		// 反射服务 Register registers the server reflection service on the given gRPC server
		// 在 Server 端添加后可以通过该服务获取所有服务的信息，包括服务定义，方法，属性等；
		reflection.Register(s)
	}

	// 判断是否作为集群bootstrap节点启动，  配置 Servers 添加 ID和Address
	if r.raftBootstrap {
		err = r.startWithBootstrapIfNeed(wal)
		if err != nil {
			return err
		}
	}

	// call back FSM 接口
	fn(r.r)

	// 启动网络服务监听
	sock, err := net.Listen("tcp", fmt.Sprintf(":%s", r.localport) /* r.getAddress() */)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	// 启动rpc服务， 绑定socket监听
	if err := s.Serve(sock); err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return nil
}

// getAddress 返回 RaftX 实例绑定的地址信息
// 返回一个格式为 "host:port" 的字符串
func (r *RaftX) GetAddress() string {
	return fmt.Sprintf("%s:%s", r.localip, r.localport)
}

// startWithBootstrapIfNeed 是一个方法，用于检查是否需要对 RaftX 进行引导启动
// 如果日志已经存在，则不进行引导启动，直接返回 nil
// 如果日志不存在，则进行引导启动，并返回可能出现的错误
//
// 参数：
//
//	r *RaftX：RaftX 实例指针
//	rt *raft.Raft：Raft 实例指针
//	wal *boltdb.BoltStore：BoltDB 存储实例指针
//
// 返回值：
//
//	error：可能出现的错误，如果引导启动成功，则返回 nil
func (r *RaftX) startWithBootstrapIfNeed(wal *boltdb.BoltStore) error {
	// check if log already exist, cuz we can't bootstrap twice
	if id, err := wal.LastIndex(); id > 0 && err == nil {
		return nil
	}

	// 配置 Servers 添加 ID和Address
	cfg := raft.Configuration{
		Servers: []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       r.config.LocalID,
				Address:  raft.ServerAddress(r.GetAddress()),
			},
		},
	}

	// 添加peer信息
	processPeers(&cfg, r.peers)

	// 作为集群bootstrap节点的模式设置, 只有第一次启动生效，一旦集群信息写入db后，就不能再调用BootstrapCluster方法
	f := r.r.BootstrapCluster(cfg)
	if err := f.Error(); err != nil {
		e := fmt.Errorf("raft.Raft.BootstrapCluster: %v", err)
		return e
	}

	return nil
}

// processPeers 函数用于处理raft节点配置，将节点信息添加到raft配置中
//
// 参数：
// cfg *raft.Configuration - raft配置指针
// peers []*Node - 待处理的节点列表
//
// 返回值：
// 无
func processPeers(cfg *raft.Configuration, peers []*Node) {
	for _, n := range peers {

		host, _ := os.Hostname()
		addr := strings.Replace(n.Addr, LOCAL_HOST, host, 1)

		cfg.Servers = append(cfg.Servers, raft.Server{
			ID:       raft.ServerID(n.Id),
			Address:  raft.ServerAddress(addr),
			Suffrage: raft.Voter,
		})
	}

}
