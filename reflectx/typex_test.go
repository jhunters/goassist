package reflectx_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/jhunters/goassist/reflectx"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertibleTo(t *testing.T) {

	Convey("TestConvertibleTo", t, func() {
		var age int = 100
		type Old int

		var age2 Old = 100
		fmt.Println(age2)

		result := reflectx.ConvertibleTo(age, age2)
		So(result, ShouldBeTrue)
	})

}

func TestAssignableIfConvertibleTo(t *testing.T) {

	Convey("TestAssignableIfConvertibleTo", t, func() {
		var age int = 100
		type Old int

		var age2 Old = 200

		// check age2 can convert to age and do convert
		result, ok := reflectx.AssignIfConvertibleTo(age, age2)
		So(ok, ShouldBeTrue)
		So(result, ShouldEqual, 200)
	})

}

func TestIsByteType(t *testing.T) {
	Convey("TestIsByteType", t, func() {
		i := 0
		ok := reflectx.IsByteType(i)
		So(ok, ShouldBeFalse)

		var b byte = 0
		ok = reflectx.IsByteType(b)
		So(ok, ShouldBeTrue)
	})

}

func TestTypeOfName(t *testing.T) {

	Convey("TestTypeOfName", t, func() {

		s := "hello"
		tn := reflectx.TypeOfName(s)
		So(tn, ShouldEqual, "string")

		i := 0
		tn = reflectx.TypeOfName(i)
		So(tn, ShouldEqual, "int")

		ptri := &i
		tn = reflectx.TypeOfName(unsafe.Pointer(ptri))
		So(tn, ShouldEqual, "Pointer")
	})

}
