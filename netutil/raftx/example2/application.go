package example2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/maputil"

	"github.com/hashicorp/raft"
)

const (
	PUT_C    = "__PUT__"
	DEL_C    = "__DEL__"
	CLR_C    = "__CLEAR__"
	C_offset = len(PUT_C)
)

// to implement raft.FSM 接口
type cacheTracker struct {
	mtx      sync.RWMutex
	cache    map[string]string
	addr     string
	id       string
	changed  chan int
	httpaddr string
}

var _ raft.BatchingFSM = &cacheTracker{} // check cacheTracker struct if implements raft.BatchingFSM interface

func startHttpServer(addr string, wt *cacheTracker, r *raft.Raft) {
	router := gin.Default()
	// get cluster info list
	router.GET("/cache", func(c *gin.Context) {

		if r.State() != raft.Leader {
			c.String(http.StatusOK, "reject, current node is not leader!")

			return
		}

		action := c.Query("action")
		k := c.Query("key")
		v := c.Query("value")
		if strings.EqualFold(action, PUT_C) {
			if _, ok := wt.cache[k]; ok {
				c.String(http.StatusOK, "key already exists")
				return
			}
			syncData2Peers(k, v, PUT_C, r, time.Second)
		} else if strings.EqualFold(action, DEL_C) {
			if _, ok := wt.cache[k]; !ok {
				c.String(http.StatusOK, "key not exists")
				return
			}
			syncData2Peers(k, v, DEL_C, r, time.Second)
		} else if strings.EqualFold(action, CLR_C) {
			log := raft.Log{Data: []byte(""), Extensions: []byte(CLR_C)}
			r.ApplyLog(log, time.Second)
		}

		c.String(http.StatusOK, "ok")
	})

	router.GET("/query", func(c *gin.Context) {
		b, err := maputil.JsonEncode(wt.cache)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}

		c.String(http.StatusOK, string(b))
	})

	router.Run(addr)
}

// syncData2Peers 用于将数据同步到 Raft 集群中的其他节点
//
// 参数：
// k: 数据的键
// v: 数据的值
// action: 对数据进行的操作（例如"add"、"delete"等）
// r: Raft 集群的 Raft 实例指针
// timeout: 应用日志的超时时间
//
// 返回值：无
func syncData2Peers(k, v, action string, r *raft.Raft, timeout time.Duration) {
	mp := make(map[string]string)
	mp[k] = v
	data, _ := maputil.JsonEncode(mp)
	log := raft.Log{Data: data, Extensions: []byte(action)}
	r.ApplyLog(log, timeout)
}

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
			c := raft.DecodeConfiguration(data)
			return c
		})

		for _, s := range cfg.Servers {
			p := fmt.Sprintf("[%s=%s http server=%s] server sync log. address=%s, id=%s, suffrage=%d ", f.id, f.addr, f.httpaddr, s.Address, s.ID, s.Suffrage)
			PrintlnLog(p)
		}
		return nil
	}

	sdata := ""
	if strings.EqualFold(action, PUT_C) {
		buf := bytes.NewBuffer(data)
		jsondecode := json.NewDecoder(buf)
		mp := make(map[string]string)
		jsondecode.Decode(&mp)
		for k, v := range mp {
			f.cache[k] = v
			sdata = fmt.Sprintf("key=%s, value=%v", k, v)
		}

		f.changeCallback()
	} else if strings.EqualFold(action, DEL_C) {
		buf := bytes.NewBuffer(data)
		jsondecode := json.NewDecoder(buf)
		mp := make(map[string]string)
		jsondecode.Decode(&mp)
		for k, _ := range mp {
			sdata = fmt.Sprintf("key=%s, value=", k)
			delete(f.cache, k)
			// println log
			p := fmt.Sprintf("[%s=%s http server=%s] action=%s, data=%v, index=%d, term=%d", f.id, f.addr, f.httpaddr, action, sdata, l.Index, l.Term)
			PrintlnLog(p)
		}

		f.changeCallback()

	} else if strings.EqualFold(action, CLR_C) {
		f.cache = make(map[string]string)
		f.changeCallback()
	}

	// println log
	p := fmt.Sprintf("[%s=%s http server=%s] action=%s, data=%v, index=%d, term=%d", f.id, f.addr, f.httpaddr, action, sdata, l.Index, l.Term)
	PrintlnLog(p)

	return len(f.cache)
}

func (f *cacheTracker) changeCallback() {
	go func() {
		f.changed <- len(f.cache)
	}()
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

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(f.cache)

	return &snapshot{data: buf.Bytes(), id: f.id, addr: f.addr}, err
}

func (f *cacheTracker) Restore(r io.ReadCloser) error {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(b)
	json.NewDecoder(buf).Decode(&f.cache)

	p := fmt.Sprintf("[%s=%s] Restore cache from snapshot, size=%d", f.id, f.addr, len(f.cache))
	PrintlnLog(p)
	f.changeCallback()
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
		return fmt.Errorf("sink.Write(): %v", err)
	}
	p := fmt.Sprintf("[%s=%s]Snap short persist to=%s", s.id, s.addr, sink.ID())
	PrintlnLog(p)
	return nil
}

func (s *snapshot) Release() {
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
