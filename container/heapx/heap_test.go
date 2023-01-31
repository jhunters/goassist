package heapx_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jhunters/goassist/container/heapx"
	. "github.com/smartystreets/goconvey/convey"
)

// Player 结构体，按级别来优化排序
type Player struct {
	level int
	name  string
}

func TestHeap(t *testing.T) {

	h := heapx.NewHeap([]Player{}, func(p1, p2 Player) int {
		return p1.level - p2.level // level小的先出
	})

	// 初始化 100个数据
	for i := 100; i > 0; i-- {
		h.Push(Player{i, "name" + strconv.Itoa(i)})
	}

	Convey("TestHeap sort", t, func() {
		player := h.Pop()
		So(player.level, ShouldEqual, 1)
	})

}

func ExampleNewHeap() {
	h := heapx.NewHeap([]Player{}, func(p1, p2 Player) int {
		return p1.level - p2.level // level小的先出
	})

	// 初始化 100个数据
	for i := 100; i > 0; i-- {
		h.Push(Player{i, "name" + strconv.Itoa(i)})
	}

	player := h.Pop()
	fmt.Println(player.level)
	player = h.Pop()
	fmt.Println(player.level)

	// Output:
	// 1
	// 2
}

func TestHeapCopy(t *testing.T) {
	Convey("TestHeapCopy", t, func() {

		h := heapx.NewHeap([]Player{}, func(p1, p2 Player) int {
			return p1.level - p2.level // level小的先出
		})
		h.Push(Player{1, "matthew"})
		h.Push(Player{2, "matt"})

		h2 := h.Copy()

		player := h2.Pop()
		So(player, ShouldNotBeNil)
		So(player.level, ShouldEqual, 1)
		player = h2.Pop()
		So(player, ShouldNotBeNil)
	})
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	h := heapx.NewHeap(IntHeap{2, 1, 5}, func(p1, p2 int) int {
		return p1 - p2
	})

	h.Push(3)

	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	// Output:
	// 1 2 3 5
}
