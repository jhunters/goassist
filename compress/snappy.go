package compress

import "github.com/golang/snappy"

// SnappyEncode do compress by snappy
func SnappyEncode(b []byte) []byte {
	dst := make([]byte, snappy.MaxEncodedLen(len(b)))
	return snappy.Encode(dst, b)
}

// SnappyDecode do unCompress by snappy
func SnappyDecode(b []byte) ([]byte, error) {
	dst := make([]byte, 1)
	return snappy.Decode(dst, b)
}
