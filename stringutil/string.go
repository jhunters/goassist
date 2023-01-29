/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:12:42
 */

package stringutil

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/unsafex"
)

const (
	INDEX_NOT_FOUND = -1
	EMPTY_STRING    = ""
)

// Abbreviates a String using a given replacement marker
func Abbreviate(str, abbrevMarker string, offset, maxWidth int) (string, error) {
	if IsEmpty(str) && IsEmpty(abbrevMarker) {
		return str, nil
	} else if !IsEmpty(str) && abbrevMarker == "" && maxWidth > 0 {
		return SubString(str, 0, maxWidth), nil
	} else if IsEmpty(str) || IsEmpty(abbrevMarker) {
		return str, nil
	}

	abbrevMarkerLength := len(abbrevMarker)
	minAbbrevWidth := abbrevMarkerLength + 1
	minAbbrevWidthOffset := abbrevMarkerLength + abbrevMarkerLength + 1

	if maxWidth < minAbbrevWidth {
		return str, fmt.Errorf("Minimum abbreviation width is %d", minAbbrevWidth)
	}
	l := len(str)
	if l <= maxWidth {
		return str, nil
	}
	if offset > l {
		offset = l
	}
	if l-offset < maxWidth-abbrevMarkerLength {
		offset = l - (maxWidth - abbrevMarkerLength)
	}
	if offset <= abbrevMarkerLength+1 {
		return SubString(str, 0, maxWidth-abbrevMarkerLength) + abbrevMarker, nil
	}
	if maxWidth < minAbbrevWidthOffset {
		return str, fmt.Errorf("Minimum abbreviation width with offset is %d", minAbbrevWidthOffset)
	}
	if offset+maxWidth-abbrevMarkerLength < l {
		ns, err := Abbreviate(SubString(str, offset, -1), abbrevMarker, 0, maxWidth-abbrevMarkerLength)
		if err != nil {
			return str, err
		}
		return abbrevMarker + ns, nil
	}
	return abbrevMarker + SubString(str, l-(maxWidth-abbrevMarkerLength), -1), nil
}

// AbbreviateMiddle a String to the length passed, replacing the middle characters with the supplied replacement String.
func AbbreviateMiddle(str, middle string, length int) string {
	if IsEmpty(str) || IsEmpty(middle) {
		return str
	}

	if length >= len(str) || length < len(middle)+2 {
		return str
	}

	targetSting := length - len(middle)
	startOffset := targetSting/2 + targetSting%2
	endOffset := len(str) - targetSting/2

	return SubString(str, 0, startOffset) +
		middle +
		SubString(str, endOffset, -1)
}

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

// SubString a string that is a substring of this string.
func SubString(s string, beginIndex, endIndex int) string {
	if beginIndex < 0 {
		return s
	}
	if endIndex > len(s) || endIndex < 0 {
		return s[beginIndex:]
	}

	return s[beginIndex:endIndex]
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
	ret := arrayutil.CreateAndFill(count, s)
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

// IsNumber return string 'str' is a validate number
// stringutil.IsNumber("0x12") = true
// stringutil.IsNumber("0x") = false
// stringutil.IsNumber("0o10") = true
// stringutil.IsNumber("0o18") = false
// stringutil.IsNumber("0b10") = true
// stringutil.IsNumber("0B093g") = false
// stringutil.IsNumber("-12.11") = true
// stringutil.IsNumber("12e-9") = true
// stringutil.IsNumber("19.1") = true
// stringutil.IsNumber("-12.1.1") = false
func IsNumber(str string) bool {
	if IsEmpty(str) {
		return false
	}

	chars := []byte(str)
	sz := len(chars)
	hasExp := false
	hasDecPoint := false
	allowSigns := false
	foundDigit := false

	// deal with any possible sign up front
	start := 0
	if chars[0] == '-' || chars[0] == '+' {
		start = 1
	}

	if sz > start+1 && chars[start] == '0' && !strings.Contains(str, ".") { // leading 0, skip if is a decimal number
		if chars[start+1] == 'x' || chars[start+1] == 'X' { // leading 0x/0X for hex number
			i := start + 2
			if i == sz {
				return false // str == "0x"
			}

			// checking hex (it can't be anything else)
			for ; i < len(chars); i++ {
				if (chars[i] < '0' || chars[i] > '9') && (chars[i] < 'a' || chars[i] > 'f') && (chars[i] < 'A' || chars[i] > 'F') {
					return false
				}
			}
			return true
		} else if chars[start+1] == 'o' || chars[start+1] == 'O' { // leading 0o/0O for octal number
			i := start + 2
			for ; i < len(chars); i++ {
				if chars[i] < '0' || chars[i] > '7' {
					return false
				}
			}
			return true
		} else if chars[start+1] == 'b' || chars[start+1] == 'B' { // leading 0o/0O for binary number
			i := start + 2
			for ; i < len(chars); i++ {
				if chars[i] < '0' || chars[i] > '1' {
					return false
				}
			}
			return true
		}
	}

	sz-- // don't want to loop to the last char, check it afterwords
	// for type qualifiers
	i := start
	// loop to the next to last char or to the last char if we need another digit to
	// make a valid number (e.g. chars[0..5] = "1234E")
	for i < sz || i < sz+1 && allowSigns && !foundDigit {
		if chars[i] >= '0' && chars[i] <= '9' {
			foundDigit = true
			allowSigns = false

		} else if chars[i] == '.' {
			if hasDecPoint || hasExp {
				// two decimal points or dec in exponent
				return false
			}
			hasDecPoint = true
		} else if chars[i] == 'e' || chars[i] == 'E' {
			// we've already taken care of hex.
			if hasExp {
				// two E's
				return false
			}
			if !foundDigit {
				return false
			}
			hasExp = true
			allowSigns = true
		} else if chars[i] == '+' || chars[i] == '-' {
			if !allowSigns {
				return false
			}
			allowSigns = false
			foundDigit = false // we need a digit after the E
		} else {
			return false
		}
		i++
	}

	if i < len(chars) {
		if chars[i] >= '0' && chars[i] <= '9' {
			// no type qualifier, OK
			return true
		}
		if chars[i] == 'e' || chars[i] == 'E' {
			// can't have an E at the last byte
			return false
		}
		if chars[i] == '.' {
			if hasDecPoint || hasExp {
				// two decimal points or dec in exponent
				return false
			}
			// single trailing decimal point after non-exponent is ok
			return foundDigit
		}
		if !allowSigns {
			return foundDigit
		}
		// last character is illegal
		return false
	}
	// allowSigns is true iff the val ends in 'E'
	// found digit it to make sure weird stuff like '.' and '1E-' doesn't pass
	return !allowSigns && foundDigit

}
