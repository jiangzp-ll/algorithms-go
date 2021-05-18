package stack

// Stack , stack interface{}
type Stack interface {
	// Flush ,clear the stack
	Flush()
	// IsEmpty ,determine whether the stack is empty
	IsEmpty() bool
	// Pop , pop the element from the top of the stack
	Pop() interface{}
	// Push ,push the element to top of the stack
	Push(v interface{})
	// Top , get top of the stack
	Top() interface{}
}
