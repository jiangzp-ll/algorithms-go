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
