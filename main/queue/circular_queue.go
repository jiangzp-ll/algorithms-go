package queue

import "fmt"

// 循环队列

// CircularQueue 循环队列
type CircularQueue struct {
	queue    []interface{}
	capacity int
	head     int
	tail     int
}

// NewCircularQueue 初始化
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{
		queue:    make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// IsEmpty 是否为空
// 栈空条件：head==tail为true
func (q *CircularQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

// IsFull 是否排满
// 栈满条件：(tail+1)%capacity==head为true
func (q *CircularQueue) IsFull() bool {
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

// EnQueue 入队
func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.queue[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// DeQueue 出队
func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.queue[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

// ToString 转成字符串
func (q *CircularQueue) ToString() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.queue[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
