package stack

import "fmt"

// 基于数组实现的栈

// ArrayStack 栈
type ArrayStack struct {
	// data 数据
	data []interface{}
	// top 栈顶指针
	top int
}

// NewArrayStack 初始化
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// GetTop
func (s *ArrayStack) GetTop() int {
	return s.top
}

// GetData
func (s *ArrayStack) GetData() []interface{} {
	return s.data
}

// IsEmpty 是否为空
func (s *ArrayStack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

// Push 压栈
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

// Pop 出栈
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top -= 1
	return v
}

// Top 返回栈顶元素
func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

// Flush
func (s *ArrayStack) Flush() {
	s.data = make([]interface{}, 0, 32)
	s.top = -1
}

// Print 打印栈
func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is null")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Printf("%v, ",s.data[i])
		}
		fmt.Println()
	}
}

