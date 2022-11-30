package mapx_test

import (
	"testing"

	"github.com/jhunters/goassist/mapx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClone(t *testing.T) {

	m := make(map[string]string)
	m["hello"] = "world"
	m["name"] = "matthew"

	newM := mapx.Clone(m)

	Convey("Test add all", t, func() {
		So(m, ShouldResemble, newM)
	})
}

func TestAddAll(t *testing.T) {

	m1 := make(map[string]string)
	m1["hello"] = "world"
	m1["name"] = "matthew"

	m2 := make(map[string]string)

	m2 = mapx.AddAll(m2, m1)

	Convey("Test add all", t, func() {
		So(m2, ShouldResemble, m1)
	})

}
