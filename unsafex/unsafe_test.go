package unsafex_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"

	"github.com/jhunters/goassist/unsafex"
	. "github.com/smartystreets/goconvey/convey"
)

type Simple struct {
	Age    int64
	Gender bool
	Other  int64
	Other2 bool
}

func TestStructMappingToArray(t *testing.T) {
	Convey("test struct convert to array", t, func() {

		s := Simple{Age: 100}
		size := unsafe.Sizeof(s)

		ret := unsafex.MappingToArray(s)
		So(len(ret), ShouldEqual, size)
		buff := bytes.NewBuffer(ret)
		var age int64
		binary.Read(buff, binary.LittleEndian, &age)
		So(age, ShouldEqual, s.Age)
	})
}

func TestArrayMappingToStruct(t *testing.T) {
	Convey("test convert array to struct ", t, func() {

		s := Simple{Age: 100}
		size := unsafe.Sizeof(s)

		arr := make([]byte, size)
		result := unsafex.ArrayMapping[Simple](arr)

		binary.LittleEndian.PutUint64(arr, 1000)
		So(1000, ShouldEqual, result.Age)
		So(100, ShouldEqual, s.Age)
	})
}

func TestStringToSlice(t *testing.T) {
	Convey("TestStringToSlice", t, func() {
		b := unsafex.StringToSlice("hello world!")
		s := unsafex.SliceToString(b)
		So(s, ShouldEqual, "hello world!")
	})
}

func TestValueSizeof(t *testing.T) {
	Convey("TestValueSizeof", t, func() {
		i := 0
		sz := unsafex.ValueSizeof(i)
		So(sz, ShouldEqual, 8)

		sz = unsafex.ValueSizeof(&i)
		So(sz, ShouldEqual, 8)
	})
}

type VIP struct {
	Name string
}

func TestCopy(t *testing.T) {
	Convey("TestCopy", t, func() {
		s := &VIP{}
		s1 := unsafex.Copy(s)
		s.Name = "matt"
		So(s1.Name, ShouldBeEmpty)
	})
}

func ExampleCopy() {

	s := struct {
		Name string
	}{}
	s.Name = "hello"
	s2 := unsafex.Copy(&s) // to a new value
	s2.Name = "world"      // 传递给 s2, 但 s2 已经是新创建的
	fmt.Println(s.Name)

	// Output:
	// hello
}

type AllInt struct {
	V1 int
	V2 int
	V3 int
	V4 int
	V5 int
}

func TestSlice(t *testing.T) {
	v := &AllInt{1, 2, 3, 4, 5}

	Convey("TestSlice", t, func() {

		expect := []int{1, 2, 3, 4, 5}
		// convert AllInt struct's field value to  int slice
		nv := unsafex.Slice[AllInt, int](v, 5)
		So(nv, ShouldResemble, expect)

		// convert AllInt struct's field value to  int slice
		nv = unsafex.Slice[AllInt, int](v, 2)
		So(nv, ShouldResemble, expect[:2])
	})

}

func ExampleSlice() {
	v := &AllInt{1, 2, 3, 4, 5}

	// convert AllInt struct's field value to  int slice
	nv := unsafex.Slice[AllInt, int](v, 5)
	fmt.Println(nv)

	// convert AllInt struct's field value to  int slice
	nv = unsafex.Slice[AllInt, int](v, 2)
	fmt.Println(nv)

	// Output:
	// [1 2 3 4 5]
	// [1 2]
}

func TestOffsetValue(t *testing.T) {

	v := &AllInt{1, 2, 3, 4, 5}

	Convey("TestOffsetValue", t, func() {
		nv := unsafex.OffsetValue[AllInt, int](v, 2)
		So(nv, ShouldEqual, &v.V3)
	})

}

func ExampleOffsetValue() {
	v := &AllInt{1, 2, 3, 4, 5}

	nv := unsafex.OffsetValue[AllInt, int](v, 2) // equal to v[2]
	fmt.Println(*nv)

	// Output:
	// 3
}

func ExampleAs() {
	var ui64 uint64 = 10
	// convert uint64 to int64
	var v *int64 = unsafex.As[uint64, int64](&ui64)
	fmt.Println(*v)

	var f64 float64 = 100
	// convert float64 to int64
	v = unsafex.As[float64, int64](&f64)
	fmt.Println(*v) // will get a unsafe return

	// Output:
	// 10
	// 4636737291354636288

}
