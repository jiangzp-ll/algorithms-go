package errors

import "errors"

var (
	// General error
	NotExistError              = errors.New("the value not exist")
	InvalidCapacityError       = errors.New("capacity must be greater than or equal to zero")
	InvalidTypeError           = errors.New("invalid type")
	InputValueCannotBeNilError = errors.New("input value cannot be nil")

	// ArrayList error
	IndexOutOfBoundsError      = errors.New("array index out of bounds error")
	SubArrayListWithIndexError = errors.New("start must be less than or equal to end")

	// Single/Double LinkedList error
	LinkedListIsEmptyError = errors.New("LinkedList is empty")
	InputNodeIsEmptyError  = errors.New("input node is empty")
	InvalidIndexError      = errors.New("invalid index")
	ValueNotExistError     = errors.New("value exist in the LinkedList")

	// Queue
	QueueIsFullError  = errors.New("queue is full")
	QueueIsEmptyError = errors.New("queue is empty")

	// Stack
	StackIsEmptyError = errors.New("stack is empty")
)
