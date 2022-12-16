package containerx_test

import (
	"testing"

	"github.com/jhunters/goassist/containerx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewStack(t *testing.T) {
	Convey("TestNewStack", t, func() {
		s := containerx.NewStack[*RingPojo]()
		So(s.IsEmpty(), ShouldBeTrue)
	})
	Convey("TestNewStackSize", t, func() {
		s := containerx.NewStackSize[*RingPojo](16)
		size := s.Cap()
		So(size, ShouldEqual, 16)

		s = containerx.NewStackSize[*RingPojo](-1)
		So(s.IsEmpty(), ShouldBeTrue)
	})

}

func TestStackPushPoP(t *testing.T) {
	Convey("TestStackPushPoP", t, func() {
		s := containerx.NewStackSize[*RingPojo](16)
		So(s.IsEmpty(), ShouldBeTrue)
		s.Push(NewRingPojo("matt"))
		size := s.Cap()
		So(size, ShouldEqual, 15)
		So(s.IsEmpty(), ShouldBeFalse)

		rp := s.Pop()
		So(rp, ShouldNotBeNil)
		So(rp.Name, ShouldEqual, "matt")

		rp = s.Pop()
		So(rp, ShouldBeNil)

		So(s.IsEmpty(), ShouldBeTrue)
	})

}

func TestStackResize(t *testing.T) {
	Convey("TestStackResize", t, func() {
		s := containerx.NewStackSize[*RingPojo](2)
		So(s.Size(), ShouldEqual, 0)
		So(s.Cap(), ShouldEqual, 2)

		s.Push(NewRingPojo("matt"))
		s.Push(NewRingPojo("xml"))

		So(s.Cap(), ShouldEqual, 0)
		So(s.Size(), ShouldEqual, 2)

		// resize
		s.Push(NewRingPojo("michael"))
		So(s.Size(), ShouldEqual, 3)
		So(s.Cap(), ShouldEqual, 1)
	})
}
