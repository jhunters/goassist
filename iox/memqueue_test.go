package iox_test

import (
	"testing"
	"time"

	"github.com/jhunters/goassist/iox"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemQueue(t *testing.T) {

	Convey("TestMemQueue", t, func() {
		queue := iox.NewMemQueue()
		defer queue.Close()

		dataArr := make([][]byte, 100)
		for i := 0; i < 100; i++ {
			dataArr[i] = make([]byte, i)
		}

		go func() {
			for _, v := range dataArr {
				time.Sleep(10 * time.Millisecond)
				queue.Enqueue(v)
			}
		}()

		for i := 0; i < 100; i++ {
			b, _ := queue.Dequeue()
			So(b, ShouldResemble, dataArr[i])
		}
	})

}
