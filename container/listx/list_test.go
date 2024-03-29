package listx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunters/goassist/container/listx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInitList(t *testing.T) {

	Convey("TestList", t, func() {
		Convey("Test init list", func() {
			l := listx.NewList[string]()
			size := l.Len()
			So(size, ShouldEqual, 0)
		})

		Convey("Test init list and fetch", func() {
			l := listx.NewList[string]()
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
			l := listx.NewList[string]()
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
			l := listx.NewList[string]()
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
			l := listx.NewList[string]()
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
		l := listx.NewList[string]()
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
		l.Iterate(func(s string) bool {
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
			So(l.Len(), ShouldEqual, 5)
			So(l.Front().Value, ShouldEqual, "matt")
			So(b, ShouldBeTrue)

			b = l.Set(1, "xml")
			So(l.Len(), ShouldEqual, 5)
			So(l.Front().Next().Value, ShouldEqual, "xml")
			So(b, ShouldBeTrue)

			b = l.Set(l.Len(), "out of pos")
			So(l.Len(), ShouldEqual, 5)
			So(b, ShouldBeFalse)

			b = l.Set(l.Len()-1, "bottom of pos")
			So(l.Len(), ShouldEqual, 5)
			So(b, ShouldBeTrue)
			So("bottom of pos", ShouldEqual, l.Back().Value)

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

func createListX2() *listx.List[string] {
	arr1 := []string{"1", "2", "3", "4", "5"}
	l := listx.NewList[string]()
	for _, v := range arr1 {
		l.PushBack(v)
		l.PushBack(v)
	}
	return l
}

func createList() *listx.List[string] {
	arr1 := []string{"10", "20", "30", "40", "50"}
	l := listx.NewList[string]()
	for _, v := range arr1 {
		l.PushBack(v)
	}
	return l
}

func TestNewListFromArray(t *testing.T) {
	Convey("TestNewListFromArray", t, func() {
		want := listx.NewList[string]()
		want.PushBack("1")
		want.PushBack("2")

		l := listx.NewListFromArray([]string{"1", "2"})
		So(want.Len(), ShouldEqual, l.Len())
		So(want.Front(), ShouldResemble, l.Front())
	})
}

func TestWriteToArray(t *testing.T) {

	Convey("TestWriteToArray", t, func() {

		l := createListX2()

		// read empty
		l.WriteToArray(nil)

		// read partial elements
		v := make([]string, 5)
		l.WriteToArray(v)
		arr1 := []string{"1", "1", "2", "2", "3"}
		So(arr1, ShouldResemble, v)

		// read full
		v = make([]string, l.Len())
		l.WriteToArray(v)
		So(l.ToArray(), ShouldResemble, v)

		// read full
		v = make([]string, l.Len()<<1)
		l.WriteToArray(v)
		So(l.ToArray(), ShouldResemble, v[:l.Len()])
	})
}

func TestIndex(t *testing.T) {
	Convey("TestIndex", t, func() {
		l := createListX2()
		index := l.Index("2", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
		So(2, ShouldEqual, index)

		index = l.Index("22", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
		So(-1, ShouldEqual, index)
	})

	Convey("TestLastIndex", t, func() {
		l := createListX2()
		index := l.LastIndex("2", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
		So(3, ShouldEqual, index)

		index = l.LastIndex("22", func(s1, s2 string) bool { return strings.Compare(s1, s2) == 0 })
		So(-1, ShouldEqual, index)
	})

}

func TestFilter(t *testing.T) {
	Convey("TestFilter", t, func() {
		l := createListX2()
		ll := l.Filter(func(s string) bool {
			return strings.Compare(s, "2") > 0
		})
		So(ll.Len(), ShouldEqual, 6)
	})

}

func TestMin(t *testing.T) {
	Convey("TestMin", t, func() {
		l := createListX2()
		e := l.Min(func(o1, o2 string) int {
			return strings.Compare(o1, o2)
		})
		So(e, ShouldEqual, "1")
	})

}
func TestMax(t *testing.T) {
	Convey("TestMax", t, func() {
		l := createListX2()
		e := l.Max(func(o1, o2 string) int {
			return strings.Compare(o1, o2)
		})
		So(e, ShouldEqual, "5")
	})

}

func TestSort(t *testing.T) {
	Convey("TestSort", t, func() {
		l := listx.NewListOf("3", "4", "9", "6", "2", "5", "1")

		l.Sort(func(o1, o2 string) int {
			return strings.Compare(o1, o2)
		})
		So(l.FrontValue(), ShouldEqual, "1")
	})

}

func TestCopy(t *testing.T) {
	Convey("TestCopy", t, func() {
		l := listx.NewListOf("3", "4", "9", "6", "2", "5", "1")

		l2 := l.Copy()
		So(l2.Len(), ShouldEqual, l.Len())
		So(l2.ToArray(), ShouldResemble, l.ToArray())
	})
}

func ExampleNewList() {

	l := listx.NewList[string]()
	fmt.Println(l.Len()) // len is 0

	l.Add(0, "golang")
	fmt.Println(l.Len()) // len is 1

	l.PushFront("java")
	fmt.Println(l.FrontValue())

	l.PushBack("python")
	fmt.Println(l.BackValue())

	exist := l.Contains("java", func(s1, s2 string) bool { return strings.EqualFold(s1, s2) })
	fmt.Println(exist)

	l.Sort(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	l.Range(func(s string) bool {
		fmt.Println(s)
		return true
	})

	// Output:
	// 0
	// 1
	// java
	// python
	// true
	// golang
	// java
	// python
}

func TestInsertAndMove(t *testing.T) {
	arr1 := []string{"10", "20", "30", "40", "50"}
	l := listx.NewListOf(arr1...)

	Convey("test insert and move", t, func() {
		fe := l.Front()
		So(fe.Value, ShouldEqual, "10") // 10, 20, 30, 40, 50

		// move to back
		l.MoveToBack(fe) // 20, 30, 40, 50, 10

		fe = l.Front()
		So(fe.Value, ShouldEqual, "20")

		be := l.Back()
		So(be.Value, ShouldEqual, "10")

		e := l.PushFront("1") // 1, 20, 30, 40, 50, 10
		So(e.Value, ShouldEqual, "1")

		l.MoveAfter(e, be) // 20, 30, 40, 50, 10, 1
		be = l.Back()
		So(be.Value, ShouldEqual, "1")

		be2 := l.PushBack("100") // 20, 30, 40, 50, 10, 1, 100

		l.MoveBefore(be2, be) // 20, 30, 40, 50, 10, 100, 1
		So(l.BackValue(), ShouldEqual, "1")

		l.InsertBefore("200", be2) // 20, 30, 40, 50, 10, 200, 100, 1

		v, exist := l.Get(5)
		So(v, ShouldEqual, "200")
		So(exist, ShouldBeTrue)

	})
}
