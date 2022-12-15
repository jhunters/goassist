package base

type (
	// compare function
	CMP[E any] func(E, E) int

	// equal function
	EQL[E any] func(E, E) bool

	Null struct{}

	// Evaluate use the specified parameter to perform a test that returns true or false.
	Evaluate[E any] func(E) bool
)

var (
	Empty Null // const var for nil usage marker
)
