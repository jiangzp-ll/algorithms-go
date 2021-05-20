package errors

import "errors"

var (
	// General error
	NotExistError = errors.New("the value not exist")

	// Array error
	IndexOutOfBoundsError  = errors.New("array index out of bounds error")
	InvalidCapacityError   = errors.New("capacity must be greater than or equal to zero")
	SubArrayWithIndexError = errors.New("start must be less than or equal to end")

	// Single/Double LinkedList error
	LinkedListIsEmptyError = errors.New("LinkedList is empty")
	InputNodeIsEmptyError  = errors.New("input node is empty")
	InvalidIndexError      = errors.New("invalid index")
	ValueNotExistError     = errors.New("value exist in the LinkedList")

	// Stack
	StackIsEmptyError = errors.New("stack is empty")

	// Queue
	QueueIsFullError  = errors.New("queue is full")
	QueueIsEmptyError = errors.New("queue is empty")
)
