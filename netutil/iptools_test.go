package netutil_test

import (
	"fmt"
	"net"
	"net/netip"
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

func TestXxx(t *testing.T) {
	addr, err := netip.ParseAddr("fe80::c56b:90cd:b4fd:8a5c%16")
	// addr, err := netip.ParseAddr("ABCD:EF01:2345:6789:ABCD:EF01:127.0.0.1")
	fmt.Println(addr, err)

	fmt.Println(addr.As16())

	addr, err = netip.ParseAddr("fe80::c56b:90cd:b4fd:8a5c")
	// addr, err := netip.ParseAddr("ABCD:EF01:2345:6789:ABCD:EF01:127.0.0.1")
	fmt.Println(addr, err)

	fmt.Println(addr.As16())

}

func TestUint32ToIPv4_1(t *testing.T) {

	Convey("TestUint32ToIPv4_1", t, func() {
		So("12.12.12.12", ShouldEqual, netutil.Uint32ToIPv4(0x0C0C0C0C).String())
		So("192.168.0.1", ShouldEqual, netutil.Uint32ToIPv4(0x0100A8C0).String())
		So("0.0.0.0", ShouldEqual, netutil.Uint32ToIPv4(0x00000000).String())
		So("255.255.255.255", ShouldEqual, netutil.Uint32ToIPv4(0xFFFFFFFF).String())
	})

}
