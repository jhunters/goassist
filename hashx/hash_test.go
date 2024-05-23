package hashx_test

import (
	"encoding/json"
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

	Convey("Hashcode64", t, func() {
		hc := hashx.Hashcode64([]byte("nodeA"))
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)
	})

	Convey("HashcodeSum", t, func() {
		hc := hashx.HashcodeSum([]byte("sdfdsfdsfsdfdd"))
		r, err := json.Marshal(hc)
		So(err, ShouldBeNil)
		So(r, ShouldNotBeNil)
	})
}
