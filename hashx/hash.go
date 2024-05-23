package hashx

import (
	"hash/crc32"
	"hash/crc64"
)

// Hashcode to generate hash code
func Hashcode(b []byte) uint32 {
	v := crc32.ChecksumIEEE(b)
	return v
}

// Hashcode to generate hash code
func HashcodeString(s string) uint32 {
	v := crc32.ChecksumIEEE([]byte(s))
	return v
}

// Hashcode to generate hash code
func Hashcode64(b []byte) uint64 {
	v := crc64.Checksum(b, crc64.MakeTable(crc64.ISO))
	return v
}

// HashcodeSum 计算给定字节数组的 CRC64-ECMA 哈希值并返回。
// 参数 b 为需要计算哈希值的字节数组。
// 返回值为一个字节数组，表示计算得到的哈希值。
func HashcodeSum(b []byte) []byte {
	crc := crc64.New(crc64.MakeTable(crc64.ECMA))
	crc.Write(b)
	return crc.Sum(nil)
}
