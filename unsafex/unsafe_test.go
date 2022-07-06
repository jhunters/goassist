package unsafex_test

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"

	"github.com/jhunters/goassist/unsafex"
	"github.com/smartystreets/goconvey/convey"
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
	convey.Convey("test convert array to struct ", t, func() {

		s := Simple{age: 100}
		size := unsafe.Sizeof(s)

		arr := make([]byte, size)
		result := unsafex.ArrayMapping[Simple](arr)

		binary.LittleEndian.PutUint64(arr, 1000)
		So(1000, ShouldEqual, result.age)
	})
}
