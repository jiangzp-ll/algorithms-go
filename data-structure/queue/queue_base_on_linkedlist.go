package queue

import (
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"github.com/zepeng-jiang/go-basic-demo/pkg/generic"
)

// node ,LinkedList element
type node struct {
	prev  *node
	next  *node
	value interface{}
}

// LinkedListQueue ,Queue based on LinkedList
type LinkedListQueue struct {
	// head ,head Node
	head *node
	// tail ,tail Node
	tail *node
	// len ,the number of elements in the queue
	len int
	// typeOf , LinkedList type
	// Because Go not have Generic
	typeOf string
}

// newNode ,init a node
func newNode(val interface{}) *node {
	return &node{
		prev:  nil,
		next:  nil,
		value: val,
	}
}

// NewLinkedListQueue ,init LinkedListQueue
func NewLinkedListQueue(typeOf string) (*LinkedListQueue, error) {
	if typeOf == "" {
		return nil, errors2.InvalidTypeError
	}
	return &LinkedListQueue{
		head:   newNode(nil),
		tail:   newNode(nil),
		len:    0,
		typeOf: typeOf,
	}, nil
}

// Add ,add the element to the end of the queue
func (q *LinkedListQueue) Add(val interface{}) error {
	if nil == val {
		return errors2.InputValueCannotBeNilError
	}
	if err := q.Check(val); err != nil {
		return err
	}
	n := newNode(val)
	if q.len == 0 {
		q.head, q.tail = n, n
		q.len++
		return nil
	}
	t := q.tail
	t.next, n.prev = n, t
	q.tail = n
	q.len++
	return nil
}

// Check ,check input value type
func (q *LinkedListQueue) Check(val interface{}) error {
	if nil == val {
		return errors2.InputValueCannotBeNilError
	}
	if err := generic.CheckType(q.typeOf, val); err != nil {
		return err
	}
	return nil
}

// Clear ,clear the stack
func (q *LinkedListQueue) Clear() {
	q.head, q.tail = newNode(nil), newNode(nil)
	q.len = 0
	return
}

// Contain ,determine whether the value contain in the queue
func (q *LinkedListQueue) Contain(val interface{}) bool {
	if q.len == 0 || nil == val {
		return false
	}
	if err := q.Check(val); err != nil {
		return false
	}
	cur := q.head
	for nil != cur.next {
		if val == cur.value {
			return true
		}
		cur = cur.next
	}
	return false
}

// IsEmpty ,determine whether the queue is empty
func (q *LinkedListQueue) IsEmpty() bool {
	return q.len <= 0
}

// Len ,get the number of elements in the queue
func (q *LinkedListQueue) Len() int {
	return q.len
}

// Peek ,get and not remove the element from the header of the queue
func (q *LinkedListQueue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors2.QueueIsEmptyError
	}
	return q.head.value, nil
}

// Remove ,get and remove the element from the header of the queue
func (q *LinkedListQueue) Remove() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors2.QueueIsEmptyError
	}
	head := q.head
	if q.len == 1 {
		q.Clear()
		return head.value, nil
	}
	next := q.head.next
	q.head = next
	q.len--
	return head.value, nil
}
