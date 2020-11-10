package stack

import "fmt"

// 基于链表实现的栈

// Node 节点
type Node struct {
	// 	next 指针
	next *Node
	// 数据
	value interface{}
}

// LinkedListStack 栈
type LinkedListStack struct {
	topNode *Node
}

// NewLinkedListStack 初始化栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

// IsEmpty 是否为空
func (s *LinkedListStack) IsEmpty() bool {
	return s.topNode == nil
}

// Push 压栈
func (s *LinkedListStack) Push(v interface{}) {
	s.topNode = &Node{
		next:  s.topNode,
		value: v,
	}
}

// Pop 出栈
func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.topNode.value
	s.topNode = s.topNode.next
	return v
}

// Top 返回栈顶元素
func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.topNode.value
}

// Flush 清空栈
func (s LinkedListStack) Flush() {
	s.topNode = nil
}

// Print 打印栈
func (s *LinkedListStack) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is null")
	} else {
		cur := s.topNode
		for nil != cur {
			fmt.Printf("%v, ", cur.value)
			cur = cur.next
		}
	}
}
