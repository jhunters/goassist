/*
 */
package conv_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/jhunters/goassist/conv"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToString(t *testing.T) {

	Convey("Test ToString", t, func() {
		Convey("Test ToString for integer type ", func() {
			expect := "10"
			i8 := int8(10)
			result := conv.Itoa(i8)
			So(result, ShouldEqual, expect)

			ui8 := uint8(10)
			result = conv.Itoa(ui8)
			So(result, ShouldEqual, expect)

			i64 := int64(10)
			result = conv.Itoa(i64)
			So(result, ShouldEqual, expect)

			i := int(10)
			result = conv.Itoa(i)
			So(result, ShouldEqual, expect)

		})

		Convey("Test ToString for float type ", func() {
			expect := "10.01"
			f32 := float32(10.01)
			result := conv.Itoa(f32)
			So(result, ShouldEqual, expect)

			f64 := float64(10.01)
			result = conv.Itoa(f64)
			So(result, ShouldEqual, expect)

		})
	})

}

func TestFormatInt(t *testing.T) {

	Convey("Test FormatInt", t, func() {

		format := conv.FormatInt(math.MaxInt64, 2)
		So(format, ShouldEqual, strconv.FormatInt(math.MaxInt64, 2))

		// negative
		format = conv.FormatInt(math.MinInt64, 2)
		So(format, ShouldEqual, strconv.FormatInt(math.MinInt64, 2))

		var bigUnit64 uint64 = math.MaxInt64 + 100
		format = conv.FormatInt(bigUnit64, 2)
		So(format, ShouldEqual, strconv.FormatUint(bigUnit64, 2))

		var smartUnit64 uint64 = 100
		format = conv.FormatInt(smartUnit64, 2)
		So(format, ShouldEqual, strconv.FormatUint(smartUnit64, 2))

		var sint int8 = 10
		format = conv.FormatInt(sint, 10)
		So(format, ShouldEqual, strconv.FormatInt(int64(sint), 10))
	})
}

func TestAtoi(t *testing.T) {

	fmt.Println(^uint(0), conv.FormatInt(^uint(0), 2))
	fmt.Println(^uint(0) >> 63)

	v, err := strconv.Atoi("-100")
	fmt.Println(v, err)
}

func TestToPrt(t *testing.T) {
	Convey("TestToPrt", t, func() {
		iPtr := conv.ToPtr(100)
		So(*iPtr, ShouldEqual, 100)
	})

}

func TestIsNumber(t *testing.T) {
	Convey("TestIsNumber", t, func() {
		Convey("hex", func() {
			isNumber := conv.IsNumber("0x12")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("0xac")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("0x")
			So(isNumber, ShouldBeFalse)

			isNumber = conv.IsNumber("0x093g")
			So(isNumber, ShouldBeFalse)
		})
		Convey("octal", func() {
			isNumber := conv.IsNumber("0o10")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("0O17")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("0o18")
			So(isNumber, ShouldBeFalse)

			isNumber = conv.IsNumber("0O093g")
			So(isNumber, ShouldBeFalse)
		})
		Convey("binary", func() {
			isNumber := conv.IsNumber("0b10")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("0B11")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("0b18")
			So(isNumber, ShouldBeFalse)

			isNumber = conv.IsNumber("0B093g")
			So(isNumber, ShouldBeFalse)
		})
		Convey("common", func() {
			isNumber := conv.IsNumber("-12.11")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("19.1")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("-12.1.1")
			So(isNumber, ShouldBeFalse)

			isNumber = conv.IsNumber("12e1")
			So(isNumber, ShouldBeTrue)

			isNumber = conv.IsNumber("12e-9")
			So(isNumber, ShouldBeTrue)
		})

	})

}

