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

func TestSetGet(t *testing.T) {
	Convey("TestSetGet", t, func() {
		atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))
		atomInt.Set(100)
		So(atomInt.Get(), ShouldEqual, 100)

		atomInt.Store(200)
		So(atomInt.Load(), ShouldEqual, 200)
	})

	Convey("GetAndSet", t, func() {
		atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))
		v := atomInt.GetAndSet(1)
		So(v, ShouldEqual, 0)
		So(atomInt.Get(), ShouldEqual, 1)
	})

}

func ExampleAtomicInt_AddandGet() {
	atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](100))

	v := atomInt.AddandGet(20)
	fmt.Println(v)

	v = atomInt.AddandGet(-20)
	fmt.Println(v)

	// Output:
	// 120
	// 100
}

func ExampleAtomicInt_IncrementAndGet() {
	atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](100))

	v := atomInt.IncrementAndGet()
	fmt.Println(v)

	v = atomInt.IncrementAndGet()
	fmt.Println(v)

	// Output:
	// 101
	// 102
}

func ExampleAtomicInt_CompareAndSet() {
	atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](100))

	update := atomInt.CompareAndSet(100, 101)
	fmt.Println(update)

	update = atomInt.CompareAndSet(100, 101)
	fmt.Println(update)

	// Output:
	// true
	// false
}

func ExampleAtomicInt_GetAndSet() {
	atomInt := atomicx.NewAtomicInt(conv.ToPtr[int64](0))
	v := atomInt.GetAndSet(1)

	fmt.Println(v, atomInt.Get())

	// Output:
	// 0 1
}
