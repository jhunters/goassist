package stack_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/container/stack"
	"github.com/jhunters/goassist/conv"
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
		s := stack.NewStack[*StackPojo]()
		So(s.IsEmpty(), ShouldBeTrue)
	})
	Convey("TestNewStackSize", t, func() {
		s := stack.NewStackSize[*StackPojo](16)
		size := s.Cap()
		So(size, ShouldEqual, 16)

		s = stack.NewStackSize[*StackPojo](-1)
		So(s.IsEmpty(), ShouldBeTrue)
	})

}

func TestStackPushPoP(t *testing.T) {
	Convey("TestStackPushPoP", t, func() {
		s := stack.NewStackSize[*StackPojo](16)
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
		s := stack.NewStackSize[*StackPojo](2)
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
		s := stack.NewStackSize[*StackPojo](2)
		s.Push(NewStackPojo("matt"))
		s.Push(NewStackPojo("xml"))

		cs := s.Copy()
		So(cs.Cap(), ShouldEqual, s.Cap())
		So(cs.Size(), ShouldEqual, s.Size())
	})
}

func ExampleNewStack() {
	// event bus send event as FILO mode
	eventBus := stack.NewStack[*string]()
	fmt.Println(eventBus.IsEmpty())

	eventBus.Push(conv.ToPtr("Sig No 109"))
	eventBus.Push(conv.ToPtr("Sig No 282"))

	fmt.Println(eventBus.Size())
	fmt.Println(*eventBus.Pop())
	fmt.Println(*eventBus.Pop())

	// Output:
	// true
	// 2
	// Sig No 282
	// Sig No 109
}

func ExampleNewStackSize() {
	// event bus send event as FILO mode
	eventBus := stack.NewStackSize[*string](10)
	fmt.Println(eventBus.IsEmpty())

	eventBus.Push(conv.ToPtr("Sig No 109"))
	eventBus.Push(conv.ToPtr("Sig No 282"))

	fmt.Println(eventBus.Size())
	fmt.Println(eventBus.Cap())
	fmt.Println(*eventBus.Pop())
	fmt.Println(*eventBus.Pop())

	// Output:
	// true
	// 2
	// 8
	// Sig No 282
	// Sig No 109

}
