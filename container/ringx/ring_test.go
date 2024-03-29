package ringx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/container/ringx"
	. "github.com/smartystreets/goconvey/convey"
)

type RingPojo struct {
	Name string
}

func NewRingPojo(name string) *RingPojo {
	return &RingPojo{Name: name}
}

func compareRingPojo(o1, o2 *RingPojo) int {
	if o1 == nil && o2 == nil {
		return 0
	} else {
		var n1 string
		if o1 == nil {
			return -1
		} else {
			n1 = o1.Name
		}

		var n2 string
		if o2 == nil {
			return 1
		} else {
			n2 = o2.Name
		}
		return strings.Compare(n1, n2)
	}
}

func createRing() *ringx.Ring[*RingPojo] {
	r := ringx.NewRingOf(NewRingPojo("matt"), NewRingPojo("xml"), NewRingPojo("ant_miracle"),
		NewRingPojo("michael"), NewRingPojo(""), NewRingPojo("ryan"))
	return r
}

func createRingPojoArray() []*RingPojo {
	return arrayutil.AsList(NewRingPojo("matt"), NewRingPojo("xml"), NewRingPojo("ant_miracle"),
		NewRingPojo("michael"), NewRingPojo(""), NewRingPojo("ryan"))
}

func TestNewRing(t *testing.T) {

	Convey("TestNewRing", t, func() {
		r := ringx.NewRing[*RingPojo](10)
		So(10, ShouldEqual, r.Len())

		r.Do(func(rp *RingPojo) {
			So(rp, ShouldBeNil)
		})
	})
	Convey("TestNewRingOf", t, func() {
		r := createRing()
		So(r.Len(), ShouldEqual, 6)

		r.Do(func(rp *RingPojo) {
			So(rp, ShouldNotBeNil)
		})
	})
}

func ExampleNewRingOf() {
	r := ringx.NewRingOf("a", "b", "c", "d", "e")
	fmt.Println(r.Len())

	r.Do(func(s string) {
		fmt.Println(s)
	})

	// move to prev
	r = r.Prev()
	fmt.Println(r.Value)

	// link
	r = r.Link(ringx.NewRingOf("1", "2", "3"))
	fmt.Println(r.Len())

	r.Do(func(s string) {
		fmt.Println(s)
	})

	// Output:
	// 5
	// a
	// b
	// c
	// d
	// e
	// e
	// 8
	// a
	// b
	// c
	// d
	// e
	// 1
	// 2
	// 3
}

func TestMinMax(t *testing.T) {
	Convey("TestMinMax", t, func() {
		Convey("Test min in initial ring", func() {
			r := ringx.NewRing[*RingPojo](10)
			v := r.Min(compareRingPojo)
			So(v, ShouldBeNil)
		})
		Convey("Test max in initial ring", func() {
			r := ringx.NewRing[*RingPojo](10)
			v := r.Max(compareRingPojo)
			So(v, ShouldBeNil)
		})
		Convey("Test min in ring", func() {
			r := createRing()
			v := r.Min(compareRingPojo)
			So(v, ShouldNotBeNil)
			So(v.Name, ShouldBeEmpty)
		})
		Convey("Test max in ring", func() {
			r := createRing()
			v := r.Max(compareRingPojo)
			So(v.Name, ShouldEqual, "xml")
		})

	})

}

func TestRingLink(t *testing.T) {
	Convey("TestRingLink", t, func() {
		r := createRing()
		r2 := createRing()
		r3 := r.Link(r2)
		So(r.Len(), ShouldEqual, 12)
		So(r3, ShouldNotBeNil)
		So(r3.Len(), ShouldEqual, 12)
		So(r3.Value.Name, ShouldEqual, "xml")
	})
}