func TestParseInt(t *testing.T) {
	Convey("TestParseInt", t, func() {
		// hex string
		v, err := conv.ParseInt("-0x19ac")
		So(v, ShouldEqual, -6572)
		So(err, ShouldBeNil)

		v, err = conv.ParseInt("0X19ac")
		So(v, ShouldEqual, 6572)
		So(err, ShouldBeNil)

		// octal string
		v, err = conv.ParseInt("0o1312")
		So(v, ShouldEqual, 714)
		So(err, ShouldBeNil)

		v, err = conv.ParseInt("-0O1312")
		So(v, ShouldEqual, -714)
		So(err, ShouldBeNil)

		// binary string
		v, err = conv.ParseInt("0b1011")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)

		v, err = conv.ParseInt("-0B1011")
		So(v, ShouldEqual, -11)
		So(err, ShouldBeNil)

		// e
		v, err = conv.ParseInt("11e10")
		So(v, ShouldEqual, 110000000000)
		So(err, ShouldBeNil)

		v, err = conv.ParseInt("110e-1")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)

		v, err = conv.ParseInt("11e-1")
		So(v, ShouldEqual, 1)
		So(err, ShouldBeNil)

		// invalid int number
		v, err = conv.ParseInt("11ee-1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)

		v, err = conv.ParseInt("--11e-1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)
		v, err = conv.ParseInt("-+11e-1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)
		v, err = conv.ParseInt("1.1e1")
		So(v, ShouldEqual, -1)
		So(err, ShouldNotBeNil)
	})

}
func ExampleParseInt() {
	// hex string
	v, err := conv.ParseInt("-0x19ac")
	fmt.Println(v, err)

	v, err = conv.ParseInt("0X19ac")
	fmt.Println(v, err)

	// octal string
	v, err = conv.ParseInt("0o1312")
	fmt.Println(v, err)
	v, err = conv.ParseInt("-0O1312")
	fmt.Println(v, err)

	// binary string
	v, err = conv.ParseInt("0b1011")
	fmt.Println(v, err)
	v, err = conv.ParseInt("-0B1011")
	fmt.Println(v, err)

	// e
	v, err = conv.ParseInt("11e10")
	fmt.Println(v, err)

	v, err = conv.ParseInt("110e-1")
	fmt.Println(v, err)

	v, err = conv.ParseInt("11e-1")
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
		v, err := conv.ParseFloat("-0x19ac")
		So(v, ShouldEqual, -6572)
		So(err, ShouldBeNil)

		v, err = conv.ParseFloat("0X19ac")
		So(v, ShouldEqual, 6572)
		So(err, ShouldBeNil)

		// octal string
		v, err = conv.ParseFloat("0o1312")
		So(v, ShouldEqual, 714)
		So(err, ShouldBeNil)
		v, err = conv.ParseFloat("-0O1312")
		So(v, ShouldEqual, -714)
		So(err, ShouldBeNil)

		// binary string
		v, err = conv.ParseFloat("0b1011")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)
		v, err = conv.ParseFloat("-0B1011")
		So(v, ShouldEqual, -11)
		So(err, ShouldBeNil)

		// e
		v, err = conv.ParseFloat("11e10")
		So(v, ShouldEqual, 110000000000)
		So(err, ShouldBeNil)

		v, err = conv.ParseFloat("110e-1")
		So(v, ShouldEqual, 11)
		So(err, ShouldBeNil)

		v, err = conv.ParseFloat("11e-1")
		So(v, ShouldEqual, 1.1)
		So(err, ShouldBeNil)
	})

}

