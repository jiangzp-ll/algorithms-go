package stack

// Stack , stack interface{}
type Stack interface {
	// Flush ,clear the stack
	Flush()
	// IsEmpty ,determine whether the stack is empty
	IsEmpty() bool
	// Peek ,get and not remove the element from the top of the stack
	Peek() (interface{}, error)
	// Pop ,pop and remove the element from the top of the stack
	Pop() (interface{}, error)
	// Push ,push the element to top of the stack
	Push(v interface{})
	// Search ,return the index of the value in the stack
	Search(val interface{}) (int, error)
}
