package containerx_test

import (
	"strings"
	"testing"

	"github.com/jhunters/goassist/containerx"
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

func createRing() *containerx.Ring[*RingPojo] {
	r := containerx.NewRingOf(NewRingPojo("matt"), NewRingPojo("xml"), NewRingPojo("ant_miracle"),
		NewRingPojo("michael"), NewRingPojo(""), NewRingPojo("ryan"))
	return r
}

func TestNewRing(t *testing.T) {

	Convey("TestNewRing", t, func() {
		r := containerx.NewRing[*RingPojo](10)
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

func TestMinMax(t *testing.T) {
	Convey("TestMinMax", t, func() {
		Convey("Test min in initial ring", func() {
			r := containerx.NewRing[*RingPojo](10)
			v := r.Min(compareRingPojo)
			So(v, ShouldBeNil)
		})
		Convey("Test max in initial ring", func() {
			r := containerx.NewRing[*RingPojo](10)
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

func TestSortRing(t *testing.T) {
	Convey("TestSortRing", t, func() {
		r := createRing()
		r.Sort(compareRingPojo)

		So(r.Value.Name, ShouldBeEmpty)
		So(r.Prev().Value.Name, ShouldEqual, "xml")
	})
}
