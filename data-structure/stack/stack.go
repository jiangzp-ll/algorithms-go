package stack

// Stack , stack interface{}
type Stack interface {
	// Check ,check input value
	Check(val interface{}) error
	// Flush ,clear the stack
	Flush()
	// IsEmpty ,determine whether the stack is empty
	IsEmpty() bool
	// Len ,get the number of elements in the stack
	Len() int
	// Peek ,get and not remove the element from the top of the stack
	Peek() (interface{}, error)
	// Pop ,pop and remove the element from the top of the stack
	Pop() (interface{}, error)
	// Push ,push the element to top of the stack
	Push(val interface{}) error
	// Search ,return the index of the value in the stack
	Search(val interface{}) (int, error)
}
