package set_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/container/set"
	. "github.com/smartystreets/goconvey/convey"
)

type SetPojo struct {
	Name string
}

func newSetPojo(name string) SetPojo {
	return SetPojo{name}
}

func createSet() *set.Set[SetPojo] {
	mp := set.NewSet[SetPojo]()
	mp.Add(newSetPojo("hello"))
	mp.Add(newSetPojo("world"))
	mp.Add(newSetPojo("to"))
	mp.Add(newSetPojo("you"))
	mp.Add(newSetPojo("!"))
	return mp
}

func TestNewSet(t *testing.T) {
	Convey("TestNewSet", t, func() {
		mp := set.NewSet[*SetPojo]()
		So(mp.Size(), ShouldBeZeroValue)
		So(mp.IsEmpty(), ShouldBeTrue)
	})

}

func TestSetAddRemove(t *testing.T) {
	Convey("TestAddRemove", t, func() {
		s := createSet()
		So(s.Size(), ShouldEqual, 5)

		s.Add(newSetPojo("world"))
		So(s.Size(), ShouldEqual, 5)

		ok := s.Remove(newSetPojo("world"))
		So(s.Size(), ShouldEqual, 4)
		So(ok, ShouldBeTrue)

		// remove not exist key
		ok = s.Remove(newSetPojo("world"))
		So(s.Size(), ShouldEqual, 4)
		So(ok, ShouldBeFalse)

	})

	Convey("TestClear", t, func() {
		s := createSet()
		So(s.Size(), ShouldEqual, 5)

		s.Clear()
		So(s.Size(), ShouldEqual, 0)

	})
}

func TestSetExist(t *testing.T) {
	Convey("TestExist", t, func() {
		s := createSet()
		ok := s.Exist(newSetPojo("world"))
		So(ok, ShouldBeTrue)

		ok = s.Exist(newSetPojo("unknown"))
		So(ok, ShouldBeFalse)
	})
}

func TestSetRange(t *testing.T) {
	Convey("TestSetRange", t, func() {
		s := createSet()

		count := 0
		s.Range(func(sp SetPojo) bool {
			ok := s.Exist(sp)
			So(ok, ShouldBeTrue)
			count++
			return true
		})

		So(count, ShouldEqual, 5)

		Convey("TestSetRange shot break", func() {
			count := 0
			s.Range(func(sp SetPojo) bool {
				count++
				return false
			})
			So(count, ShouldEqual, 1)
		})

	})
}

func TestCopy(t *testing.T) {
	Convey("TestCopy", t, func() {
		Convey("Test copy empty set", func() {
			mp := set.NewSet[*SetPojo]()

			mp2 := mp.Copy()
			So(mp2.Size(), ShouldBeZeroValue)
			So(mp2.IsEmpty(), ShouldBeTrue)
		})
		Convey("Test copy set", func() {
			mp := createSet()

			mp2 := mp.Copy()
			So(mp2.Size(), ShouldEqual, mp.Size())
			So(mp2.IsEmpty(), ShouldEqual, mp.IsEmpty())
		})
	})
}

func TestSetToArray(t *testing.T) {
	Convey("TestSetToArray", t, func() {
		s := createSet()
		arr := s.ToArray()
		So(len(arr), ShouldEqual, s.Size())

		for _, v := range arr {
			So(s.Exist(v), ShouldBeTrue)
		}
	})
}

func ExampleNewSet() {
	// create a new set
	st := set.NewSet[string]()
	fmt.Println(st.IsEmpty())

	ok := st.Add("v1")
	fmt.Println(ok) // true

	ok = st.Add("v1")
	fmt.Println(ok, st.Size())

	exist := st.Exist("v1")
	fmt.Println(exist) // true

	exist = st.Remove("v1")
	fmt.Println(exist) // true

	// Output:
	// true
	// true
	// false 1
	// true
	// true
}
