package stack

import (
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
)

// node ,LinkedList element
type node struct {
	prev  *node
	next  *node
	value interface{}
}

// LinkedListStack ,Stack based on Double LinkedList
type LinkedListStack struct {
	top *node
}

// newNode ,init a node
func newNode(val interface{}) *node {
	return &node{
		prev:  nil,
		next:  nil,
		value: val,
	}
}

// NewLinkedListStack ,init LinkedListStack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{newNode(nil)}
}

// Flush ,clear the stack
func (s *LinkedListStack) Flush() {
	s.top = newNode(nil)
}

// IsEmpty ,determine whether the stack is empty
func (s *LinkedListStack) IsEmpty() bool {
	return nil == s.top.value
}

// Peek ,get and not remove the element from the top of the stack
func (s *LinkedListStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors2.StackIsEmptyError
	}
	cur := s.top
	for nil != cur.next {
		cur = cur.next
	}
	return cur.value, nil
}

// Pop ,pop and remove the element from the top of the stack
func (s *LinkedListStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors2.StackIsEmptyError
	}
	cur := s.top
	for nil != cur.next {
		cur = cur.next
	}
	pre := cur.prev
	pre.next = nil
	return cur.value, nil
}

// Push ,push the element to top of the stack
func (s *LinkedListStack) Push(val interface{}) {
	node := newNode(val)
	if s.IsEmpty() {
		s.top = node
		return
	}
	cur := s.top
	for nil != cur.next {
		cur = cur.next
	}
	cur.next, node.prev = node, cur
	return
}

// Search , return the index of the value in the stack
// Why return index when there is an error?
//All returned indexes are invalid. Avoid not checking error, but insist on using the returned value
func (s *LinkedListStack) Search(val interface{}) (int, error) {
	if s.IsEmpty() {
		return -1, errors2.StackIsEmptyError
	}
	if nil == s.top.next {
		if s.top.value == val {
			return 1, nil
		} else {
			return -2, errors2.NotExistError
		}
	}
	var index = 1
	cur := s.top
	for nil != cur.next {
		if cur.value == val {
			return index, nil
		}
		index++
		cur = cur.next
	}
	return index - 1, errors2.NotExistError
}
