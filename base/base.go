/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-12 17:27:04
 */
package base

type Comparable[E any] interface {
	CompareTo(other E) int
}

type Equal[E any] interface {
	Comparable[E]
}

type Student struct {
	Name string
	Age  int
}
