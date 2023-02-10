/*
 * Package conv provides convert utility apis
 */
package conv

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/jhunters/goassist/generic"
	"github.com/jhunters/goassist/stringutil"
)

const (
	Binary = 2
	Octal  = 10
	Hex    = 16
)

// Itoa formate integer and float value to string type
func Itoa[E generic.Integer | generic.Float](i E) string {
	return fmt.Sprintf("%v", i)
}

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatInt[E generic.Integer](i E, base int) string {
	if i > 0 && uint64(i) > math.MaxInt64 {
		return strconv.FormatUint(uint64(i), base)
	}
	return strconv.FormatInt(int64(i), base)
}

// ToPtr convert to pointer
func ToPtr[E any](t E) *E {
	return &t
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
	if stringutil.IsEmpty(str) {
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

// ParseInt to parse string to int
func ParseInt(str string) (int, error) {
	if !IsNumber(str) && !strings.Contains(str, ".") {
		return -1, fmt.Errorf("string '%s' is not a valid int number.", str)
	}

	chars := []byte(str)
	neg := 1
	start := 0
	if chars[0] == '-' {
		neg = -1
		start = 1
	} else if chars[0] == '+' {
		start = 1
	}

	if chars[start] == '0' { // leading 0, maybe base specified int eg 0x, 0b, 0o
		base := 0
		if chars[start+1] == 'x' || chars[start+1] == 'X' {
			base = 16
		} else if chars[start+1] == 'o' || chars[start+1] == 'O' {
			base = 8
		} else if chars[start+1] == 'b' || chars[start+1] == 'B' {
			base = 2
		}
		v, err := strconv.ParseInt(string(chars[start+2:]), base, 64)
		if err != nil {
			return -1, err
		}
		return int(v) * neg, nil
	}

	// if has e or E
	s := strings.ToLower(string(chars[start:]))
	if strings.Contains(s, "e") {
		splits := strings.Split(s, "e")
		v, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return -1, err
		}
		v2, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			return -1, err
		}
		pow := math.Pow10(int(v2))
		return int(float64(v) * pow), nil
	}

	v, err := strconv.ParseInt(str, 10, 64)
	return int(v), err
}

// ParseFloat to parse string to float64
func ParseFloat(str string) (float64, error) {
	if !IsNumber(str) {
		return -1, fmt.Errorf("string '%s' is not a valid int number.", str)
	}

	chars := []byte(str)
	neg := 1
	start := 0
	if chars[0] == '-' {
		neg = -1
		start = 1
	} else if chars[0] == '+' {
		start = 1
	}

	if chars[start] == '0' { // leading 0, maybe base specified int eg 0x, 0b, 0o
		base := 0
		if chars[start+1] == 'x' || chars[start+1] == 'X' {
			base = 16
		} else if chars[start+1] == 'o' || chars[start+1] == 'O' {
			base = 8
		} else if chars[start+1] == 'b' || chars[start+1] == 'B' {
			base = 2
		}
		v, err := strconv.ParseInt(string(chars[start+2:]), base, 64)
		if err != nil {
			return -1, err
		}
		return float64(int(v) * neg), nil
	}

	// if has e or E
	s := strings.ToLower(string(chars[start:]))
	if strings.Contains(s, "e") {
		splits := strings.Split(s, "e")
		v, err := strconv.ParseFloat(splits[0], 10)
		if err != nil {
			return -1, err
		}
		v2, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			return -1, err
		}
		pow := math.Pow10(int(v2))
		return v * pow, nil
	}

	return strconv.ParseFloat(str, 10)
}

// Append convert e to string and appends to dst
func Append[E any](dst []byte, e E) []byte {
	toAppend := fmt.Sprintf("%v", e)
	return append(dst, []byte(toAppend)...)
}

// AppendString  convert e to string and appends to str
func AppendString[E any](str string, e E) string {
	return string(Append([]byte(str), e))
}

// ParseBool It accepts 1, t, T, TRUE(captital ignore), 0, f, F, FALSE(captital ignore).
func ParseBool[E string | generic.Signed](e E) (bool, error) {
	val := reflect.ValueOf(e)

	switch val.Kind() {
	case reflect.String:
		s := val.String()
		s = strings.ToLower(s)
		return strconv.ParseBool(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := val.Int()
		if i == 0 {
			return false, nil
		} else if i == 1 {
			return true, nil
		}
		return false, fmt.Errorf("%d is not a bool int ", i)
	}
	// should not go here
	panic(fmt.Sprintf("invalid value type, %s", val.Kind().String()))
}
