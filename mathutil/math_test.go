/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-17 20:13:17
 */
package mathutil_test

import (
	"testing"

	"github.com/jhunters/goassist/mathutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSafeAdd(t *testing.T) {
	Convey("Test SafeAdd positive", t, func() {

		a := 1 << 62
		b := 1 << 62

		c, err := mathutil.SafeAdd(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeLessThan, 0)

	})

	Convey("Test SafeAdd negative", t, func() {

		a := -1 << 63
		b := -1 << 63

		c, err := mathutil.SafeAdd(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeGreaterThanOrEqualTo, 0)

	})
}

func TestSafeAddUnsigned(t *testing.T) {
	Convey("Test SafeAddUnsigned positive", t, func() {

		a := uint64(1 << 63)
		b := uint64(1 << 63)

		c, err := mathutil.SafeAddUnsigned(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldEqual, 0)

	})

}

func TestSafeSubstract(t *testing.T) {
	Convey("Test Substract positive", t, func() {

		a := 1 << 62
		b := -1 << 63

		c, err := mathutil.SafeSubstract(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeLessThan, 0)

	})

	Convey("Test Substract negative", t, func() {

		a := -1 << 63
		b := 1 << 62

		c, err := mathutil.SafeSubstract(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeGreaterThanOrEqualTo, 0)

	})

	Convey("Test SubstractUnsigned negative", t, func() {

		var a uint64 = 1 << 62
		var b uint64 = 1 << 63

		c, err := mathutil.SafeSubstractUnsigned(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeGreaterThanOrEqualTo, 0)

	})
}

func TestMod(t *testing.T) {

	Convey("Test Mod positive", t, func() {
		v, err := mathutil.Mod(10, 3)
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 1)
	})

	Convey("Test Mod negative", t, func() {
		v, err := mathutil.Mod(-10, 3)
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 2)
	})

	Convey("Test Mod failed", t, func() {
		_, err := mathutil.Mod(-10, -3)
		So(err, ShouldNotBeNil)
	})

}

func TestMaxMin(t *testing.T) {

	Convey("TestMaxMin", t, func() {

		v := mathutil.Max(1, 0)
		So(v, ShouldEqual, 1)

		v = mathutil.Min(1, 0)
		So(v, ShouldEqual, 0)

	})

}

func TestAbs(t *testing.T) {

	Convey("TestAbs", t, func() {
		v := mathutil.Abs(1)
		So(v, ShouldEqual, 1)

		v = mathutil.Abs(-1)
		So(v, ShouldEqual, 1)

	})

}
