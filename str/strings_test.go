/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:14:37
 */
package str_test

import (
	"testing"

	str "github.com/jhunters/goassist/str"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {

	Convey("TestReverse", t, func() {
		ret := str.Reverse("hello")
		So(ret, ShouldEqual, "olleh")
	})
}

func TestCapitalize(t *testing.T) {

	Convey("TestCapitalize", t, func() {
		ret := str.Capitalize("hello")
		So(ret, ShouldEqual, "Hello")

		ret = str.Capitalize("HEllo")
		So(ret, ShouldEqual, "HEllo")

		ret = str.Capitalize("")
		So(ret, ShouldEqual, "")

		ret = str.Capitalize("121H")
		So(ret, ShouldEqual, "121H")

	})
}

func TestUncapitalize(t *testing.T) {

	Convey("TestCapitalize", t, func() {
		ret := str.Uncapitalize("hello")
		So(ret, ShouldEqual, "hello")

		ret = str.Uncapitalize("HEllo")
		So(ret, ShouldEqual, "hEllo")

		ret = str.Uncapitalize("")
		So(ret, ShouldEqual, "")

		ret = str.Uncapitalize("121H")
		So(ret, ShouldEqual, "121H")
	})
}
