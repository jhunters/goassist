package atomicx_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/concurrent/syncx/atomicx"
	"github.com/jhunters/goassist/conv"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddandGet(t *testing.T) {

	Convey("TestAddandGet", t, func() {
		Convey("int 32", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[int32](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.AddandGet(16)
			So(v, ShouldEqual, 16)
			v = atomInt.AddandGet(-16)
			So(v, ShouldEqual, 0)
		})
		Convey("int 64", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.AddandGet(16)
			So(v, ShouldEqual, 16)
			v = atomInt.AddandGet(-16)
			So(v, ShouldEqual, 0)
		})

		Convey("uint 32", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uint32](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.AddandGet(16)
			So(v, ShouldEqual, 16)
		})

		Convey("uint 64", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uint64](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.AddandGet(16)
			So(v, ShouldEqual, 16)
		})

		Convey("uintptr", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uintptr](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.AddandGet(16)
			So(v, ShouldEqual, 16)
		})

	})

}

func ExampleNewAtomicInt() {
	// int32
	atomInt := atomicx.NewAtomicInt(conv.ToPtr[int32](0))
	v := atomInt.Get()
	fmt.Println(v)

	// AddandGet
	v = atomInt.AddandGet(16)
	fmt.Println(v)
	v = atomInt.AddandGet(-16)
	fmt.Println(v)

	// CompareAndSet
	b := atomInt.CompareAndSet(0, -100)
	fmt.Println(atomInt.Get(), b)

	// GetAndSet
	v = atomInt.GetAndSet(100)
	fmt.Println(v, atomInt.Get())

	// IncrementAndGet
	v = atomInt.IncrementAndGet()
	fmt.Println(v)

	// decrement and get
	v = atomInt.AddandGet(-1)
	fmt.Println(v)
	// Output:
	// 0
	// 16
	// 0
	// -100 true
	// -100 100
	// 101
	// 100
}

func TestCompareAndSet(t *testing.T) {

	Convey("TestCompareAndSet", t, func() {

		Convey("int64", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			b := atomInt.CompareAndSet(0, 10)
			So(b, ShouldBeTrue)
			So(atomInt.Get(), ShouldEqual, 10)

			b = atomInt.CompareAndSet(0, 20)
			So(b, ShouldBeFalse)
			So(atomInt.Get(), ShouldEqual, 10)
		})

		Convey("int32", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[int32](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			b := atomInt.CompareAndSet(0, 10)
			So(b, ShouldBeTrue)
			So(atomInt.Get(), ShouldEqual, 10)

			b = atomInt.CompareAndSet(0, 20)
			So(b, ShouldBeFalse)
			So(atomInt.Get(), ShouldEqual, 10)
		})

		Convey("uint32", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uint32](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			b := atomInt.CompareAndSet(0, 10)
			So(b, ShouldBeTrue)
			So(atomInt.Get(), ShouldEqual, 10)

			b = atomInt.CompareAndSet(0, 20)
			So(b, ShouldBeFalse)
			So(atomInt.Get(), ShouldEqual, 10)
		})

		Convey("uint64", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uint64](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			b := atomInt.CompareAndSet(0, 10)
			So(b, ShouldBeTrue)
			So(atomInt.Get(), ShouldEqual, 10)

			b = atomInt.CompareAndSet(0, 20)
			So(b, ShouldBeFalse)
			So(atomInt.Get(), ShouldEqual, 10)
		})

		Convey("uintptr", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uintptr](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			b := atomInt.CompareAndSet(0, 10)
			So(b, ShouldBeTrue)
			So(atomInt.Get(), ShouldEqual, 10)

			b = atomInt.CompareAndSet(0, 20)
			So(b, ShouldBeFalse)
			So(atomInt.Get(), ShouldEqual, 10)
		})

	})
}

func TestIncrementAndGet(t *testing.T) {
	Convey("TestIncrementAndGet", t, func() {
		Convey("int64", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.IncrementAndGet()
			So(v, ShouldEqual, 1)
			So(atomInt.Get(), ShouldEqual, 1)
		})

		Convey("int32", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[int32](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.IncrementAndGet()
			So(v, ShouldEqual, 1)
			So(atomInt.Get(), ShouldEqual, 1)
		})

		Convey("uint64", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uint64](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.IncrementAndGet()
			So(v, ShouldEqual, 1)
			So(atomInt.Get(), ShouldEqual, 1)
		})

		Convey("uint32", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uint32](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.IncrementAndGet()
			So(v, ShouldEqual, 1)
			So(atomInt.Get(), ShouldEqual, 1)
		})

		Convey("uintptr", func() {
			atomInt := atomicx.NewAtomicInt(conv.ToPtr[uintptr](0))

			v := atomInt.Get()
			So(v, ShouldBeZeroValue)

			v = atomInt.IncrementAndGet()
			So(v, ShouldEqual, 1)
			So(atomInt.Get(), ShouldEqual, 1)
		})

	})
}

func TestSet(t *testing.T) {
	Convey("TestSet", t, func() {

		atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))
		atomInt.Set(100)
		So(atomInt.Get(), ShouldEqual, 100)
	})

}
