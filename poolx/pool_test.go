package poolx_test

import (
	"testing"

	"github.com/jhunters/goassist/poolx"

	. "github.com/smartystreets/goconvey/convey"
)

type Pojo struct {
	name string
}

func TestNewPool(t *testing.T) {
	name1 := "matt"
	name2 := "matthew"
	Convey("TestNewPool", t, func() {
		p := poolx.NewPool(func() *Pojo {
			return &Pojo{name1}
		})

		p.Put(&Pojo{name2})

		Convey("pool Get test", func() {
			get1 := p.Get()
			So(get1.name, ShouldEqual, name2)    // pop the last put one
			So(p.Get().name, ShouldEqual, name1) // no element, so New function will be called
			p.Put(get1)
			So(p.Get().name, ShouldEqual, name2) // no element, so New function will be called

		})

	})

}
