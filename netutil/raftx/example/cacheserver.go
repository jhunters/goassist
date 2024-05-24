package example

import (
	"fmt"
	"log"
	"time"

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

	// cache大小
	cacheSize int

	// 当前的实例的 host信息
	raftHost raft.ServerAddress
}

func StartRaft(raftId, raftAddress, raftDir string, raftBootstrap bool, peers []*raftx.Node) {

	// 实现 raft.FSM 接口
	wt := &cacheTracker{}
	wt.id = raftId
	wt.addr = raftAddress
	wt.cache = &proto.CacheMgr{Cache: make(map[string][]byte)}

	node := &raftx.Node{Id: raftId, Addr: raftAddress}
	rfx, err := raftx.NewRaftX(node, raftBootstrap, wt)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	rfx.SetDataDir(raftDir)
	rfx.EnableRaftAdmin()
	rfx.EnableRefelctionService()
	rfx.SetHealthService("CacheManager")

	// add peers
	rfx.AddPeers(peers...)

	// 创建grpc服务
	s := grpc.NewServer()

	status := &CacheStatus{}

	rfx.Start(s, func(r *raft.Raft) {
		// 业务功能实现，实现业务CacheManager的接口实现， 也是用grpc来注册服务
		proto.RegisterCacheManagerServer(s, &CacheSvrInterface{
			wt:   wt,
			raft: r,
		})

		go watchStatus(r, wt, status)
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
				statusLog.Write([]byte(fmt.Sprintf("[%s=%s] leader: %v\n", wt.id, wt.addr, r.State() == raft.Leader)))
			}

			if r.Leader() != status.raftHost {
				status.raftHost = r.Leader()
				statusLog.Write([]byte(fmt.Sprintf("[%s=%s] leader info: %s\n", wt.id, wt.addr, status.raftHost)))
			}
		}
	}()

	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if len(wt.cache.Cache) != status.cacheSize {
				status.cacheSize = len(wt.cache.Cache)
				statusLog.Write([]byte(fmt.Sprintf("[%s=%s] cache size: %d\n", wt.id, wt.addr, len(wt.cache.Cache))))
			}
		}
	}
}
