/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:14:37
 */
package stringutil_test

import (
	"fmt"
	"testing"
	"unsafe"

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

func ExampleReverse() {
	ret, err := stringutil.Reverse("hello")
	fmt.Println(ret, err)

	ret, err = stringutil.Reverse("helloｃｈｉｎｅｓｅ")
	fmt.Println(ret, err)

	// Output:
	// olleh <nil>
	// ｅｓｅｎｉｈｃolleh <nil>
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

func ExampleCapitalize() {
	ret := stringutil.Capitalize("hello")
	fmt.Println(ret)

	ret = stringutil.Capitalize("121H")
	fmt.Println(ret)

	ret = stringutil.Capitalize("HEllo")
	fmt.Println(ret)

	// Output:
	// Hello
	// 121H
	// HEllo
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

func ExampleUncapitalize() {
	ret := stringutil.Uncapitalize("HEllo")
	fmt.Println(ret)

	ret = stringutil.Uncapitalize("hello")
	fmt.Println(ret)

	// Output:
	// hEllo
	// hello
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

func ExampleSubstringAfter() {
	s := "abc"
	r := stringutil.SubstringAfter(s, "a")
	fmt.Println(r)

	s = "abcba"
	r = stringutil.SubstringAfter(s, "b")
	fmt.Println(r)

	// Output:
	// bc
	// cba
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

func ExampleSubstringAfterLast() {
	s := "abc"
	r := stringutil.SubstringAfterLast(s, "a")
	fmt.Println(r)

	s = "abcba"
	r = stringutil.SubstringAfterLast(s, "b")
	fmt.Println(r)

	// Output:
	// bc
	// a
}

func TestWrap(t *testing.T) {
	Convey("TestWrap", t, func() {
		s := stringutil.Wrap("hello", "|")
		So(s, ShouldEqual, "|hello|")

	})
}

func ExampleWrap() {
	s := stringutil.Wrap("hello", "|")
	fmt.Println(s)

	// Output:
	// |hello|
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

		_, err = stringutil.Abbreviate("abcdefghij", "abra", 0, 4)
		So(err, ShouldNotBeNil)
	})

}

func ExampleAbbreviate() {
	s, err := stringutil.Abbreviate("abcdefghijklmno", "---", -1, 10)
	fmt.Println(s, err)

	s, err = stringutil.Abbreviate("abcdefghijklmno", ",", 0, 10)
	fmt.Println(s, err)

	s, err = stringutil.Abbreviate("abcdefghijklmno", "...", 6, 10)
	fmt.Println(s, err)

	// Output:
	// abcdefg--- <nil>
	// abcdefghi, <nil>
	// ...ghij... <nil>
}

func TestAbbreviateMiddle(t *testing.T) {
	Convey("TestAbbreviateMiddle", t, func() {
		s := stringutil.AbbreviateMiddle("abcdef", ".", 4)
		So(s, ShouldEqual, "ab.f")

		s = stringutil.AbbreviateMiddle("abc", ".", 3)
		So(s, ShouldEqual, "abc")
	})

}

func ExampleAbbreviateMiddle() {
	s := stringutil.AbbreviateMiddle("abcdef", ".", 4)
	fmt.Println(s)

	s = stringutil.AbbreviateMiddle("abc", ".", 3)
	fmt.Println(s)

	// Output:
	// ab.f
	// abc
}

func TestSubstringBefore(t *testing.T) {
	Convey("TestSubstringBefore", t, func() {
		s := "hello world to beijin"
		separator := "wor"

		// exist
		ns := stringutil.SubstringBefore(s, separator)
		So(ns, ShouldEqual, "hello ")

		ns = stringutil.SubstringBefore(s, "notexist")
		So(ns, ShouldEqual, "")

		s = "hello world hello world"
		ns = stringutil.SubstringBefore(s, separator)
		So(ns, ShouldEqual, "hello ")
	})

}

func ExampleSubstringBefore() {
	s := "helloworld to beijin"
	separator := "wor"

	// exist
	ns := stringutil.SubstringBefore(s, separator)
	fmt.Println(ns)

	s = "helloworld hello world"
	ns = stringutil.SubstringBefore(s, separator)
	fmt.Println(ns)

	// Output:
	// hello
	// hello
}
func TestSubstringBeforeLast(t *testing.T) {
	Convey("TestSubstringBefore", t, func() {
		s := "hello world to beijin"
		separator := "wor"

		// exist
		ns := stringutil.SubstringBeforeLast(s, separator)
		So(ns, ShouldEqual, "hello ")

		ns = stringutil.SubstringBeforeLast(s, "notexist")
		So(ns, ShouldEqual, s)

		s = "hello world hello world"
		ns = stringutil.SubstringBeforeLast(s, separator)
		So(ns, ShouldEqual, "hello world hello ")
	})

}

func ExampleSubstringBeforeLast() {
	s := "helloworldtobeijin"
	separator := "wor"

	// exist
	ns := stringutil.SubstringBeforeLast(s, separator)
	fmt.Println(ns)

	ns = stringutil.SubstringBeforeLast(s, "notexist")
	fmt.Println(ns)

	s = "helloworldhelloworld"
	ns = stringutil.SubstringBeforeLast(s, separator)
	fmt.Println(ns)

	// Output:
	// hello
	// helloworldtobeijin
	// helloworldhello
}

func TestStringSlice(t *testing.T) {
	Convey("TestStringSlice", t, func() {
		s := "hello world"
		arr := stringutil.StringToSlice(s)
		So(len(arr), ShouldEqual, len(s))
		So(string(arr), ShouldEqual, s)

		ns := stringutil.SliceToString(arr)
		So(len(arr), ShouldEqual, len(ns))
		So(string(arr), ShouldEqual, ns)

		narr := stringutil.StringToSlice(s)
		So(unsafe.Pointer(&narr[0]), ShouldEqual, unsafe.Pointer(&arr[0]))
	})
}

func ExampleStringToSlice() {
	s := "hello world"
	arr := stringutil.StringToSlice(s)
	fmt.Println(string(arr))

	// Output:
	// hello world

}

func TestRepeat(t *testing.T) {
	Convey("TestRepeat", t, func() {
		s := stringutil.Repeat(97, 10)
		So(s, ShouldEqual, "aaaaaaaaaa")
	})
}

func ExampleRepeat() {
	s := stringutil.Repeat(97, 10)
	fmt.Println(s)

	// Output:
	// aaaaaaaaaa
}

func TestIsEmptyOrBlank(t *testing.T) {
	Convey("TestIsEmptyOrBlank", t, func() {
		So(stringutil.IsEmpty(""), ShouldBeTrue)
		So(stringutil.IsEmpty(" "), ShouldBeFalse)
		So(stringutil.IsBlank(""), ShouldBeTrue)
		So(stringutil.IsBlank(" "), ShouldBeTrue)
		So(stringutil.IsBlank(" a"), ShouldBeFalse)
	})
}

func ExampleIsEmpty() {
	fmt.Println(stringutil.IsEmpty(""))
	fmt.Println(stringutil.IsEmpty(" "))

	// Output:
	// true
	// false
}

func ExampleIsBlank() {
	fmt.Println(stringutil.IsBlank(""))
	fmt.Println(stringutil.IsBlank(" "))
	fmt.Println(stringutil.IsBlank(" a"))

	// Output:
	// true
	// true
	// false
}

func TestExapnd(t *testing.T) {
	Convey("TestExapnd", t, func() {

		Convey("TestExapnd simple", func() {

			v, err := stringutil.Expand("dfsdfsd ${ab} ${sdabc} fsdfsd", "${", "}", func(key string) string {
				So(key, ShouldBeIn, []string{"ab", "sdabc"})
				return key
			})
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "dfsdfsd ab sdabc fsdfsd")
		})
		Convey("TestExapnd recursive", func() {

			v, err := stringutil.Expand("dfsdfsd ${sd${ab}c} fsdfsd", "${", "}", func(key string) string {
				So(key, ShouldBeIn, []string{"ab", "sdabc"})
				return key
			})
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "dfsdfsd sdabc fsdfsd")
		})

		Convey("TestExapnd no match", func() {

			v, err := stringutil.Expand("dfsdfsd ${sd${ab}c} fsdfsd", "#{", "}", func(key string) string {
				return key
			})
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "dfsdfsd ${sd${ab}c} fsdfsd")
		})
	})
}

func ExampleExpand() {
	// Expand simple expression
	v, err := stringutil.Expand("please send mail to ${name} at ${address}!", "${", "}", func(placeholderKey string) string {
		mp := map[string]string{"name": "matt", "address": "shanghai pudong"}
		return mp[placeholderKey]
	})
	if err == nil {
		fmt.Println(v)
	}

	// Expand recursive expression
	v, err = stringutil.Expand("please send mail to #{to#{name}} at #{address}", "#{", "}", func(placeholderKey string) string {
		mp := map[string]string{"name": "matt", "tomatt": "matt's company", "address": "shanghai pudong"}
		return mp[placeholderKey]
	})
	if err == nil {
		fmt.Println(v)
	}

	// Output:
	// please send mail to matt at shanghai pudong!
	// please send mail to matt's company at shanghai pudong
}
