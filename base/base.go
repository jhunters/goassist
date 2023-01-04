package base

type (
	Null struct{}

	Call func()

	// function provides one input argument and one return
	Func[R, T any] func(T) R

	// two-arity specialization of function
	BiFunc[R, T, U any] func(T, U) R

	// function provides one input argument and no returns
	Consumer[T any] func(T)

	// function provides two input arguments and no returns
	BiConsumer[T, U any] func(T, U)

	// function provides one input and one return
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
	Empty Null // const var for nil usage marker
)
