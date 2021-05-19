package stack

import errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"

const BasicCapacity = 16

// ArrayStack ,Stack based on Array
type ArrayStack struct {
	// data ,store data
	data []interface{}
	// top ,top of the stack
	top int
}

// NewArrayStack ,init ArrayStack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, BasicCapacity),
		top:  -1,
	}
}

// expansion ,stack expansion
// When the number of elements in the stack is less than 1024, the capacity becomes twice of the original stack;
// When the number of elements in the stack is greater than 1024, the capacity becomes 1.25 times of the original stack.
func (s *ArrayStack) expansion() {
	length := len(s.data)
	if length < 1024 {
		newData := make([]interface{}, length, 2*length)
		for i := 0; i < length; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
	} else {
		nl := float64(length) * 1.25
		newData := make([]interface{}, length, int(nl))
		for i := 0; i < length; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
	}
}

// Flush ,clear the stack
func (s *ArrayStack) Flush() {
	s.data = make([]interface{}, 0, BasicCapacity)
	s.top = -1
}

// IsEmpty ,determine whether the stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

// Peek ,get and not remove the element from the top of the stack
func (s *ArrayStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors2.StackIsEmptyError
	}
	v := s.data[s.top]
	return v, nil
}

// Pop ,pop and remove the element from the top of the stack
func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors2.StackIsEmptyError
	}
	v := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return v, nil
}

// Push ,push the element to top of the stack
func (s *ArrayStack) Push(val interface{}) {
	if s.top == cap(s.data)-1 {
		s.expansion()
	}
	s.data = append(s.data, val)
	s.top++
}

// Search , return the index of the value in the stack
// Why return index when there is an error?
//All returned indexes are invalid. Avoid not checking error, but insist on using the returned value
func (s *ArrayStack) Search(val interface{}) (int, error) {
	if s.IsEmpty() {
		return -1, errors2.StackIsEmptyError
	}
	var index int
	for i := 0; i < len(s.data); i++ {
		if val == s.data[i] {
			index = i + 1
			break
		}
	}
	if index == 0 {
		return -2, errors2.NotExistError
	}
	return index, nil
}
