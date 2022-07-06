/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:14:37
 */
package stringx_test

import (
	"testing"

	"github.com/jhunters/goassist/stringx"
	str "github.com/jhunters/goassist/stringx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {

	Convey("TestReverse", t, func() {
		ret, err := str.Reverse("hello")
		So(err, ShouldBeNil)
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

func TestSubstringAfter(t *testing.T) {
	Convey("TestSubstringAfter", t, func() {
		Convey("Test found", func() {
			s := "abc"
			r := stringx.SubstringAfter(s, "a")
			So(r, ShouldEqual, "bc")

			s = "abcba"
			r = stringx.SubstringAfter(s, "b")
			So(r, ShouldEqual, "cba")

			s = "abcba"
			r = stringx.SubstringAfterLast(s, "a")
			So(r, ShouldEqual, "")
		})

		Convey("Test not found", func() {
			s := "abc"
			r := stringx.SubstringAfter(s, "")
			So(r, ShouldEqual, stringx.EMPTY_STRING)

			s = "abcba"
			r = stringx.SubstringAfter(s, "d")
			So(r, ShouldEqual, stringx.EMPTY_STRING)
		})

	})
}

func TestSubstringAfterLast(t *testing.T) {
	Convey("TestSubstringAfter", t, func() {
		Convey("Test found", func() {
			s := "abc"
			r := stringx.SubstringAfterLast(s, "a")
			So(r, ShouldEqual, "bc")

			s = "abcba"
			r = stringx.SubstringAfterLast(s, "b")
			So(r, ShouldEqual, "a")

			s = "abcba"
			r = stringx.SubstringAfterLast(s, "a")
			So(r, ShouldEqual, "")
		})

		Convey("Test not found", func() {
			s := "abc"
			r := stringx.SubstringAfterLast(s, "")
			So(r, ShouldEqual, stringx.EMPTY_STRING)

			s = "abcba"
			r = stringx.SubstringAfterLast(s, "d")
			So(r, ShouldEqual, stringx.EMPTY_STRING)
		})

	})
}
