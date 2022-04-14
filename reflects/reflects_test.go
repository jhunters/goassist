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
	"unsafe"

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

func TestCase(t *testing.T) {
	s := Student{"matthew"}
	s.ChangeName2("xml")
	fmt.Println(s.GetName()) // print  matthew   value ref way just copy a  new struct object

	s2 := &Student{"matthew"}
	s2.ChangeName("xml")
	fmt.Println(s2.GetName()) // print  xml   ptr ref way

	// ptr := unsafe.Pointer(s2)
	// fmt.Println(ptr)
	// sz := unsafe.Sizeof(s2)
	// fmt.Println(sz)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	r := unsafe.Slice(&arr[0], 2) // 从地址时，获取切片内容
	fmt.Println(r)

	// 这里的代码等同于上面
	arr2 := (*[0x7FFFFFFF]int)(unsafe.Pointer(&arr[0])) // 把cap大小设置了0x7FFFFFFF， 表示可以越界读到更多内容
	fmt.Println(arr2[:2])

	fmt.Println(0x7FFFFFFF, len(arr2))
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
	s := &Student{"matthew"}
	Convey("TestCallMethod with no parameter", t, func() {
		result, err := reflects.CallMethodName(s, "GetName") // 问题， 反射下  地址引用方式，可以调用值引用方式的方法，但返过来，不行
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.GetName())
	})

	Convey("TestCallMethod with no such method name", t, func() {
		result, err := reflects.CallMethodName(s, "NoSuchMethod")
		So(err, ShouldNotBeNil)

		So(result, ShouldBeNil)
	})

	Convey("TestCallMethod with one parameter", t, func() {

		var i I = &Value{"Hello"}
		result, err := reflects.CallMethodName(s, "Greet", i)
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.Greet(i))
	})

	Convey("TestCallMethod with two parameters", t, func() {
		var i I = &Value{"Hello"}
		result, err := reflects.CallMethodName(s, "GreetWithOption", i, "name")
		So(err, ShouldBeNil)

		So(result[0].Interface(), ShouldEqual, s.GreetWithOption(i, "name"))
	})

}
