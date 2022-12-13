package containerx_test

import (
	"strconv"
	"testing"

	"github.com/jhunters/goassist/containerx"
	. "github.com/smartystreets/goconvey/convey"
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

	Convey("TestHeap sort", t, func() {
		player := h.Pop()
		So(player.level, ShouldEqual, 1)
	})

}
