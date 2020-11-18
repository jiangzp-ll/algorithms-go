package queue

import "fmt"

// 数组实现顺序队列

// ArrayQueue 队列
type ArrayQueue struct {
	queue    []interface{}
	capacity int
	head     int
	tail     int
}

// NewArrayQueue 初始化
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		queue:    make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// EnQueue 入队
func (q *ArrayQueue) EnQueue(v interface{}) bool {
	if q.capacity == q.tail {
		return false
	}
	q.queue[q.tail] = v
	q.tail++
	return true
}

// DeQueue 出队
func (q *ArrayQueue) DeQueue() interface{} {
	if q.tail == q.head {
		return nil
	}
	h := q.queue[q.head]
	q.head++
	return h
}

// ToString 转成字符串
func (q *ArrayQueue) ToString() string {
	if q.head == q.tail {
		return "empty queue"
	}
	ret := "head"
	for i := q.head; i < q.tail; i++ {
		ret += fmt.Sprintf("<-%+v", q.queue[i])
	}
	ret+="<-tail"
	return ret
}
