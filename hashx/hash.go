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
