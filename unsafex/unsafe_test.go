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
	age    int64
	gender bool
	other  int64
	other2 bool
}

func TestStructMappingToArray(t *testing.T) {
	Convey("test struct convert to array", t, func() {

		s := Simple{age: 100}
		size := unsafe.Sizeof(s)

		ret := unsafex.MappingToArray(s)
		So(len(ret), ShouldEqual, size)
		buff := bytes.NewBuffer(ret)
		var age int64
		binary.Read(buff, binary.LittleEndian, &age)
		So(age, ShouldEqual, s.age)
	})
}

func TestArrayMappingToStruct(t *testing.T) {
	Convey("test convert array to struct ", t, func() {

		s := Simple{age: 100}
		size := unsafe.Sizeof(s)

		arr := make([]byte, size)
		result := unsafex.ArrayMapping[Simple](arr)

		binary.LittleEndian.PutUint64(arr, 1000)
		So(1000, ShouldEqual, result.age)
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
