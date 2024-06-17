package example

import (
	"fmt"
	"log"

	"github.com/jhunters/goassist/logutil"
	"github.com/jhunters/goassist/netutil/raftx"
	"github.com/jhunters/goassist/netutil/raftx/example/proto"

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

func StartRaft(raftId, raftAddress, raftDir string, raftBootstrap bool, peers []*raftx.Node) {

	// 初始化 cacheTracker，实现 raft.FSM 接口
	// 实现 raft.FSM 接口
	wt := &cacheTracker{}
	wt.id = raftId
	wt.addr = raftAddress
	wt.cache = &proto.CacheMgr{Cache: make(map[string][]byte)}
	wt.changed = make(chan int, 1)

	// 创建一个 raftx.Node 实例
	node := &raftx.Node{Id: raftId, Addr: raftAddress}

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
		// 注册 CacheManagerServer 到 gRPC 服务，并启动一个 goroutine 来监听状态变化
		// 业务功能实现，实现业务CacheManager的接口实现， 也是用grpc来注册服务
		proto.RegisterCacheManagerServer(s, &CacheSvrInterface{
			wt:   wt,
			raft: r.Raft,
		})

		go watchStatus(r.Raft, wt, status)
	})

}

// 这段代码用于监控raft的状态和缓存大小。它首先设置一个状态变更时的回调函数并开始一个goroutine，用于监控raft的leader状态和host信息的更新。
// 同时，它开启了一个定时器ticker，每隔两秒触发一次，用于监控缓存大小的变更并输出到控制台
func watchStatus(r *raft.Raft, wt *cacheTracker, status *CacheStatus) {
	ch := make(chan raft.Observation, 1)
	// 状态变更时回调
	r.RegisterObserver(raft.NewObserver(ch, true, func(o *raft.Observation) bool {
		_, ok := o.Data.(raft.LeaderObservation) // 过滤条件，只处理 leader状态变更通知
		return ok
	}))
	go func() {
		for range ch {
			if r.State() != status.raftState {
				status.raftState = r.State()
				// fmt.Print("\x0c", "leader info:", leader)
				statusLog.Log(fmt.Sprintf("[%s=%s] leader: %v\n", wt.id, wt.addr, r.State() == raft.Leader))
			}

			if r.Leader() != status.raftHost {
				raftHost, id := r.LeaderWithID()
				status.raftHost = raftHost
				statusLog.Log(fmt.Sprintf("[%s=%s] leader info: [id=%s, host= %s]\n", wt.id, wt.addr, id, status.raftHost))
			}
		}
	}()

	for sz := range wt.changed {
		statusLog.Log(fmt.Sprintf("[%s=%s] cache size: %d\n", wt.id, wt.addr, sz))
	}
}
