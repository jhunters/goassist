package example2

import (
	"fmt"
	"log"

	"github.com/jhunters/goassist/logutil"
	"github.com/jhunters/goassist/netutil/raftx"

	"github.com/hashicorp/raft"
	"google.golang.org/grpc"
)

const (
	Logs_File = "logs.dat"

	Stable_File = "stable.dat"
)

var (
	entryLog  = logutil.CreateLogger("AppendEntryLog", logutil.GREEN)
	statusLog = logutil.CreateLogger("StatusLog", logutil.YELLOW)
)

type CacheStatus struct {
	// 当前的leader状态
	raftState raft.RaftState

	// 当前的实例的 host信息
	raftHost raft.ServerAddress
}

func StartRaft(addr, raftId, raftAddress, raftDir string, raftBootstrap bool, peers []*raftx.Node) {

	// 初始化 cacheTracker，实现 raft.FSM 接口
	// 实现 raft.FSM 接口
	wt := &cacheTracker{}
	wt.id = raftId
	wt.addr = raftAddress
	wt.cache = make(map[string]string)
	wt.changed = make(chan int, 1)
	wt.httpaddr = addr

	// 创建一个 raftx.Node 实例
	node := &raftx.Node{Id: raftId, Addr: raftAddress, Desc: addr}

	// 创建 raftx 实例
	rfx, err := raftx.NewRaftX(node, raftBootstrap, wt)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	// 设置数据文件目录
	rfx.SetDataDir(raftDir)

	// 启用 raftadmin 管理功能
	rfx.EnableRaftAdmin()

	// 启用 gRPC 反射调用服务
	rfx.EnableRefelctionService()

	// 设置健康检查服务，只在 leader 节点可用
	rfx.SetHealthService("CacheManager")

	// 添加 peers
	rfx.AddPeers(peers...) // 设置 peer id与addr信息

	// 创建 gRPC 服务
	s := grpc.NewServer()

	// 创建一个 CacheStatus 实例
	status := &CacheStatus{}

	rfx.Start(s, func(r *raftx.RaftWrapper) {
		// 启动http服务， 实现缓存的增删改查功能服务
		go startHttpServer(addr, wt, r.Raft)
		// 启动对缓存的数据变化监控
		go watchStatus(r, wt, status)
	})

}

// 这段代码用于监控raft的状态和缓存大小。它首先设置一个状态变更时的回调函数并开始一个goroutine，用于监控raft的leader状态和host信息的更新。
// 同时，它开启了一个定时器ticker，每隔两秒触发一次，用于监控缓存大小的变更并输出到控制台
func watchStatus(r *raftx.RaftWrapper, wt *cacheTracker, status *CacheStatus) {
	ch := make(chan raft.Observation, 1)
	// 状态变更时回调
	r.Raft.RegisterObserver(raft.NewObserver(ch, true, func(o *raft.Observation) bool {
		_, ok := o.Data.(raft.LeaderObservation) // 过滤条件，只处理 leader状态变更通知
		return ok
	}))
	go func() {
		for range ch {
			if r.Raft.State() != status.raftState {
				status.raftState = r.Raft.State()
				if r.Raft.State() == raft.Leader { // if leader
					statusLog.Log(fmt.Sprintf("[%s=%s] leader: %v http server address: %v\n", wt.id, wt.addr, r.Raft.State() == raft.Leader, r.LocalNodeInfo.Desc))
				}

			}

			if r.Raft.Leader() != status.raftHost {
				raftHost, id := r.Raft.LeaderWithID()
				status.raftHost = raftHost
				statusLog.Log(fmt.Sprintf("[%s=%s http server=%s] leader info: [id=%s, host= %s]\n", wt.id, wt.addr, r.LocalNodeInfo.Desc, id, status.raftHost))
			}
		}
	}()

	for sz := range wt.changed {
		statusLog.Log(fmt.Sprintf("[%s=%s http server=%s] cache size: %d\n", wt.id, wt.addr, r.LocalNodeInfo.Desc, sz))
	}
}
