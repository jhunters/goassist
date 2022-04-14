/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 14:12:42
 */

package str

import "strings"

// Reverse to reverse the string
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
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
