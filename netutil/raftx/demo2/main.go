package main

import (
	"strconv"

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/netutil/raftx"
	"github.com/jhunters/goassist/netutil/raftx/example2"
)

func main() {

	size := 3

	nodes := make([]*raftx.Node, size)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = &raftx.Node{Id: "node" + strconv.Itoa(i), Addr: "localhost:5005" + strconv.Itoa(i)}
	}

	// start multiple raft instances and with peers
	for i := 0; i < len(nodes); i++ {
		Id := nodes[i].Id
		Addr := nodes[i].Addr

		cNodes := arrayutil.Clone(nodes)
		cNodes = arrayutil.RemoveIndex(cNodes, i)

		httpaddr := ":888" + strconv.Itoa(i+1)
		if i == size-1 {
			// start block way
			example2.StartRaft(httpaddr, Id, Addr, "data", true, cNodes)
		} else {
			go example2.StartRaft(httpaddr, Id, Addr, "data", false, cNodes)
		}

	}

}
