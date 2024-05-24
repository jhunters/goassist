package example

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/maputil"
	"github.com/jhunters/goassist/netutil/raftx/example/proto"

	"github.com/hashicorp/raft"
	pb "google.golang.org/protobuf/proto"
)

const (
	PUT_C    = "__PUT__"
	DEL_C    = "__DEL__"
	CLR_C    = "__CLEAR__"
	C_offset = len(PUT_C)
)

// to implement raft.FSM 接口
type cacheTracker struct {
	mtx   sync.RWMutex
	cache *proto.CacheMgr
	addr  string
	id    string
}

var _ raft.BatchingFSM = &cacheTracker{} // check cacheTracker struct if implements raft.BatchingFSM interface

func (f *cacheTracker) Apply(l *raft.Log) interface{} {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	return f.processLog(l)
}

func (f *cacheTracker) processLog(l *raft.Log) interface{} {
	data := l.Data
	action := string(l.Extensions)

	if l.Extensions == nil { //server config append entry log

		cfg := base.SafetyFunc(data, func(b []byte) raft.Configuration {
			cfg := raft.DecodeConfiguration(data)
			return cfg
		})

		for _, s := range cfg.Servers {
			p := fmt.Sprintf("[%s=%s] server sync log. address=%s, id=%s, suffrage=%d ", f.id, f.addr, s.Address, s.ID, s.Suffrage)
			PrintlnLog(p)
		}
		return nil
	}

	sdata := ""
	if strings.EqualFold(action, PUT_C) {
		req := &proto.PutRequest{}
		pb.Unmarshal(data, req)
		f.cache.Cache[string(req.Key)] = req.Value
		sdata = fmt.Sprintf("key=%s, value=%v", req.Key, string(req.Value))
	} else if strings.EqualFold(action, DEL_C) {
		req := &proto.GetRequest{}
		pb.Unmarshal(data, req)
		value, exist := f.cache.Cache[string(req.Key)]
		if exist {
			delete(f.cache.Cache, string(req.Key))
		}
		sdata = fmt.Sprintf("key=%s", req.Key)
		// println log
		p := fmt.Sprintf("[%s=%s] action=%s, data=%v, index=%d, term=%d", f.id, f.addr, action, sdata, l.Index, l.Term)
		PrintlnLog(p)
		return value

	} else if strings.EqualFold(action, CLR_C) {
		f.cache.Cache = make(map[string][]byte)
	}

	// println log
	p := fmt.Sprintf("[%s=%s] action=%s, data=%v, index=%d, term=%d", f.id, f.addr, action, sdata, l.Index, l.Term)
	PrintlnLog(p)

	return len(f.cache.Cache)
}

func (f *cacheTracker) ApplyBatch(logs []*raft.Log) []interface{} {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	ret := make([]interface{}, 0)
	for _, l := range logs {
		ret = append(ret, f.processLog(l))
	}

	return ret
}

func (f *cacheTracker) Snapshot() (raft.FSMSnapshot, error) {
	// Make sure that any future calls to f.Apply() don't change the snapshot.
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	data, err := pb.Marshal(f.cache)
	return &snapshot{data: data, id: f.id, addr: f.addr}, err
}

func (f *cacheTracker) Restore(r io.ReadCloser) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	f.mtx.Lock()
	defer f.mtx.Unlock()

	err = pb.Unmarshal(b, f.cache)

	p := fmt.Sprintf("[%s=%s] Restore cache from snapshot, size=%d", f.id, f.addr, len(f.cache.Cache))
	PrintlnLog(p)
	return err
}

type snapshot struct {
	data []byte
	addr string
	id   string
}

var _ raft.FSMSnapshot = &snapshot{} // check snapshot struct if implements raft.FSMSnapshot interface

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	_, err := sink.Write(s.data)
	if err != nil {
		sink.Cancel()
		return fmt.Errorf("sink.Write(): %v", err)
	}
	p := fmt.Sprintf("[%s=%s]Snap short persist to=%s", s.id, s.addr, sink.ID())
	PrintlnLog(p)
	return sink.Close()
}

func (s *snapshot) Release() {
}

type CacheSvrInterface struct {
	proto.UnimplementedCacheManagerServer
	raft *raft.Raft
	wt   *cacheTracker
}

func (q *CacheSvrInterface) Put(c context.Context, req *proto.PutRequest) (*proto.PutResponse, error) {
	data, err := pb.Marshal(req)
	if err != nil {
		return &proto.PutResponse{Error: err.Error()}, err
	}

	// send raft sync
	log := raft.Log{Data: data, Extensions: []byte(PUT_C)}
	res := q.raft.ApplyLog(log, 5*time.Second)

	if res.Error() != nil {
		fmt.Println(res.Error())
		return &proto.PutResponse{Error: res.Error().Error()}, nil
	}

	fmt.Println(res.Index(), res.Response())
	v := res.Response().(int)
	return &proto.PutResponse{Size: int32(v)}, nil
}

func (q *CacheSvrInterface) Get(c context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {

	key := req.Key
	value, exist := q.wt.cache.Cache[string(key)]
	if exist {
		return &proto.GetResponse{Value: value}, nil
	}
	err := fmt.Errorf("key %s not found", string(key))
	return &proto.GetResponse{Error: err.Error()}, err
}

func (q *CacheSvrInterface) Del(c context.Context, req *proto.DelRequest) (*proto.DelResponse, error) {
	key := req.Key
	value, exist := q.wt.cache.Cache[string(key)]

	if exist {
		data, err := pb.Marshal(req)
		if err != nil {
			return &proto.DelResponse{Error: err.Error()}, err
		}

		// send raft sync
		log := raft.Log{Data: data, Extensions: []byte(DEL_C)}
		q.raft.ApplyLog(log, 5*time.Second)
		return &proto.DelResponse{Value: value}, nil
	}

	err := fmt.Errorf("key %s not found", string(key))
	return &proto.DelResponse{Error: err.Error()}, err
}

func (q *CacheSvrInterface) Clear(c context.Context, req *proto.ClearRequest) (*proto.ClearResponse, error) {
	// send raft sync
	log := raft.Log{Data: []byte{}, Extensions: []byte(CLR_C)}
	q.raft.ApplyLog(log, 5*time.Second)
	return &proto.ClearResponse{Error: ""}, nil

}

func (q *CacheSvrInterface) Query(ctx context.Context, in *proto.QueryRequest) (*proto.QueryResponse, error) {
	ret := &proto.QueryResponse{}
	ret.Cache = maputil.Clone(q.wt.cache.Cache)
	return ret, nil
}

func PrintlnLog(p string) {
	entryLog.Log(p)
}

func PrintKeys(mp map[string][]byte) {
	for k := range mp {
		fmt.Print(k, " ")
	}
	fmt.Println()
}
