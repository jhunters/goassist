package reflects_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/reflects"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertibleTo(t *testing.T) {

	Convey("TestConvertibleTo", t, func() {
		var age int = 100
		type Old int

		var age2 Old = 100
		fmt.Println(age2)

		result := reflects.ConvertibleTo(age, age2)
		So(result, ShouldBeTrue)
	})

}

func TestAssignableIfConvertibleTo(t *testing.T) {

	Convey("TestAssignableIfConvertibleTo", t, func() {
		var age int = 100
		type Old int

		var age2 Old = 200

		// check age2 can convert to age and do convert
		result, ok := reflects.AssignIfConvertibleTo(age, age2)
		So(ok, ShouldBeTrue)
		So(result, ShouldEqual, 200)
	})

}
