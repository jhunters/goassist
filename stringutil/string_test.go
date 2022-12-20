/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:14:37
 */
package stringutil_test

import (
	"testing"

	"github.com/jhunters/goassist/stringutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {

	Convey("TestReverse", t, func() {
		ret, err := stringutil.Reverse("hello")
		So(err, ShouldBeNil)
		So(ret, ShouldEqual, "olleh")

		Convey("TestReverse by chinese", func() {
			ret, err := stringutil.Reverse("hello中国")
			So(err, ShouldBeNil)
			So(ret, ShouldEqual, "国中olleh")
		})

		Convey("TestReverse by upper case chinese", func() {
			ret, err := stringutil.Reverse("helloｃｈｉｎｅｓｅ")
			So(err, ShouldBeNil)
			So(ret, ShouldEqual, "ｅｓｅｎｉｈｃolleh")
		})
	})
}

func TestCapitalize(t *testing.T) {

	Convey("TestCapitalize", t, func() {
		ret := stringutil.Capitalize("hello")
		So(ret, ShouldEqual, "Hello")

		ret = stringutil.Capitalize("HEllo")
		So(ret, ShouldEqual, "HEllo")

		ret = stringutil.Capitalize("")
		So(ret, ShouldEqual, "")

		ret = stringutil.Capitalize("121H")
		So(ret, ShouldEqual, "121H")

	})
}

func TestUncapitalize(t *testing.T) {

	Convey("TestCapitalize", t, func() {
		ret := stringutil.Uncapitalize("hello")
		So(ret, ShouldEqual, "hello")

		ret = stringutil.Uncapitalize("HEllo")
		So(ret, ShouldEqual, "hEllo")

		ret = stringutil.Uncapitalize("")
		So(ret, ShouldEqual, "")

		ret = stringutil.Uncapitalize("121H")
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

func TestWrap(t *testing.T) {
	Convey("TestWrap", t, func() {
		s := stringutil.Wrap("hello", "|")
		So(s, ShouldEqual, "|hello|")

	})

}
