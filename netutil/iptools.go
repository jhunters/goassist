package netutil

import (
	"net"
)

// IPv4ToUInt32 convert to ipv4 as uint32
func IPv4ToUInt32(ip net.IP) uint32 {
	ipv4 := ip.To4()
	if ipv4 == nil {
		return 0
	}
	var ipInt32 uint32
	ipInt32 = ipInt32 | uint32(ipv4[0])
	ipInt32 = ipInt32 | uint32(ipv4[1])<<8
	ipInt32 = ipInt32 | uint32(ipv4[2])<<16
	ipInt32 = ipInt32 | uint32(ipv4[3])<<24

	return ipInt32
}

// Uint32ToIPv4 convert unit32 to ipv4
func Uint32ToIPv4(v uint32) net.IP {
	a := byte(v & 0x000000FF)
	b := byte((v & 0x0000FF00) >> 8)
	c := byte((v & 0x00FF0000) >> 16)
	d := byte((v & 0xFF000000) >> 24)
	return net.IPv4(a, b, c, d)
}
