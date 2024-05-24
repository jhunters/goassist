/*
 * Package base to defines const and base types
 */
package base

type (
	Null struct{} // a struct representing a null value.

	Call func() // a function type with no parameters and no return value

	// function provides one input argument and one return
	Func[R, T any] func(T) R

	// two-arity specialization of function
	BiFunc[R, T, U any] func(T, U) R

	// function provides one input argument and no returns
	Consumer[T any] func(T)

	// function provides two input arguments and no returns
	BiConsumer[T, U any] func(T, U)

	// function provides no input and one return
	Supplier[R any] func() R

	// Evaluate use the specified parameter to perform a test that returns true or false.
	Evaluate[E any] Func[bool, E]

	// compare function
	CMP[E any] BiFunc[int, E, E]

	// equal function
	EQL[E any] BiFunc[bool, E, E]

	// Comparable is a interface to compare action
	Comparable[E any] interface {
		CompareTo(v E) int
	}
)

var (
	Empty Null // serves as a constant marker for nil usage

	Dummy func() = func() { recover() }
)

// SafetyCall 函数接收一个无参无返回值的函数作为参数，并在调用该函数前后执行一些操作
// 参数fn：需要被调用的函数
func SafetyCall(fn Call) {
	defer Dummy()
	fn()
}

// SafetyFunc 是一个泛型函数，用于安全地执行给定函数 fn 并返回结果
// 参数 t 为 fn 函数的输入参数类型 T 的实例
// 参数 fn 为一个函数类型 Func[R, T]，其中 R 为 fn 函数返回值的类型，T 为 fn 函数输入参数的类型
// 函数返回 fn 函数执行后的结果，类型为 R
func SafetyFunc[R, T any](t T, fn Func[R, T]) R {
	defer Dummy()
	r := fn(t)
	return r
}

// SafetyBiFunc 是一个泛型函数，接收两个任意类型的参数 t 和 u，以及一个 BiFunc 函数类型的参数 fn
// 函数的功能是：调用 fn 函数，并将 t 和 u 作为参数传入，返回 fn 函数的返回值
// 参数：
//     t：任意类型的参数
//     u：任意类型的参数
//     fn：BiFunc 函数类型的参数，接收两个参数，返回值为 R 类型
// 返回值：
//     R：fn 函数的返回值
func SafetyBiFunc[R, T, U any](t T, u U, fn BiFunc[R, T, U]) R {
	defer Dummy()
	r := fn(t, u)
	return r
}

// SafetyConsumer 是一个泛型函数，用于安全地消费类型为R的值
//
// 参数：
//     r R - 需要被消费的值
//     fn Consumer[R] - 消费函数，接收一个类型为R的参数，没有返回值
//
// 返回值：
//     无
//
// 说明：
//     使用defer语句调用Dummy函数，确保在消费函数执行完毕后执行一些清理操作或记录日志等。
//     该函数主要用于防止消费函数执行时出现panic等异常情况，导致整个程序崩溃。
func SafetyConsumer[T any](t T, fn Consumer[T]) {
	defer Dummy()
	fn(t)
}

// SafetyBiConsumer 是一个安全的二元消费者函数，用于接收两个参数并调用传入的二元消费者函数
// 参数t和u分别为传入二元消费者函数的两个参数
// 参数fn为传入的二元消费者函数，其类型必须为BiConsumer[T, U]
// 该函数会调用Dummy函数确保在函数结束时进行清理工作
func SafetyBiConsumer[T, U any](t T, u U, fn BiConsumer[T, U]) {
	defer Dummy()
	fn(t, u)
}

// SafetySupplier 是一个泛型函数，用于安全地执行 Supplier 函数 fn 并返回其结果。
// 参数 fn 是一个 Supplier 函数，返回类型为 R。
// 函数使用 defer 语句调用了 Dummy 函数，用于处理可能的异常或进行清理工作。
// 函数执行 fn 并将其结果赋值给 r，然后返回 r。
func SafetySupplier[R any](fn Supplier[R]) R {
	defer Dummy()
	r := fn()
	return r
}

// SafetyEvaluate 是一个泛型函数，用于安全地评估某个值e是否满足Evaluate[E]函数的要求
// 参数e为待评估的值，类型为泛型E
// 参数fn为Evaluate[E]类型的函数，用于对e进行评估
// 函数会延迟调用Dummy函数，以处理可能存在的异常情况
// 函数返回值为bool类型，表示e是否满足fn的要求
func SafetyEvaluate[E any](e E, fn Evaluate[E]) bool {
	defer Dummy()
	return fn(e)
}
