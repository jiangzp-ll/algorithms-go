package stack

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

// LinkedListStack ,Stack based on Double LinkedList
type LinkedListStack struct {
	// top ,top of the stack
	top *node
	// len ,the number of elements in the stack
	len int
	// typeOf ,LinkedListStack type
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

// NewLinkedListStack ,init LinkedListStack
func NewLinkedListStack(typeOf string) (*LinkedListStack, error) {
	if typeOf == "" {
		return nil, errors2.InvalidTypeError
	}
	return &LinkedListStack{
		top:    newNode(nil),
		len:    0,
		typeOf: typeOf,
	}, nil
}

// Check ,check input value
func (s *LinkedListStack) Check(val interface{}) error {
	if nil == val {
		return errors2.InputValueCannotBeNilError
	}
	if err := generic.CheckType(s.typeOf, val); err != nil {
		return err
	}
	return nil
}

// Flush ,clear the stack
func (s *LinkedListStack) Flush() {
	s.top = newNode(nil)
}

// IsEmpty ,determine whether the stack is empty
func (s *LinkedListStack) IsEmpty() bool {
	return nil == s.top.value
}

// Len ,get the number of elements in the stack
func (s *LinkedListStack) Len() int {
	return s.len
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
	s.len--
	return cur.value, nil
}

// Push ,push the element to top of the stack
func (s *LinkedListStack) Push(val interface{}) error {
	if err := s.Check(val); err != nil {
		return err
	}
	node := newNode(val)
	if s.IsEmpty() {
		s.top = node
		s.len = 1
		return nil
	}
	cur := s.top
	for nil != cur.next {
		cur = cur.next
	}
	cur.next, node.prev = node, cur
	s.len++
	return nil
}

// Search , return the index of the value in the stack
// Why return index when there is an error?
//All returned indexes are invalid. Avoid not checking error, but insist on using the returned value
func (s *LinkedListStack) Search(val interface{}) (int, error) {
	if s.IsEmpty() {
		return -1, errors2.StackIsEmptyError
	}
	if err := s.Check(val); err != nil {
		return -1, err
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
