package errors

import "errors"

var (
	// Data Structure Sentinel Error
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

	// Queue error
	QueueIsFullError  = errors.New("queue is full")
	QueueIsEmptyError = errors.New("queue is empty")

	// Stack error
	StackIsEmptyError = errors.New("stack is empty")

	// Algorithm Sentinel Error
	// BinarySearch error
	InputSliceIsEmptyError = errors.New("slice not has element")
	InvalidInputValueError = errors.New("input value not in slice element range")
)
