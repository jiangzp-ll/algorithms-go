package stack

import "fmt"

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
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// Data ,get stack all data
func (s *ArrayStack) Data() []interface{} {
	return s.data
}

// Flush ,clear the stack
func (s *ArrayStack) Flush() {
	s.data = make([]interface{}, 0, 32)
	s.top = -1
}

// IsEmpty ,determine whether the stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

// Pop ,pop the element from the top of the stack
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

// Push ,push the element to top of the stack
func (s *ArrayStack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top += 1
	}
	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

// Top ,get top of the stack
func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

// Print 打印栈
func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is null")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Printf("%v, ", s.data[i])
		}
		fmt.Println()
	}
}
