package hashx_test

import (
	"testing"

	"github.com/jhunters/goassist/hashx"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHashcode(t *testing.T) {

	Convey("HashcodeString", t, func() {
		hc := hashx.HashcodeString("nodeA")
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)

		hc = hashx.HashcodeString("nodeB")
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)

		hc = hashx.HashcodeString("nodeC")
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)
	})

	Convey("TestHashcode", t, func() {
		hc := hashx.Hashcode([]byte("nodeA"))
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)

		hc = hashx.Hashcode([]byte("nodeB"))
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)

		hc = hashx.Hashcode([]byte("nodeC"))
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)
	})
}