func ExampleParseFloat() {
	// hex string
	v, err := conv.ParseFloat("-0x19ac")
	fmt.Println(v, err)

	v, err = conv.ParseFloat("0X19ac")
	fmt.Println(v, err)

	// octal string
	v, err = conv.ParseFloat("0o1312")
	fmt.Println(v, err)
	v, err = conv.ParseFloat("-0O1312")
	fmt.Println(v, err)

	// binary string
	v, err = conv.ParseFloat("0b1011")
	fmt.Println(v, err)
	v, err = conv.ParseFloat("-0B1011")
	fmt.Println(v, err)

	// e
	v, err = conv.ParseFloat("11e10")
	fmt.Println(v, err)

	v, err = conv.ParseFloat("110e-1")
	fmt.Println(v, err)

	v, err = conv.ParseFloat("11e-1")
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

func TestAppend(t *testing.T) {
	Convey("TestAppend", t, func() {
		b10 := []byte("int:")
		b10 = conv.Append(b10, -42)
		So(string(b10), ShouldEqual, "int:-42")

		// append bool
		b := []byte("bool:")
		b = conv.Append(b, true)
		So(string(b), ShouldEqual, "bool:true")

		// append float
		b32 := []byte("float32:")
		b32 = conv.Append(b32, 3.1415926535)
		So(string(b32), ShouldEqual, "float32:3.1415926535")

		b64 := []byte("float64:")
		b64 = conv.Append(b64, 3.1415926535)
		So(string(b64), ShouldEqual, "float64:3.1415926535")

		// append quote
		b = []byte("quote:")
		b = conv.Append(b, `"Fran & Freddie's Diner"`)
		So(string(b), ShouldEqual, `quote:"Fran & Freddie's Diner"`)
	})

}
func TestAppendString(t *testing.T) {
	Convey("TestAppendString", t, func() {
		b10 := "int:"
		b10 = conv.AppendString(b10, -42)
		So(string(b10), ShouldEqual, "int:-42")

		// append bool
		b := "bool:"
		b = conv.AppendString(b, true)
		So(string(b), ShouldEqual, "bool:true")

		// append float
		b32 := "float32:"
		b32 = conv.AppendString(b32, 3.1415926535)
		So(string(b32), ShouldEqual, "float32:3.1415926535")

		b64 := "float64:"
		b64 = conv.AppendString(b64, 3.1415926535)
		So(string(b64), ShouldEqual, "float64:3.1415926535")

		// append quote
		b = "quote:"
		b = conv.AppendString(b, `"Fran & Freddie's Diner"`)
		So(string(b), ShouldEqual, `quote:"Fran & Freddie's Diner"`)
	})

}

func ExampleAppend() {
	// append int
	b10 := []byte("int:")
	b10 = conv.Append(b10, -42)
	fmt.Println(string(b10))

	// append bool
	b := []byte("bool:")
	b = conv.Append(b, true)
	fmt.Println(string(b))

	// append float
	b32 := []byte("float32:")
	b32 = conv.Append(b32, 3.1415926535)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = conv.Append(b64, 3.1415926535)
	fmt.Println(string(b64))

	// append quote
	b = []byte("quote:")
	b = conv.Append(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	// Output:
	// int:-42
	// bool:true
	// float32:3.1415926535
	// float64:3.1415926535
	// quote:"Fran & Freddie's Diner"

}

func TestParseBool(t *testing.T) {

	Convey("TestParseBool", t, func() {

		// case for parsed return true
		b, err := conv.ParseBool(1)
		So(b, ShouldBeTrue)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("1")
		So(b, ShouldBeTrue)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("true")
		So(b, ShouldBeTrue)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("True")
		So(b, ShouldBeTrue)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("TRUE")
		So(b, ShouldBeTrue)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("tRUE")
		So(b, ShouldBeTrue)
		So(err, ShouldBeNil)

		// case for parsed return false
		b, err = conv.ParseBool(0)
		So(b, ShouldBeFalse)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("0")
		So(b, ShouldBeFalse)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("false")
		So(b, ShouldBeFalse)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("False")
		So(b, ShouldBeFalse)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("FALSE")
		So(b, ShouldBeFalse)
		So(err, ShouldBeNil)

		b, err = conv.ParseBool("FaLSe")
		So(b, ShouldBeFalse)
		So(err, ShouldBeNil)

		// case for parsed failed
		b, err = conv.ParseBool(2)
		So(b, ShouldBeFalse)
		So(err, ShouldNotBeNil)

		b, err = conv.ParseBool("100")
		So(b, ShouldBeFalse)
		So(err, ShouldNotBeNil)

		b, err = conv.ParseBool("abc")
		So(b, ShouldBeFalse)
		So(err, ShouldNotBeNil)
	})

}

func TestCitoa(t *testing.T) {

	Convey("TestCitoa", t, func() {

		numStr, err := conv.CItoa("一千万零一百一十五")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "10000115")

		numStr, err = conv.CItoa("一亿零五百万零一十五")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "105000015")

		numStr, err = conv.CItoa("一千零一百一十五")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "1115")

		numStr, err = conv.CItoa("一千一十一")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "1011")

		numStr, err = conv.CItoa("一千零十一")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "1011")

		numStr, err = conv.CItoa("十一")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "11")

		numStr, err = conv.CItoa("一万零百")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "10100")

		numStr, err = conv.CItoa("一百十一")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "111")

		numStr, err = conv.CItoa("五百万亿零十一")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "500000000000011")

		numStr, err = conv.CItoa("八")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "8")

		numStr, err = conv.CItoa("零")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "0")

		numStr, err = conv.CItoa("一亿二千三百五十万零一十五")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "123500015")

		numStr, err = conv.CItoa("一千万零一百亿")
		So(err, ShouldBeNil)
		So(numStr, ShouldEqual, "1000010000000000")

	})

	Convey("TestCitoa failed", t, func() {
		numStr, err := conv.CItoa("一q千万零一百一十五")
		So(err, ShouldNotBeNil)
		So(numStr, ShouldBeEmpty)

	})

}

func ExampleCItoa() {
	numStr, err := conv.CItoa("一千万零一百一十五")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("五百万亿零十一")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("一亿二千三百五十万零一十五")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("一千万零一百亿")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("八")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("零")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("十五")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("二亿五万万")
	if err == nil {
		fmt.Println(numStr)
	}

	numStr, err = conv.CItoa("十八亿五万万")
	if err == nil {
		fmt.Println(numStr)
	}

	// Output:
	// 10000115
	// 500000000000011
	// 123500015
	// 1000010000000000
	// 8
	// 0
	// 15
	// 700000000
	// 2300000000
}
