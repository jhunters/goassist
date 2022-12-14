package concurrent_test

import (
	"testing"
	"time"

	"github.com/jhunters/goassist/concurrent"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCloseChan(t *testing.T) {

	Convey("common close", t, func() {
		ch := make(chan int)
		ok := concurrent.SafeCloseChan(ch)
		So(ok, ShouldBeTrue)
	})

	Convey("close nil chan", t, func() {
		var ch chan string
		ok := concurrent.SafeCloseChan(ch)
		So(ok, ShouldBeFalse)
	})

	Convey("close closed chan", t, func() {
		ch := make(chan int)
		ok := concurrent.SafeCloseChan(ch)
		So(ok, ShouldBeTrue)
		ok = concurrent.SafeCloseChan(ch)
		So(ok, ShouldBeFalse)

	})
}

func TestTrySend(t *testing.T) {
	Convey("TestTrySend", t, func() {
		s := "hello"
		Convey("TestTrySend in time", func() {
			ch := make(chan string, 1)
			now := time.Now()
			ok := concurrent.TrySendChan(s, ch, 1*time.Second)
			took := time.Now().Sub(now)
			So(ok, ShouldBeTrue)
			So(took, ShouldBeLessThan, 1*time.Second)
			So(s, ShouldEqual, <-ch)
		})

		Convey("TestTrySend out of time", func() {
			ch := make(chan string, 0)
			now := time.Now()
			ok := concurrent.TrySendChan(s, ch, 1*time.Second)
			took := time.Now().Sub(now)
			So(ok, ShouldBeFalse)
			So(took, ShouldBeGreaterThanOrEqualTo, 1*time.Second)
			So(took, ShouldBeLessThan, 2*time.Second)
			v := "hello1"
			go func(s string) {
				ch <- s
			}(v)

			So(v, ShouldEqual, <-ch)
		})

		Convey("TestTrySend with no timeout", func() {
			ch := make(chan string, 0)
			go func() {
				concurrent.TrySendChan(s, ch, 0)
			}()
			So(s, ShouldEqual, <-ch)

		})
	})

}

func TestTryReceive(t *testing.T) {
	Convey("TestTryReceive", t, func() {
		s := "hello"
		Convey("TestTryReceive in time", func() {
			ch := make(chan string, 1)
			ch <- s
			now := time.Now()
			ok, f := concurrent.TryRecevieChan(ch, 1*time.Second)
			took := time.Now().Sub(now)
			So(ok, ShouldBeTrue)
			So(s, ShouldEqual, f)
			So(took, ShouldBeLessThan, 1*time.Second)
		})

		Convey("TestTryReceive out of time", func() {
			ch := make(chan string, 1)
			now := time.Now()
			ok, _ := concurrent.TryRecevieChan(ch, 1*time.Second)
			took := time.Now().Sub(now)
			So(ok, ShouldBeFalse)
			So(took, ShouldBeGreaterThanOrEqualTo, 1*time.Second)
			So(took, ShouldBeLessThan, 2*time.Second)

			ch <- s
			So(s, ShouldEqual, <-ch)
		})

		Convey("TestTryReceive with no timeout", func() {
			ch := make(chan string, 1)
			go func(s string) {
				ch <- s
			}(s)
			ok, v := concurrent.TryRecevieChan(ch, 0)
			So(ok, ShouldBeTrue)
			So(s, ShouldEqual, v)

		})
	})

}
