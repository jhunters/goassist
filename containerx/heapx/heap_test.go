package heapx_test

import (
	"strconv"
	"testing"

	"github.com/jhunters/goassist/containerx/heapx"
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
