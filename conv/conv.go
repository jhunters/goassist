/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-26 11:42:28
 */
package conv

import (
	"constraints"
	"fmt"
	"math"
	"strconv"
)

// Itoa formate integer and float value to string type
func Itoa[E constraints.Integer | constraints.Float](i E) string {
	return fmt.Sprintf("%v", i)
}

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatInt[E constraints.Integer](i E, base int) string {
	if i > 0 && uint64(i) > math.MaxInt64 {
		return strconv.FormatUint(uint64(i), base)
	}
	return strconv.FormatInt(int64(i), base)
}
