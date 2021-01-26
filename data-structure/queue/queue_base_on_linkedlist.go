package queue

import "fmt"

// 基于连边的顺序队列

// Node 节点
type Node struct {
	value interface{}
	next  *Node
}

// LinkedListQueue 队列
type LinkedListQueue struct {
	head   *Node
	tail   *Node
	length int
}

// NewLinkedListQueue 初始化
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// GetValue
func (n *Node) GetValue() interface{}{
	return n.value
}

// GetNext
func (n *Node) GetNext() *Node{
	return n.next
}

// GetHead
func (q *LinkedListQueue) GetHead() *Node {
	return q.head
}

// GetTail
func (q *LinkedListQueue) GetTail() *Node {
	return q.tail
}

// GetTail
func (q *LinkedListQueue) GetLength() int {
	return q.length
}

// EnQueue 入队
func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &Node{
		value: v,
		next:  nil,
	}
	if nil == q.tail {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

// DeQueue 出队
func (q *LinkedListQueue) DeQueue() interface{} {
	if nil == q.tail {
		return nil
	}
	old := q.head
	q.head = q.head.next
	q.length--
	return old.value
}

// ToString 转成字符串
func (q *LinkedListQueue) ToString() string {
	if q.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.value)
	}
	result += "<-tail"
	return result
}
