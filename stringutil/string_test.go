/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:14:37
 */
package stringutil_test

import (
	"testing"

	"github.com/jhunters/goassist/stringutil"
	str "github.com/jhunters/goassist/stringutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {

	Convey("TestReverse", t, func() {
		ret, err := str.Reverse("hello")
		So(err, ShouldBeNil)
		So(ret, ShouldEqual, "olleh")

		Convey("TestReverse by chinese", func() {
			ret, err := str.Reverse("hello中国")
			So(err, ShouldBeNil)
			So(ret, ShouldEqual, "国中olleh")
		})

		Convey("TestReverse by upper case chinese", func() {
			ret, err := str.Reverse("helloｃｈｉｎｅｓｅ")
			So(err, ShouldBeNil)
			So(ret, ShouldEqual, "ｅｓｅｎｉｈｃolleh")
		})
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
			r := stringutil.SubstringAfter(s, "a")
			So(r, ShouldEqual, "bc")

			s = "abcba"
			r = stringutil.SubstringAfter(s, "b")
			So(r, ShouldEqual, "cba")

			s = "abcba"
			r = stringutil.SubstringAfterLast(s, "a")
			So(r, ShouldEqual, "")
		})

		Convey("Test not found", func() {
			s := "abc"
			r := stringutil.SubstringAfter(s, "")
			So(r, ShouldEqual, stringutil.EMPTY_STRING)

			s = "abcba"
			r = stringutil.SubstringAfter(s, "d")
			So(r, ShouldEqual, stringutil.EMPTY_STRING)
		})

	})
}

func TestSubstringAfterLast(t *testing.T) {
	Convey("TestSubstringAfter", t, func() {
		Convey("Test found", func() {
			s := "abc"
			r := stringutil.SubstringAfterLast(s, "a")
			So(r, ShouldEqual, "bc")

			s = "abcba"
			r = stringutil.SubstringAfterLast(s, "b")
			So(r, ShouldEqual, "a")

			s = "abcba"
			r = stringutil.SubstringAfterLast(s, "a")
			So(r, ShouldEqual, "")
		})

		Convey("Test not found", func() {
			s := "abc"
			r := stringutil.SubstringAfterLast(s, "")
			So(r, ShouldEqual, stringutil.EMPTY_STRING)

			s = "abcba"
			r = stringutil.SubstringAfterLast(s, "d")
			So(r, ShouldEqual, stringutil.EMPTY_STRING)
		})

	})
}
