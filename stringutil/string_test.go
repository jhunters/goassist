/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:14:37
 */
package stringutil_test

import (
	"fmt"
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

func TestAbbreviate(t *testing.T) {
	Convey("TestAbbreviate", t, func() {
		s, err := stringutil.Abbreviate("abcdefghijklmno", "---", -1, 10)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "abcdefg---")

		s, err = stringutil.Abbreviate("abcdefghijklmno", ",", 0, 10)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "abcdefghi,")

		s, err = stringutil.Abbreviate("abcdefghijklmno", "...", 6, 10)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "...ghij...")

		s, err = stringutil.Abbreviate("abcdefghij", "abra", 0, 4)
		So(err, ShouldNotBeNil)
	})

}

func TestAbbreviateMiddle(t *testing.T) {
	Convey("TestAbbreviateMiddle", t, func() {
		s := stringutil.AbbreviateMiddle("abcdef", ".", 4)
		So(s, ShouldEqual, "ab.f")

		s = stringutil.AbbreviateMiddle("abc", ".", 3)
		So(s, ShouldEqual, "abc")
	})

}

func TestIsNumber(t *testing.T) {
	Convey("TestIsNumber", t, func() {
		Convey("hex", func() {
			isNumber := stringutil.IsNumber("0x12")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("0xac")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("0x")
			So(isNumber, ShouldBeFalse)

			isNumber = stringutil.IsNumber("0x093g")
			So(isNumber, ShouldBeFalse)
		})
		Convey("octal", func() {
			isNumber := stringutil.IsNumber("0o10")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("0O17")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("0o18")
			So(isNumber, ShouldBeFalse)

			isNumber = stringutil.IsNumber("0O093g")
			So(isNumber, ShouldBeFalse)
		})
		Convey("binary", func() {
			isNumber := stringutil.IsNumber("0b10")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("0B11")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("0b18")
			So(isNumber, ShouldBeFalse)

			isNumber = stringutil.IsNumber("0B093g")
			So(isNumber, ShouldBeFalse)
		})
		Convey("common", func() {
			isNumber := stringutil.IsNumber("-12.11")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("19.1")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("-12.1.1")
			So(isNumber, ShouldBeFalse)

			isNumber = stringutil.IsNumber("12e1")
			So(isNumber, ShouldBeTrue)

			isNumber = stringutil.IsNumber("12e-9")
			So(isNumber, ShouldBeTrue)
		})

	})

}

func TestParseInt(t *testing.T) {
	Convey("TestParseInt", t, func() {
		// hex string
		v, err := stringutil.ParseInt("-0x19ac")
		So(v, ShouldEqual, -6572)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseInt("0X19ac")
		So(v, ShouldEqual, 6572)
		So(err, ShouldBeNil)

		// octal string
		v, err = stringutil.ParseInt("0o1312")
		So(v, ShouldEqual, 714)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseInt("-0O1312")
		So(v, ShouldEqual, -714)
		So(err, ShouldBeNil)

		// binary string
		v, err = stringutil.ParseInt("0b1011")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseInt("-0B1011")
		So(v, ShouldEqual, -11)
		So(err, ShouldBeNil)

		// e
		v, err = stringutil.ParseInt("11e10")
		So(v, ShouldEqual, 110000000000)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseInt("110e-1")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseInt("11e-1")
		So(v, ShouldEqual, 1)
		So(err, ShouldBeNil)

		// invalid int number
		v, err = stringutil.ParseInt("11ee-1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)

		v, err = stringutil.ParseInt("--11e-1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)
		v, err = stringutil.ParseInt("-+11e-1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)
		v, err = stringutil.ParseInt("1.1e1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)
	})

}
func ExampleParseInt() {
	// hex string
	v, err := stringutil.ParseInt("-0x19ac")
	fmt.Println(v, err)

	v, err = stringutil.ParseInt("0X19ac")
	fmt.Println(v, err)

	// octal string
	v, err = stringutil.ParseInt("0o1312")
	fmt.Println(v, err)
	v, err = stringutil.ParseInt("-0O1312")
	fmt.Println(v, err)

	// binary string
	v, err = stringutil.ParseInt("0b1011")
	fmt.Println(v, err)
	v, err = stringutil.ParseInt("-0B1011")
	fmt.Println(v, err)

	// e
	v, err = stringutil.ParseInt("11e10")
	fmt.Println(v, err)

	v, err = stringutil.ParseInt("110e-1")
	fmt.Println(v, err)

	v, err = stringutil.ParseInt("11e-1")
	fmt.Println(v, err)

	// Output:
	// 	-6572 <nil>
	// 6572 <nil>
	// 714 <nil>
	// -714 <nil>
	// 11 <nil>
	// -11 <nil>
	// 110000000000 <nil>
	// 11 <nil>
	// 1 <nil>

}
func TestParseFloat(t *testing.T) {

	Convey("TestParseFloat", t, func() {
		// hex string
		v, err := stringutil.ParseFloat("-0x19ac")
		So(v, ShouldEqual, -6572)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseFloat("0X19ac")
		So(v, ShouldEqual, 6572)
		So(err, ShouldBeNil)

		// octal string
		v, err = stringutil.ParseFloat("0o1312")
		So(v, ShouldEqual, 714)
		So(err, ShouldBeNil)
		v, err = stringutil.ParseFloat("-0O1312")
		So(v, ShouldEqual, -714)
		So(err, ShouldBeNil)

		// binary string
		v, err = stringutil.ParseFloat("0b1011")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)
		v, err = stringutil.ParseFloat("-0B1011")
		So(v, ShouldEqual, -11)
		So(err, ShouldBeNil)

		// e
		v, err = stringutil.ParseFloat("11e10")
		So(v, ShouldEqual, 110000000000)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseFloat("110e-1")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)

		v, err = stringutil.ParseFloat("11e-1")
		So(v, ShouldEqual, 1.1)
		So(err, ShouldBeNil)
	})

}

func ExampleParseFloat() {
	// hex string
	v, err := stringutil.ParseFloat("-0x19ac")
	fmt.Println(v, err)

	v, err = stringutil.ParseFloat("0X19ac")
	fmt.Println(v, err)

	// octal string
	v, err = stringutil.ParseFloat("0o1312")
	fmt.Println(v, err)
	v, err = stringutil.ParseFloat("-0O1312")
	fmt.Println(v, err)

	// binary string
	v, err = stringutil.ParseFloat("0b1011")
	fmt.Println(v, err)
	v, err = stringutil.ParseFloat("-0B1011")
	fmt.Println(v, err)

	// e
	v, err = stringutil.ParseFloat("11e10")
	fmt.Println(v, err)

	v, err = stringutil.ParseFloat("110e-1")
	fmt.Println(v, err)

	v, err = stringutil.ParseFloat("11e-1")
	fmt.Println(v, err)

	// Output:
	// 	-6572 <nil>
	// 6572 <nil>
	// 714 <nil>
	// -714 <nil>
	// 11 <nil>
	// -11 <nil>
	// 1.1e+11 <nil>
	// 11 <nil>
	// 1.1 <nil>
}
