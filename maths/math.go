/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-05 15:52:22
 */
package maths

import "constraints"

// Max return the greater one
func Max[E constraints.Ordered](a, b E) E {
	if a >= b {
		return a
	}
	return b
}

// Min return the smaller one
func Min[E constraints.Ordered](a, b E) E {
	if a <= b {
		return a
	}
	return b
}

// Abs returns the absolute value of an target value.
func Abs[E constraints.Signed](value E) E {
	if value < 0 {
		return -1 * value
	}
	return value
}
