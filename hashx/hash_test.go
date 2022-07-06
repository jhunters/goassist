package hashx_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/conv"
	"github.com/jhunters/goassist/hash"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	m = 0b1111111111
)

func TestHashcode(t *testing.T) {

	Convey("TestHashcode", t, func() {

		hc := hash.HashcodeString("nodeA")
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)
		fmt.Println(hc, conv.FormatInt(hc, 2), hc%m)

		hc = hash.HashcodeString("nodeB")
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)
		fmt.Println(hc, conv.FormatInt(hc, 2), hc%m)

		hc = hash.HashcodeString("nodeC")
		So(hc, ShouldBeGreaterThanOrEqualTo, 0)
		fmt.Println(hc, conv.FormatInt(hc, 2), hc%m)
	})
}
