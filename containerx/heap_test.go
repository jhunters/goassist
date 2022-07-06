package containerx_test

import (
	"container/list"
	"fmt"
	"strconv"
	"testing"

	"github.com/jhunters/goassist/containerx"
)

func TestHeap(t *testing.T) {
	// Player 结构体，按级别来优化排序
	type Player struct {
		level int
		name  string
	}

	h := containerx.NewHeap([]Player{}, func(p1, p2 Player) int {
		return p1.level - p2.level // level小的先出
	})

	// 初始化 100个数据
	for i := 100; i > 0; i-- {
		h.Push(Player{i, "name" + strconv.Itoa(i)})
	}

	// 优先 pop level值 最小的
	for i := 0; i < 100; i++ {
		fmt.Println(h.Pop())
	}

}

func TestList(t *testing.T) {

	l := list.New()

	l.PushFront(1)
	l.PushFront(2)

	fmt.Println(l.Front())
	fmt.Println(l.Back())
}