func TestRingLinkValueAndGet(t *testing.T) {
	Convey("TestRingLinkValueAndGet", t, func() {
		r := ringx.NewRing[*RingPojo](1)
		So(r.Len(), ShouldEqual, 1)
		So(r.Value, ShouldBeNil)

		r.Value = NewRingPojo("matt")

		r.LinkValue(NewRingPojo("xml"))
		nr := r.Get(1)
		So(nr.Name, ShouldEqual, "xml")
		So(r.Len(), ShouldEqual, 2)

		e := r.Unlink(1)
		So(r.Len(), ShouldEqual, 1)
		So(e.Value.Name, ShouldEqual, "xml")

	})

}

func TestSortRing(t *testing.T) {
	Convey("TestSortRing", t, func() {
		r := createRing()
		r.Sort(compareRingPojo)

		So(r.Value.Name, ShouldBeEmpty)
		So(r.Prev().Value.Name, ShouldEqual, "xml")
	})
}

func TestRingNext(t *testing.T) {
	Convey("TestRingNext", t, func() {

		r := ringx.NewRing[*RingPojo](1)
		So(r.Len(), ShouldEqual, 1)

		e := r.Next()
		So(e, ShouldNotBeNil)
		So(e.Next(), ShouldNotBeNil)
	})
}

func TestRingMove(t *testing.T) {
	Convey("TestRingMove", t, func() {

		r := createRing()
		e := r.Move(1)
		So(e.Value.Name, ShouldEqual, "xml")

		e = r.Move(1 + r.Len())
		So(e.Value.Name, ShouldEqual, "xml")

		e = r.Move(1 - r.Len())
		So(e.Value.Name, ShouldEqual, "xml")
	})
}

func TestRingContains(t *testing.T) {
	Convey("TestRingContains element exist", t, func() {
		r := createRing()
		b := r.Contains(NewRingPojo(""), func(rp1, rp2 *RingPojo) bool { return compareRingPojo(rp1, rp2) == 0 })
		So(b, ShouldBeTrue)

		idx := r.Index(NewRingPojo(""), func(rp1, rp2 *RingPojo) bool { return compareRingPojo(rp1, rp2) == 0 })
		So(idx, ShouldBeGreaterThan, 0)
	})

	Convey("TestRingContains element not exist", t, func() {
		r := createRing()
		b := r.Contains(NewRingPojo("unknown"), func(rp1, rp2 *RingPojo) bool { return compareRingPojo(rp1, rp2) == 0 })
		So(b, ShouldBeFalse)

		idx := r.Index(NewRingPojo("unknown"), func(rp1, rp2 *RingPojo) bool { return compareRingPojo(rp1, rp2) == 0 })
		So(idx, ShouldEqual, -1)
	})

}

func TestRingToArray(t *testing.T) {
	Convey("TestRingToArray", t, func() {
		r := createRing()
		arr := r.ToArray()
		So(arr, ShouldResemble, createRingPojoArray())

		arr2 := make([]*RingPojo, 2)
		r.WriteToArray(arr2)
		So(arr2, ShouldResemble, createRingPojoArray()[:2])
	})
}

func TestRingCopy(t *testing.T) {
	Convey("TestRingCopy", t, func() {
		r := createRing()

		r2 := r.Copy()
		So(r2.Len(), ShouldEqual, 6)

	})

}

func TestRingIsLoop(t *testing.T) {
	Convey("TestRingIsRealRing", t, func() {
		r := createRing()
		loop := false
		r1 := r
		r2 := r
		checkCount := r.Len() * 3
		var b bool
		for i := 0; i < checkCount; i++ {
			r1, b = step(r1, 1)
			So(b, ShouldBeTrue)
			r2, b = step(r2, 2)
			So(b, ShouldBeTrue)
			if r1 == r2 {
				loop = true
				break
			}
		}
		So(loop, ShouldBeTrue)
	})
}

func step[E any](r *ringx.Ring[E], step int) (*ringx.Ring[E], bool) {
	if r == nil {
		return nil, false
	}
	ret := r
	for i := 0; i < step; i++ {
		ret = ret.Next()
		if ret == nil {
			return nil, false
		}
	}

	return ret, true
}
