package queue

// Queue ,Queue interface{}
type Queue interface {
	// Add ,add the element to the end of the queue
	Add(val interface{}) error
	// Check ,check input value
	Check(val interface{}) error
	// Clear ,clear the queue
	Clear()
	// Contain ,determine whether the value contain in the queue
	Contain(val interface{}) bool
	// IsEmpty ,determine whether the queue is empty
	IsEmpty() bool
	// Len ,get the number of elements in the queue
	Len() int
	// Peek ,get and not remove the element from the header of the queue
	Peek() (interface{}, error)
	// Remove ,get and remove the element from the header of the queue
	Remove() (interface{}, error)
}
