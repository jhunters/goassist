package queue_test

import (
	"testing"
	"time"

	"github.com/jhunters/goassist/container/queue"
	"github.com/jhunters/goassist/conv"
	. "github.com/smartystreets/goconvey/convey"
)

type QueuePojo struct {
	Name string
}

func newQueuePojo(name string) *QueuePojo {
	return &QueuePojo{name}
}

func TestQueueEnqueueDequeue(t *testing.T) {
	Convey("TestEnqueueDequeue", t, func() {

		q := queue.NewQueue[*QueuePojo]()

		q.Enqueue(newQueuePojo("hello"))
		q.Enqueue(newQueuePojo("world"))
		q.Enqueue(newQueuePojo("!"))

		So(q.Dequeue().Name, ShouldEqual, "hello")
		So(q.Dequeue().Name, ShouldEqual, "world")
		So(q.Dequeue().Name, ShouldEqual, "!")
	})

}

func TestQueueSubsrcibe(t *testing.T) {

	Convey("TestQueueSubsrcibe", t, func() {

		q := queue.NewQueue[*QueuePojo]()

		rcv1 := make([]string, 0)
		rcv2 := make([]string, 0)

		q.SubScribe(func(qp *QueuePojo) {
			rcv1 = append(rcv1, qp.Name)
			time.Sleep(100 * time.Millisecond)
		})

		q.SubScribe(func(qp *QueuePojo) {
			rcv2 = append(rcv2, qp.Name)
			time.Sleep(100 * time.Millisecond)
		})

		for i := 0; i < 10; i++ {
			q.Enqueue(newQueuePojo(conv.Itoa(i)))
		}

		time.Sleep(time.Second)
		So(len(rcv1)+len(rcv2), ShouldEqual, 10)
	})
}

func TestClear(t *testing.T) {

	Convey("TestQueueSubsrcibe", t, func() {
		q := queue.NewQueue[*QueuePojo]()

		q.Enqueue(newQueuePojo("hello"))
		q.Enqueue(newQueuePojo("world"))
		q.Enqueue(newQueuePojo("!"))

		q.Clear()

		So(q.Dequeue(), ShouldBeNil)
	})
}
