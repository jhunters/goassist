/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-24 19:48:22
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	c := context.WithValue(context.TODO(), "hello", "world")
	fmt.Println(c.Value("hello"))

	// ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)

	// ClassicContextTimeoutDemo(ctx)

	// fmt.Println("Now before exit wait 5 seconds.")
	// time.Sleep(5 * time.Second)
	// cancel()
}

func watch(ctx context.Context, msg string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(msg, "收到信号，监控退出， time=", time.Now().Unix())
			return
		default:
			fmt.Println(msg, "goroutine 监控中, time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}

func ClassicContextTimeoutDemo(ctx context.Context) {

	syncRet := make(chan bool)

	// here do bussiness process here
	go func() {
		time.Sleep(3 * time.Second) // wait 3 seconds
		syncRet <- true
	}()

	select {
	case <-ctx.Done():
		fmt.Println("context canceled and return.")
		return
	case <-syncRet:
		fmt.Println("process finished.")
	}
}
