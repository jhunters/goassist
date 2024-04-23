package netutil

import (
	"fmt"
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

// IPv4StringToUInt32 将IPv4地址字符串转换为无符号32位整数
// 参数ip：IPv4地址字符串
// 返回值1：转换后的无符号32位整数
// 返回值2：如果转换失败，则返回错误信息
func IPv4StringToUInt32(ip string) (uint32, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return 0, fmt.Errorf("invalid ip: %s", ip)
	}
	return IPv4ToUInt32(ipv4), nil
}

// Uint32ToIPv4 convert unit32 to ipv4
func Uint32ToIPv4(v uint32) net.IP {
	a := byte(v & 0x000000FF)
	b := byte((v & 0x0000FF00) >> 8)
	c := byte((v & 0x00FF0000) >> 16)
	d := byte((v & 0xFF000000) >> 24)
	return net.IPv4(a, b, c, d)
}
