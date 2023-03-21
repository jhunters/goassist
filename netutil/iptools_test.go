package netutil_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/jhunters/goassist/netutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIPv4ToUInt32(t *testing.T) {

	Convey("TestIPv4ToUInt32", t, func() {
		ip := net.IPv4(127, 0, 1, 2)
		v := netutil.IPv4ToUInt32(ip)
		So(v, ShouldEqual, 33620095)
	})

}

func TestUint32ToIPv4(t *testing.T) {

	Convey("TestUint32ToIPv4", t, func() {
		ip := netutil.Uint32ToIPv4(33620095)
		So(ip.String(), ShouldEqual, "127.0.1.2")
	})
}

func ExampleIPv4ToUInt32() {
	ip := net.IPv4(127, 0, 1, 2)
	v := netutil.IPv4ToUInt32(ip)
	fmt.Println(v)

	// Output:
	// 33620095
}

func ExampleUint32ToIPv4() {
	ip := netutil.Uint32ToIPv4(33620095)
	fmt.Println(ip.String())

	// Output:
	// 127.0.1.2
}
