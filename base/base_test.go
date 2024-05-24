package base_test

import (
	"testing"

	"github.com/jhunters/goassist/base"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSafetyFunc(t *testing.T) {
	Convey("TestSafetyFunc", t, func() {
		v := base.SafetyFunc[string, string]("hello", func(s string) string {
			panic("err")
		})
		So(v, ShouldEqual, "")

		v = base.SafetyFunc[string, string]("hello", func(s string) string {
			return s + " world"
		})
		So(v, ShouldEqual, "hello world")
	})

}

func TestSafetyCall(t *testing.T) {
	Convey("TestSafetyCall", t, func() {
		defer func() {
			v := recover()
			So(v, ShouldBeNil)
		}()
		base.SafetyCall(func() {
			panic("TestSafetyCall panic")
		})
	})

}

func TestSafetyBiFunc(t *testing.T) {
	Convey("TestSafetyBiFunc", t, func() {
		r := base.SafetyBiFunc(3, 4, func(t, u int) int {
			return t + u
		})
		So(7, ShouldEqual, r)
		r = base.SafetyBiFunc(3, 4, func(t, u int) int {
			panic("error")
		})
		So(r, ShouldBeZeroValue)
	})

}

func TestSafetyConsumer(t *testing.T) {
	Convey("TestSafetyConsumer", t, func() {
		defer func() {
			v := recover()
			So(v, ShouldBeNil)
		}()
		base.SafetyConsumer(3, func(i int) {
			panic("throw error")
		})

	})

}
