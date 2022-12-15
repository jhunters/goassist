/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:12:42
 */

package stringx

import (
	"errors"
	"strings"
	"unicode/utf8"
	"unsafe"

	"github.com/jhunters/goassist/arrayx"
	"github.com/jhunters/goassist/unsafex"
)

const (
	INDEX_NOT_FOUND = -1
	EMPTY_STRING    = ""
)

// Reverse to reverse the string
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

// Capitalize a String changing the first letter to upper case
// str.Capitalize("hello") = "Hello"
// str.Capitalize("HEllo") = "HEllo"
// str.Capitalize("") = ""
// str.Capitalize("12h") = "12h"
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	b := []byte(s)
	b[0] = []byte(strings.ToUpper(string(b[0])))[0]
	return string(b)
}

// Uncapitalize a String changing the first letter to lower case
// str.Capitalize("hello") = "hello"
// str.Capitalize("HEllo") = "hEllo"
// str.Capitalize("") = ""
// str.Capitalize("12h") = "12h"
func Uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	b := []byte(s)
	b[0] = []byte(strings.ToLower(string(b[0])))[0]
	return string(b)
}

// SubstringAfter Gets the substring after the first occurrence of a separator
func SubstringAfter(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.Index(s, separator)
	if pos == INDEX_NOT_FOUND {
		return EMPTY_STRING
	}

	return string(s[pos+len(separator):])
}

// SubstringAfterLast Gets the substring after the last occurrence of a separator.
func SubstringAfterLast(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.LastIndex(s, separator)
	if pos == INDEX_NOT_FOUND || pos == len(s)-len(separator) {
		return EMPTY_STRING
	}

	return string(s[pos+len(separator):])
}

// SubstringBefore Gets the substring before the first occurrence of a separator
func SubstringBefore(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.Index(s, separator)
	if pos == INDEX_NOT_FOUND {
		return EMPTY_STRING
	}

	return string(s[:pos])
}

func SubstringBeforeLast(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.LastIndex(s, separator)
	if pos == INDEX_NOT_FOUND {
		return s
	}
	return string(s[:pos])
}

// fulfill string by repeat target count of byte
func Repeat(s byte, count int) string {
	ret := arrayx.CreateAndFill(count, s)
	return string(ret)
}

// StringToSlice base on unsafe package to convert string to []byte without copy action
// key point: copy string's Data and Len to slice's Data and Len, and append Cap value
func StringToSlice(value string) []byte {
	// create a new []byte
	var ret []byte

	// 把string的引用指向 ret的空间
	*(*string)(unsafe.Pointer(&ret)) = value

	// 设置slice的Cap值 ，用unsafe.Add操作，执行偏移操作 16个字节
	offset := uintptr(8) * 2
	*(*int)(unsafe.Add(unsafe.Pointer(&ret), offset)) = len(value)

	return ret
}

// SliceToString base on unsafe packge to convert []byte to string without copy action
// key point: copy slice's Data and Len to string's Data and Len.
func SliceToString(b []byte) string {
	if b == nil {
		return EMPTY_STRING
	}

	// just share Slice's Data and Len content
	return *(*string)(unsafe.Pointer(&b))
}

// IsEmpty return if s is empty
func IsEmpty(s string) bool {
	if s == EMPTY_STRING || len(s) == 0 {
		return true
	}
	return false
}

// IsBlank return if s is empty or blank string
// IsBlank("")  == true
// IsBlank(" ")  == true
// IsBlank(" a ")  == false
func IsBlank(s string) bool {
	if IsEmpty(s) || strings.TrimSpace(s) == EMPTY_STRING {
		return true
	}
	return false
}

// Wrap Wraps a String with another String.
func Wrap(s string, wrap string) string {
	if IsEmpty(s) || IsEmpty(wrap) {
		return s
	}
	b := make([]byte, len(s)+2*len(wrap))
	b1 := unsafex.StringToSlice(s)
	b2 := unsafex.StringToSlice(wrap)
	copy(b, b2)
	copy(b[len(wrap):], b1)
	copy(b[len(s)+len(wrap):], b2)

	return unsafex.SliceToString(b)
}
