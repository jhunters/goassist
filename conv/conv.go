/*
 * Package conv provides convert utility apis
 */
package conv

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jhunters/goassist/generic"
)

const (
	Binary = 2
	Octal  = 10
	Hex    = 16
)

// Itoa formate integer and float value to string type
func Itoa[E generic.Integer | generic.Float](i E) string {
	return fmt.Sprintf("%v", i)
}

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatInt[E generic.Integer](i E, base int) string {
	if i > 0 && uint64(i) > math.MaxInt64 {
		return strconv.FormatUint(uint64(i), base)
	}
	return strconv.FormatInt(int64(i), base)
}

// ToPtr convert to pointer
func ToPtr[E any](t E) *E {
	return &t
}
