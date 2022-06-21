package concurrent_test

import (
	"testing"

	"github.com/jhunters/goassist/concurrent"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCloseChan(t *testing.T) {

	Convey("common close", t, func() {
		ch := make(chan int)
		ok := concurrent.SafeChanClose(ch)
		So(ok, ShouldBeTrue)
	})

	Convey("close nil chan", t, func() {
		var ch chan string
		ok := concurrent.SafeChanClose(ch)
		So(ok, ShouldBeFalse)
	})

	Convey("close closed chan", t, func() {
		ch := make(chan int)
		ok := concurrent.SafeChanClose(ch)
		So(ok, ShouldBeTrue)
		ok = concurrent.SafeChanClose(ch)
		So(ok, ShouldBeFalse)

	})
}
