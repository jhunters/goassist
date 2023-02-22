package netutil_test

import (
	"fmt"
	"net"
	"net/http"

	"github.com/jhunters/goassist/netutil"
)

func ExampleNewCustomListenerSelector() {

	// here we wants to process HTTP or BaiduRPC protocol with same port

	port := 1031
	var headsize uint8 = 4

	// create a new customize listener
	selector, err := netutil.NewCustomListenerSelector("tcp", "", port, headsize, netutil.Equal_Mode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// register baidurpc protocol
	rpcServerListener, err := selector.RegisterListener("PRPC") // baidu rpc protocol magic code "PRPC"
	if err != nil {
		fmt.Println(err)
		return
	}
	// use default as http protocol
	httpServerListener := selector.RegisterDefaultListener()
	// start port listen
	go selector.Serve()

	// bind listener to baidurpc
	fmt.Println(rpcServerListener)
	// serverMeta := baidurpc.ServerMeta{}
	// rpcServer := baidurpc.NewTpcServer(&serverMeta)
	// rpcServer.StartServer(rpcServerListener)

	// bind listener to http server
	http.Serve(httpServerListener, http.FileServer(http.Dir("./")))

}

func ExampleNewCustomListenerSelectorByAddr() {

	// here we wants to process HTTP or BaiduRPC protocol with same port

	server := ":1032"
	var headsize uint8 = 4

	// create a new customize listener
	selector, err := netutil.NewCustomListenerSelectorByAddr("tcp", server, headsize, netutil.Equal_Mode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// register baidurpc protocol
	rpcServerListener, err := selector.RegisterListener("PRPC") // baidu rpc protocol magic code "PRPC"
	if err != nil {
		fmt.Println(err)
		return
	}
	// use default as http protocol
	httpServerListener := selector.RegisterDefaultListener()
	// start port listen
	go selector.Serve()

	// bind listener to baidurpc
	fmt.Println(rpcServerListener)
	// serverMeta := baidurpc.ServerMeta{}
	// rpcServer := baidurpc.NewTpcServer(&serverMeta)
	// rpcServer.StartServer(rpcServerListener)

	// bind listener to http server
	http.Serve(httpServerListener, http.FileServer(http.Dir("./")))

}

func ExampleNewCustomListenerSelectorByListener() {

	// here we wants to process HTTP or BaiduRPC protocol with same port

	server := ":1033"
	listener, err := net.Listen("tcp", server)
	if err != nil {
		fmt.Println(err)
		return
	}
	var headsize uint8 = 4

	// create a new customize listener
	selector, err := netutil.NewCustomListenerSelectorByListener(listener, headsize, netutil.Equal_Mode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// register baidurpc protocol
	rpcServerListener, err := selector.RegisterListener("PRPC") // baidu rpc protocol magic code "PRPC"
	if err != nil {
		fmt.Println(err)
		return
	}
	// use default as http protocol
	httpServerListener := selector.RegisterDefaultListener()
	// start port listen
	go selector.Serve()

	// bind listener to baidurpc
	fmt.Println(rpcServerListener)
	// serverMeta := baidurpc.ServerMeta{}
	// rpcServer := baidurpc.NewTpcServer(&serverMeta)
	// rpcServer.StartServer(rpcServerListener)

	// bind listener to http server
	http.Serve(httpServerListener, http.FileServer(http.Dir("./")))

}
