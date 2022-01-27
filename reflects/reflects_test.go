/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-21 11:48:16
 */
package reflects_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jhunters/goassist/reflects"

	. "github.com/smartystreets/goconvey/convey"
)

type I interface {
	GetName() string
}

type Value struct {
	Name string
}

func (v *Value) GetName() string {
	return v.Name
}

type Student struct {
	Name string
}

func (s *Student) Greet(greeting I) string {
	return fmt.Sprintf("%s : %s", greeting.GetName(), s.Name)
}

func (s *Student) GreetWithOption(greeting I, option string) string {
	return fmt.Sprintf("%s : %s [%s]", greeting.GetName(), s.Name, option)
}

func TestValueOf(t *testing.T) {

	Convey("TestValueOf", t, func() {
		i := int64(100)
		iPtr := &i

		v, isPtr := reflects.ValueOf(iPtr)
		So(isPtr, ShouldBeTrue)
		So(v, ShouldResemble, reflect.ValueOf(iPtr).Elem())
		fmt.Println(v.String())
	})

}

func TestTypeOf(t *testing.T) {
	Convey("TestTypeOf", t, func() {
		i := int64(100)
		iPtr := &i

		v, isPtr := reflects.TypeOf(iPtr)
		So(isPtr, ShouldBeTrue)
		So(v, ShouldResemble, reflect.TypeOf(iPtr).Elem())
		fmt.Println(v.String())
	})
}

func TestCallMethod(t *testing.T) {

	Convey("TestTypeOf", t, func() {
		s := &Student{"matthew"}
		var i I = &Value{"Hello"}
		result, err := reflects.CallMethodName(s, "Greet", i)
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.Greet(i))
	})

	Convey("TestTypeOf with two parameter", t, func() {
		s := &Student{"matthew"}
		var i I = &Value{"Hello"}
		result, err := reflects.CallMethodName(s, "GreetWithOption", i, "name")
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.GreetWithOption(i, "name"))
	})

}
