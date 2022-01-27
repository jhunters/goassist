/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-26 11:45:37
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

	})
}

func TestAtoi(t *testing.T) {

	fmt.Println(^uint(0), conv.FormatInt(^uint(0), 2))
	fmt.Println(^uint(0) >> 63)

	v, err := strconv.Atoi("-100")
	fmt.Println(v, err)
}
