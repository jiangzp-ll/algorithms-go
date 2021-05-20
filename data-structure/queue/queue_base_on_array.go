package queue

import errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"

// ArrayQueue ,Queue based on Array
type ArrayQueue struct {
	// data ,store data
	data []interface{}
	// len ,the number of elements in the queue
	len int
	// head ,the head of the queue
	head int
	// tail ,the tail of the queue
	tail int
}

// NewArrayQueue ,init ArrayQueue
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		data: make([]interface{}, n),
		len:  0,
		head: -1,
		tail: -1,
	}
}

// Add ,add the element to the end of the queue
func (q *ArrayQueue) Add(val interface{}) error {
	if len(q.data) <= q.len {
		return errors2.QueueIsFullError
	}
	if q.tail < 0 {
		q.data[0] = val
		q.head, q.tail = 0, 0
		q.len++
		return nil
	}
	q.data[q.len] = val
	q.tail++
	q.len++
	return nil
}

// Clear ,clear the stack
func (q *ArrayQueue) Clear() {
	q.data = make([]interface{}, len(q.data))
	q.len = 0
	q.head, q.tail = -1, -1
	return
}

// Contain ,determine whether the value contain in the queue
func (q *ArrayQueue) Contain(val interface{}) bool {
	if q.len == 0 {
		return false
	}
	for _, v := range q.data {
		if v == val {
			return true
		}
	}
	return false
}

// IsEmpty ,determine whether the queue is empty
func (q *ArrayQueue) IsEmpty() bool {
	return q.len <= 0
}

// Len ,get the number of elements in the queue
func (q *ArrayQueue) Len() int {
	return q.len
}

// Peek ,get and not remove the element from the header of the queue
func (q *ArrayQueue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors2.QueueIsEmptyError
	}
	return q.data[q.head], nil
}

// Remove ,get and remove the element from the header of the queue
func (q *ArrayQueue) Remove() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors2.QueueIsEmptyError
	}
	head := q.data[q.head]
	q.head++
	q.len--
	return head, nil
}
