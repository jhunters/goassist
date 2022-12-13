package containerx_test

import (
	"strings"
	"testing"

	"github.com/jhunters/goassist/containerx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInitList(t *testing.T) {

	Convey("TestList", t, func() {
		Convey("Test init list", func() {
			l := containerx.New[string]()
			size := l.Len()
			So(size, ShouldEqual, 0)
		})

		Convey("Test init list and fetch", func() {
			l := containerx.New[string]()
			e := l.Front()
			So(e, ShouldBeNil)

			e = l.Back()
			So(e, ShouldBeNil)

		})

	})

	// l.PushFront("hello")

}

func TestPushElement(t *testing.T) {

	Convey("TestPushElement", t, func() {

		Convey("Test PushFront", func() {
			v := "hello"
			l := containerx.New[string]()
			e := l.PushFront(v)
			size := l.Len()
			So(size, ShouldEqual, 1)
			So(e.Value, ShouldEqual, v)
			fe := l.Front()
			So(e, ShouldResemble, fe)

			be := l.Back()
			So(e, ShouldResemble, be)

		})
		Convey("Test PushBack", func() {
			v := "hello"
			l := containerx.New[string]()
			e := l.PushBack(v)
			size := l.Len()
			So(size, ShouldEqual, 1)
			So(e.Value, ShouldEqual, v)
			fe := l.Front()
			So(e, ShouldResemble, fe)

			be := l.Back()
			So(e, ShouldResemble, be)

		})
		Convey("Test PushBack & PushFront", func() {
			v := "hello"
			v2 := "world"
			l := containerx.New[string]()
			l.PushFront(v)
			l.PushBack(v2)
			size := l.Len()
			So(size, ShouldEqual, 2)

			fe := l.Front()
			So(v, ShouldResemble, fe.Value)
			So(fe.Prev(), ShouldBeNil)
			So(fe.Next().Value, ShouldEqual, v2)

			be := l.Back()
			So(v2, ShouldResemble, be.Value)

		})

	})

}

func TestToArray(t *testing.T) {
	Convey("TestToArray", t, func() {
		arr1 := []string{"1", "2", "3", "4", "5"}
		l := containerx.New[string]()
		for _, v := range arr1 {
			l.PushBack(v)
		}

		arr := l.ToArray()
		So(len(arr), ShouldEqual, 5)
		So(arr1, ShouldResemble, arr)

	})

}

