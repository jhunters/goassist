/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-21 11:48:16
 */
package reflectutil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jhunters/goassist/reflectutil"
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

func (s *Student) GetName() string {
	return s.Name
}

func (s Student) GetName2() string {
	return s.Name
}

func (s *Student) ChangeName(name string) {
	s.Name = name
}

func (s Student) ChangeName2(name string) {
	s.Name = name
	fmt.Println("print at ChangeName2", s.Name)
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

		v, isPtr := reflectutil.ValueOf(iPtr)
		So(isPtr, ShouldBeTrue)
		So(v, ShouldResemble, reflect.ValueOf(iPtr).Elem())
		fmt.Println(v.String())
	})

}

func TestTypeOf(t *testing.T) {
	Convey("TestTypeOf", t, func() {
		i := int64(100)
		iPtr := &i

		v, isPtr := reflectutil.TypeOf(iPtr)
		So(isPtr, ShouldBeTrue)
		So(v, ShouldResemble, reflect.TypeOf(iPtr).Elem())
		fmt.Println(v.String())
	})
}

func TestCallMethodByName(t *testing.T) {
	s := &Student{"matthew"}
	Convey("TestCallMethod with no parameter", t, func() {
		result, err := reflectutil.CallMethodByName(s, "GetName") // 问题， 反射下  地址引用方式，可以调用值引用方式的方法，但返过来，不行
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.GetName())
	})

	Convey("TestCallMethod with no such method name", t, func() {
		result, err := reflectutil.CallMethodByName(s, "NoSuchMethod")
		So(err, ShouldNotBeNil)

		So(result, ShouldBeNil)
	})

	Convey("TestCallMethod with one parameter", t, func() {

		var i I = &Value{"Hello"}
		result, err := reflectutil.CallMethodByName(s, "Greet", i)
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.Greet(i))
	})

	Convey("TestCallMethod with two parameters", t, func() {
		var i I = &Value{"Hello"}
		result, err := reflectutil.CallMethodByName(s, "GreetWithOption", i, "name")
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.GreetWithOption(i, "name"))
	})

}

func ExampleCallMethodByName() {
	s := &Student{"matthew"}
	result, err := reflectutil.CallMethodByName(s, "GetName")
	fmt.Println(result[0].Interface(), err)

	// Output:
	// matthew <nil>
}

type MultiFieldsPojo struct {
	IntArray2  [2]int
	Int32Slice []int
	MapValue   map[string]int
}

func TestSetValue(t *testing.T) {

	Convey("TestSetValue", t, func() {
		Convey("TestSetValue array", func() {
			arraytoSet := [2]int{3, 4}
			v := &MultiFieldsPojo{IntArray2: [2]int{1, 2}}
			r := reflectutil.SetValue(v, "IntArray2", arraytoSet)
			So(r, ShouldBeTrue)
			So(v.IntArray2, ShouldResemble, arraytoSet)
		})
		Convey("TestSetValue slice", func() {
			intSlice := make([]int, 10)
			v := &MultiFieldsPojo{}
			r := reflectutil.SetValue(v, "Int32Slice", intSlice)
			So(r, ShouldBeTrue)
			So(v.Int32Slice, ShouldResemble, intSlice)
		})
		Convey("TestSetValue map", func() {
			mp := make(map[string]int)
			mp["hello"] = 1
			v := &MultiFieldsPojo{}
			r := reflectutil.SetValue(v, "MapValue", mp)
			So(r, ShouldBeTrue)
			So(v.MapValue, ShouldResemble, mp)
		})

	})

}

type VIP struct {
	Name string
}

func TestInteface(t *testing.T) {
	vip := VIP{"matthew"}
	structV := reflect.ValueOf(vip)
	v2 := structV.Field(0) // 获取index=0字段的值
	x2 := v2.Interface()
	i2, ok2 := x2.(string)
	fmt.Printf("%s %v\n", i2, ok2) // matthew, true

	name, b := reflectutil.GetValue[string](&vip, "Name")
	fmt.Println(name, b)

	nameset := "string"
	name, b = reflectutil.GetValue[string](&nameset, "name")
	fmt.Println(name, b)
}
