/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-17 20:13:17
 */
package maths_test

import (
	"testing"

	"github.com/jhunters/goassist/maths"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSafeAdd(t *testing.T) {
	Convey("Test SafeAdd positive", t, func() {

		a := 1 << 62
		b := 1 << 62

		c, err := maths.SafeAdd(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeLessThan, 0)

	})

	Convey("Test SafeAdd negative", t, func() {

		a := -1 << 63
		b := -1 << 63

		c, err := maths.SafeAdd(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeGreaterThanOrEqualTo, 0)

	})
}

func TestSafeAddUnsigned(t *testing.T) {
	Convey("Test SafeAddUnsigned positive", t, func() {

		a := uint64(1 << 63)
		b := uint64(1 << 63)

		c, err := maths.SafeAddUnsigned(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldEqual, 0)

	})

}

func TestSafeSubstract(t *testing.T) {
	Convey("Test Substract positive", t, func() {

		a := 1 << 62
		b := -1 << 63

		c, err := maths.SafeSubstract(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeLessThan, 0)

	})

	Convey("Test SafeAdd negative", t, func() {

		a := -1 << 63
		b := 1 << 62

		c, err := maths.SafeSubstract(a, b)
		So(err, ShouldNotBeNil)
		So(c, ShouldBeGreaterThanOrEqualTo, 0)

	})
}