func TestContains(t *testing.T) {

	Convey("TestContains", t, func() {
		l := createListX2()

		b := l.Contains("3", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
		So(b, ShouldBeTrue)
		b = l.Contains("6", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
		So(b, ShouldBeFalse)
	})

}

func TestIterator(t *testing.T) {
	Convey("TestIterator", t, func() {
		l := createListX2()
		arr := make([]string, l.Len())
		i := 0
		l.Iterator(func(s string) bool {
			arr[i] = s
			i++
			return true
		})
	})
}

func TestRemove(t *testing.T) {

	Convey("TestRemove", t, func() {
		Convey("Test remove", func() {
			l := createListX2()

			e, removed := l.Remove("2", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
			So(removed, ShouldBeTrue)
			So("2", ShouldEqual, e)
			So(9, ShouldEqual, l.Len())

			e, removed = l.Remove("11", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
			So(removed, ShouldBeFalse)
			So(e, ShouldBeEmpty)
			So(9, ShouldEqual, l.Len())
		})

		Convey("Test remove all", func() {
			l := createListX2()
			e, removed := l.RemoveAll("2", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
			So(removed, ShouldBeTrue)
			So("2", ShouldEqual, e)
			So(l.Len(), ShouldEqual, 8)

			e, removed = l.RemoveAll("11", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
			So(removed, ShouldBeFalse)
			So(e, ShouldBeEmpty)
			So(l.Len(), ShouldEqual, 8)
		})

	})

}

func TestClear(t *testing.T) {
	Convey("TestClear", t, func() {
		l := createList()
		l.Clear()
		So(l.IsEmpty(), ShouldBeTrue)
	})

}

func TestPushList(t *testing.T) {
	Convey("TestPushList", t, func() {
		Convey("Test PushFrontList", func() {
			l := createListX2()
			l2 := createList()
			l.PushFrontList(l2)
			So(15, ShouldEqual, l.Len())
			So("10", ShouldEqual, l.Front().Value)
		})
		Convey("Test PushBackList", func() {
			l := createListX2()
			l2 := createList()
			l.PushBackList(l2)
			So(15, ShouldEqual, l.Len())
			So("50", ShouldEqual, l.Back().Value)
		})
	})
}

func TestGetSet(t *testing.T) {
	Convey("TestGetSet", t, func() {

		Convey("Test Set", func() {
			l := createList()
			b := l.Set(0, "matt")
			So(l.Len(), ShouldEqual, 6)
			So(l.Front().Value, ShouldEqual, "matt")
			So(b, ShouldBeTrue)

			b = l.Set(1, "xml")
			So(l.Len(), ShouldEqual, 7)
			So(l.Front().Next().Value, ShouldEqual, "xml")
			So(b, ShouldBeTrue)

			b = l.Set(l.Len(), "out of pos")
			So(l.Len(), ShouldEqual, 7)
			So(b, ShouldBeFalse)

			b = l.Set(l.Len()-1, "bottom of pos")
			So(l.Len(), ShouldEqual, 8)
			So(b, ShouldBeTrue)
			So("bottom of pos", ShouldEqual, l.Back().Prev().Value)

		})

		Convey("Test Add", func() {
			l := createList()
			b := l.Add(0, "matt")
			So(l.Len(), ShouldEqual, 6)
			So(l.Front().Next().Value, ShouldEqual, "matt")
			So(b, ShouldBeTrue)

			b = l.Add(1, "xml")
			So(l.Len(), ShouldEqual, 7)
			So(l.Front().Next().Next().Value, ShouldEqual, "xml")
			So(b, ShouldBeTrue)

			b = l.Add(l.Len(), "out of pos")
			So(l.Len(), ShouldEqual, 7)
			So(b, ShouldBeFalse)

			b = l.Add(l.Len()-1, "bottom of pos")
			So(l.Len(), ShouldEqual, 8)
			So(b, ShouldBeTrue)
			So("bottom of pos", ShouldEqual, l.Back().Value)

		})

		Convey("Test Get", func() {
			l := createList()
			v, ok := l.Get(0)
			So(ok, ShouldBeTrue)
			So(v, ShouldEqual, "10")

			v, ok = l.Get(l.Len())
			So(ok, ShouldBeFalse)
			So(v, ShouldBeEmpty)
		})

	})

}

func TestRemoveFrontAndBack(t *testing.T) {
	Convey("TestRemoveFrontAndBack", t, func() {

		Convey("Test RemoveFront", func() {
			l := createList()
			v := l.RemoveFront()
			So(v, ShouldEqual, "10")
			So(4, ShouldEqual, l.Len())
		})

		Convey("Test RemoveBack", func() {
			l := createList()
			v := l.RemoveBack()
			So(v, ShouldEqual, "50")
			So(4, ShouldEqual, l.Len())
		})
	})
}

func createListX2() *containerx.List[string] {
	arr1 := []string{"1", "2", "3", "4", "5"}
	l := containerx.New[string]()
	for _, v := range arr1 {
		l.PushBack(v)
		l.PushBack(v)
	}
	return l
}

func createList() *containerx.List[string] {
	arr1 := []string{"10", "20", "30", "40", "50"}
	l := containerx.New[string]()
	for _, v := range arr1 {
		l.PushBack(v)
	}
	return l
}

func TestNewFromArray(t *testing.T) {
	Convey("TestNewFromArray", t, func() {
		want := containerx.New[string]()
		want.PushBack("1")
		want.PushBack("2")

		l := containerx.NewFromArray([]string{"1", "2"})
		So(want.Len(), ShouldEqual, l.Len())
		So(want.Front(), ShouldResemble, l.Front())
	})
}
