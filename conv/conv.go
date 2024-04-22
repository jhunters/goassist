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

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/container/listx"
	"github.com/jhunters/goassist/generic"
	"github.com/jhunters/goassist/stringutil"
)

const (
	Binary = 2
	Octal  = 10
	Hex    = 16

	// fixed unit value to  '亿' or '万'
	FixedUnitValue = 4

	CH_ZERO   = '零'
	CH_TEN    = '十'
	CH_ONE    = '一'
	RUNE_ONE  = 1
	RUNE_ZERO = 0

	S_ZERO = '0'

	CH_WAN_BIT = 4 // 4 bit for '万'
	CH_YI_BIT  = 8 // 8 bit for '亿'
)

var (
	Chinese_Uint_Number map[rune]int
	Chinese_Number      map[rune]byte

	Number_Chinese_Uint map[int]rune
	Number_Chinese      map[rune]rune

	Chinese_Unit []rune
)

type CNum struct {
	num       byte
	unit      int
	fixedunit int
}

func init() {

	// define chinese number  unit=size
	Chinese_Uint_Number = map[rune]int{}
	Chinese_Uint_Number['亿'] = 8
	Chinese_Uint_Number['万'] = 4
	Chinese_Uint_Number['千'] = 3
	Chinese_Uint_Number['百'] = 2
	Chinese_Uint_Number[CH_TEN] = 1

	Number_Chinese_Uint = map[int]rune{}
	Number_Chinese_Uint[9] = '亿'
	Number_Chinese_Uint[5] = '万'
	Number_Chinese_Uint[3] = '千'
	Number_Chinese_Uint[2] = '百'
	Number_Chinese_Uint[1] = CH_TEN

	Chinese_Number = map[rune]byte{}
	Chinese_Number[CH_ONE] = RUNE_ONE
	Chinese_Number['二'] = 2
	Chinese_Number['三'] = 3
	Chinese_Number['四'] = 4
	Chinese_Number['五'] = 5
	Chinese_Number['六'] = 6
	Chinese_Number['七'] = 7
	Chinese_Number['七'] = 7
	Chinese_Number['八'] = 8
	Chinese_Number['九'] = 9
	Chinese_Number[CH_ZERO] = RUNE_ZERO

	Number_Chinese = map[rune]rune{}
	Number_Chinese['1'] = CH_ONE
	Number_Chinese['2'] = '二'
	Number_Chinese['3'] = '三'
	Number_Chinese['4'] = '四'
	Number_Chinese['5'] = '五'
	Number_Chinese['6'] = '六'
	Number_Chinese['7'] = '七'
	Number_Chinese['8'] = '八'
	Number_Chinese['9'] = '九'
	Number_Chinese['0'] = CH_ZERO

}

// CAtoi 函数将传入的数字 uint 转换为中文大写表示，并返回转换后的字符串和 error。
// 如果 num 为 0，则直接返回 "零" 和 nil。
func CAtoi(num uint) (string, error) {
	if num == 0 {
		return string(CH_ZERO), nil
	}
	numStr := fmt.Sprintf("%d", num)
	nums := []rune(numStr)

	ret := ""
	cbit := 0
	cunitbit := 0

	isZero := false
	isFirstNoneZero := false
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]

		if n == S_ZERO {
			isZero = true
			if cbit == CH_WAN_BIT { // 进万 （或亿)
				ret = string(Number_Chinese_Uint[(cunitbit%2+1)*4+1]) + ret
				isFirstNoneZero = false
			}
		} else {
			if isFirstNoneZero && isZero {
				if cbit == 0 {
					ret = string(CH_TEN) + ret
				} else {
					ret = string(CH_ZERO) + ret
				}
			}

			isZero = false
			isFirstNoneZero = true

			if v, ok := Number_Chinese_Uint[cbit]; ok && !isZero {
				ret = string(v) + ret
			}

			if cbit == CH_WAN_BIT { // 进万 （或亿)
				ret = string(Number_Chinese_Uint[(cunitbit%2+1)*4+1]) + ret
			}
			ret = string(Number_Chinese[n]) + ret
		}
		if cbit == CH_WAN_BIT { // 进万 （或亿)
			cbit = 0
			cunitbit++
		}

		cbit++
	}

	return ret, nil
}

// CItoa convert chinese number to asc number string
// CItoa 函数将中文数字转换为阿拉伯数字，并返回结果和可能存在的错误。
// 参数chinesenum：需要转换的中文数字。
// 返回值string：转换后的阿拉伯数字。
// 返回值error：转换过程中可能遇到的错误。
func CItoa(chinesenum string) (string, error) {
	if stringutil.IsBlank(chinesenum) {
		return stringutil.EMPTY_STRING, fmt.Errorf("invalid chinese number. %s", chinesenum)
	}

	ls := listx.NewList[CNum]()
	nums := []rune(chinesenum)
	unit := 1
	fixedunit := 0
	biggestUnit := 0
	hasNumBeforeUnit := false
	hasWanBefore := false
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if v, ok := Chinese_Number[num]; ok {
			cNum := CNum{v, unit, fixedunit}
			ls.PushFront(cNum)
			hasNumBeforeUnit = true
		} else if v, ok := Chinese_Uint_Number[num]; ok {
			if v >= FixedUnitValue { // if just like '万、亿' unit
				if biggestUnit < v {
					biggestUnit = v
					fixedunit = v // just ajust to fixed unit
				} else { // 未超过最大单位
					// base add num unit
					if v == CH_YI_BIT && hasWanBefore {
						v -= CH_WAN_BIT
					}
					fixedunit += v
				}
				hasWanBefore = v == CH_WAN_BIT
				unit = 1
			} else {
				unit = v + 1
				// check if miss num before unit, so we check next one
				i2 := i - 1
				if i2 >= 0 {
					if v, ok := Chinese_Uint_Number[nums[i2]]; ok {
						if v < FixedUnitValue {
							cNum := CNum{RUNE_ONE, unit, fixedunit} // add one as default unit value
							ls.PushFront(cNum)
						}
					} else if nums[i2] == CH_ZERO {
						cNum := CNum{RUNE_ONE, unit, fixedunit} // add one as default unit value
						ls.PushFront(cNum)
					}
				}

			}
			hasNumBeforeUnit = false
		} else {
			return stringutil.EMPTY_STRING, fmt.Errorf("invalid chinese number. %s", chinesenum)
		}
	}

	size := unit + fixedunit
	numbers := make([]byte, size)

	// fix if has no num before unit, just like as '十' or '十五'
	if !hasNumBeforeUnit {
		ls.PushFront(CNum{RUNE_ONE, unit, fixedunit})
	}

	ls.Range(func(c CNum) bool {
		if c.num != RUNE_ZERO {
			offset := size - c.unit - c.fixedunit
			numbers[offset] = c.num + numbers[offset]
		}

		return true
	})

	// refix every bit if needs carry bit
	var maxBit byte = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 9 {
			if i == 0 {
				maxBit = numbers[i] / 10
			} else {
				numbers[i-1] = numbers[i-1] + 1
			}
			numbers[i] = numbers[i] % 10
		}
	}
	if maxBit > 0 {
		numbers = append([]byte{maxBit}, numbers...)
	}

	return arrayutil.Join(numbers, stringutil.EMPTY_STRING), nil
}

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
		return -1, fmt.Errorf("string '%s' is not a valid int number", str)
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
		return -1, fmt.Errorf("string '%s' is not a valid int number", str)
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
		v, err := strconv.ParseFloat(splits[0], 64)
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

	return strconv.ParseFloat(str, 64)
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
