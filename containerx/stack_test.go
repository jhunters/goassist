package containerx_test

import (
	"testing"

	"github.com/jhunters/goassist/containerx"
	. "github.com/smartystreets/goconvey/convey"
)

type StackPojo struct {
	Name string
}

func NewStackPojo(name string) *StackPojo {
	return &StackPojo{Name: name}
}

func TestNewStack(t *testing.T) {
	Convey("TestNewStack", t, func() {
		s := containerx.NewStack[*StackPojo]()
		So(s.IsEmpty(), ShouldBeTrue)
	})
	Convey("TestNewStackSize", t, func() {
		s := containerx.NewStackSize[*StackPojo](16)
		size := s.Cap()
		So(size, ShouldEqual, 16)

		s = containerx.NewStackSize[*StackPojo](-1)
		So(s.IsEmpty(), ShouldBeTrue)
	})

}

func TestStackPushPoP(t *testing.T) {
	Convey("TestStackPushPoP", t, func() {
		s := containerx.NewStackSize[*StackPojo](16)
		So(s.IsEmpty(), ShouldBeTrue)
		s.Push(NewStackPojo("matt"))
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
		s := containerx.NewStackSize[*StackPojo](2)
		So(s.Size(), ShouldEqual, 0)
		So(s.Cap(), ShouldEqual, 2)

		s.Push(NewStackPojo("matt"))
		s.Push(NewStackPojo("xml"))

		So(s.Cap(), ShouldEqual, 0)
		So(s.Size(), ShouldEqual, 2)

		// resize
		s.Push(NewStackPojo("michael"))
		So(s.Size(), ShouldEqual, 3)
		So(s.Cap(), ShouldEqual, 1)
	})
}

func TestStackCopy(t *testing.T) {
	Convey("TestStackCopy", t, func() {
		s := containerx.NewStackSize[*StackPojo](2)
		s.Push(NewStackPojo("matt"))
		s.Push(NewStackPojo("xml"))

		cs := s.Copy()
		So(cs.Cap(), ShouldEqual, s.Cap())
		So(cs.Size(), ShouldEqual, s.Size())
	})
}
