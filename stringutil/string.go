/*
 * Package stringutil to provides utility api for string operation
 */
package stringutil

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/container/set"
	"github.com/jhunters/goassist/unsafex"
)

const (
	INDEX_NOT_FOUND = -1
	EMPTY_STRING    = ""
)

// Abbreviates a String using a given replacement marker
func Abbreviate(str, abbrevMarker string, offset, maxWidth int) (string, error) {
	// 如果输入字符串和缩写标记都为空，则直接返回原字符串
	if IsEmpty(str) && IsEmpty(abbrevMarker) {
		return str, nil
	} else if !IsEmpty(str) && abbrevMarker == "" && maxWidth > 0 {
		// 如果输入字符串不为空，缩写标记为空，且最大宽度大于0，则返回输入字符串的前maxWidth个字符
		return SubString(str, 0, maxWidth), nil
	} else if IsEmpty(str) || IsEmpty(abbrevMarker) {
		// 如果输入字符串或缩写标记为空，则直接返回原字符串
		return str, nil
	}

	abbrevMarkerLength := len(abbrevMarker)
	minAbbrevWidth := abbrevMarkerLength + 1
	minAbbrevWidthOffset := abbrevMarkerLength + abbrevMarkerLength + 1

	// 如果最大宽度小于最小缩写宽度，则返回错误提示信息
	if maxWidth < minAbbrevWidth {
		return str, fmt.Errorf("minimum abbreviation width is %d", minAbbrevWidth)
	}
	l := len(str)
	// 如果字符串长度小于等于最大宽度，则直接返回原字符串
	if l <= maxWidth {
		return str, nil
	}
	// 如果偏移量大于字符串长度，则将偏移量设为字符串长度
	if offset > l {
		offset = l
	}
	// 如果字符串长度减去偏移量小于等于最大宽度减去缩写标记长度，则将偏移量设为字符串长度减去(最大宽度减去缩写标记长度)
	if l-offset < maxWidth-abbrevMarkerLength {
		offset = l - (maxWidth - abbrevMarkerLength)
	}
	// 如果偏移量小于等于缩写标记长度加1，则返回从字符串开头到最大宽度的字符再加上缩写标记
	if offset <= abbrevMarkerLength+1 {
		return SubString(str, 0, maxWidth-abbrevMarkerLength) + abbrevMarker, nil
	}
	// 如果最大宽度小于最小缩写宽度加上偏移量，则返回错误提示信息
	if maxWidth < minAbbrevWidthOffset {
		return str, fmt.Errorf("minimum abbreviation width with offset is %d", minAbbrevWidthOffset)
	}
	// 如果偏移量加上最大宽度减去缩写标记长度小于字符串长度，则返回从偏移量位置到字符串末尾的字符再加上缩写标记
	if offset+maxWidth-abbrevMarkerLength < l {
		ns, err := Abbreviate(SubString(str, offset, -1), abbrevMarker, 0, maxWidth-abbrevMarkerLength)
		if err != nil {
			return str, err
		}
		return abbrevMarker + ns, nil
	}
	// 否则返回从字符串末尾到最大宽度的字符再加上缩写标记
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

// SubstringBeforeLast  Gets the substring before the last occurrence of a separator
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

// SubstringMatch to returns whether the given string matches the given substring
func SubstringMatch(s string, index int, sub string) bool {
	if index+len(sub) > len(s) {
		return false
	}

	for i := 0; i < len(sub); i++ {
		if s[index+i] != sub[i] {
			return false
		}
	}
	return true
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

// IndexFromOffset to return index of sub from Index
func IndexFromOffset(s, sub string, fromIndex int) int {
	if fromIndex < 0 || fromIndex+len(sub) >= len(s) {
		return -1
	}
	left := SubString(s, fromIndex, len(s))
	index := strings.Index(left, sub)
	if index != -1 {
		index += fromIndex
	}
	return index
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

// Expand to solve place holder value with prefix and suffix and replace by 'fn' call back function
func Expand(s, prefix, suffix string, fn base.Func[string, string]) (string, error) {
	if IsBlank(prefix) || IsBlank(suffix) {
		return EMPTY_STRING, fmt.Errorf("invalid prefix or suffix")
	}

	startIndex := strings.Index(s, prefix)
	if startIndex == -1 {
		return s, nil
	}

	visitedPlaceholders := set.NewSet[string]()

	for startIndex != -1 {
		endIndex := findPlaceholderEndIndex(s, startIndex, prefix, suffix)
		if endIndex != -1 {
			placeholder := SubString(s, startIndex+len(prefix), endIndex)
			originalPlaceholder := placeholder

			if !visitedPlaceholders.Add(originalPlaceholder) {
				return EMPTY_STRING, fmt.Errorf("circular placeholder reference %s in property definitions", originalPlaceholder)
			}

			// Recursive invocation, parsing placeholders contained in the placeholder key.
			placeholder, err := Expand(placeholder, prefix, suffix, fn)
			if err != nil {
				return EMPTY_STRING, err
			}

			// Now obtain the value for the fully resolved key...
			propVal := fn(placeholder)
			if !IsBlank(propVal) {
				propVal, err := Expand(propVal, prefix, suffix, fn)
				if err != nil {
					return EMPTY_STRING, err
				}
				s = ReplaceByOffset(s, startIndex, endIndex+len(suffix), propVal)
				offset := startIndex + len(propVal)
				startIndex = IndexFromOffset(s, prefix, offset)
			} else {
				offset := endIndex + len(prefix)
				startIndex = IndexFromOffset(s, prefix, offset)
			}
			visitedPlaceholders.Remove(originalPlaceholder)
		} else {
			startIndex = -1
		}

	}

	return s, nil
}

// ReplaceByOffset to replace sub string by offset begin and end index
func ReplaceByOffset(s string, begin, end int, replace string) string {
	sz := len(s)
	if begin > sz || end > sz {
		return s
	}
	if begin > end {
		return s
	}

	bb := []byte(s)
	b1 := bb[:begin]
	ret := string(b1) + replace
	b2 := bb[end:]
	return ret + string(b2)
}

func findPlaceholderEndIndex(buf string, startIndex int, prefix, suffix string) int {
	index := startIndex + len(prefix)
	withinNestedPlaceholder := 0
	for index < len(buf) {
		if SubstringMatch(buf, index, suffix) {
			if withinNestedPlaceholder > 0 {
				withinNestedPlaceholder--
				index = index + len(suffix)
			} else {
				return index
			}
		} else if SubstringMatch(buf, index, prefix) {
			withinNestedPlaceholder++
			index = index + len(prefix)
		} else {
			index++
		}
	}

	return -1
}
