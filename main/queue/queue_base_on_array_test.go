package queue

import (
	"testing"
)

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(5)
	for i := 0; i < 5; i++ {
		_ = q.EnQueue(i)
	}
	if q.queue[q.head] == 0 && q.queue[q.tail-1] == 4 {
		t.Log("入队成功")
	} else {
		t.Error("入队失败")
	}
}

func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(5)
	for i := 0; i < 5; i++ {
		_ = q.EnQueue(i)
	}
	first := q.DeQueue()
	if first == 0 {
		t.Log("出队成功")
	} else {
		t.Error("出队失败")
	}
}

func TestArrayQueue_ToString(t *testing.T) {
	q := NewArrayQueue(5)
	for i := 0; i < 5; i++ {
		_ = q.EnQueue(i)
	}
	expect := "head<-0<-1<-2<-3<-4<-tail"
	result := q.ToString()
	if expect == result {
		t.Log("转成字符串成功")
	} else {
		t.Error("转成字符串失败")
	}
}