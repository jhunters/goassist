/*
 * Package mathutil to provides utility api for math
 */
package mathutil

import (
	"fmt"

	"github.com/jhunters/goassist/generic"
)

// Max return the greater one
func Max[E generic.Ordered](a, b E) E {
	if a >= b {
		return a
	}
	return b
}

// Min return the smaller one
func Min[E generic.Ordered](a, b E) E {
	if a <= b {
		return a
	}
	return b
}

// Abs returns the absolute value of an target value.
func Abs[E generic.Signed](value E) E {
	if value < 0 {
		return -1 * value
	}
	return value
}

//
//  SafeAdd to proccess plus action if overflow return error
func SafeAdd[E generic.Signed](a, b E) (E, error) {
	ret := a + b
	if a^b < 0 || a^ret >= 0 {
		return ret, nil
	}
	return ret, fmt.Errorf("%v plus %v occures overflow.", a, b)
}

//  SafeAdd to proccess plus action if overflow return error
func SafeAddUnsigned[E generic.Unsigned](a, b E) (E, error) {
	ret := a + b
	if a == 0 || b == 0 {
		return ret, nil
	}

	if ret < a || ret < b {
		return ret, fmt.Errorf("%v plus %v occures overflow.", a, b)
	}
	return ret, nil
}

// SafeSubstract to proccess substract action if overflow return error
func SafeSubstract[E generic.Signed](a, b E) (E, error) {
	ret := a - b

	if a^b >= 0 || a^ret >= 0 {
		return ret, nil
	}

	return ret, fmt.Errorf("%v substract %v occures overflow.", a, b)
}

// SafeSubstractUnsigned to proccess substract action if overflow return error
func SafeSubstractUnsigned[E generic.Unsigned](a, b E) (E, error) {
	ret := a - b
	if a >= b {
		return ret, nil
	}
	return ret, fmt.Errorf("%v substract %v occures overflow.", a, b)
}

//  Mod returns {@code x mod m}, a non-negative value less than {@code m}.
// This differs from {@code x % m}, which might be negative.
func Mod[E generic.Integer](x, m E) (E, error) {
	if m < 0 {
		return x, fmt.Errorf("Modulus must be positive")
	}

	ret := x % m
	if ret < 0 {
		return ret + m, nil
	}
	return ret, nil
}
