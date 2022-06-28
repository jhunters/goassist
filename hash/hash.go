package hash

import "hash/crc32"

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
