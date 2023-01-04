package queue_test

import (
	"testing"

	"github.com/jhunters/goassist/container/queue"
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
