package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue(5)
	for i := 0; i < 5; i++ {
		_ = q.EnQueue(i)
	}
	fmt.Println(q.GetQueue()[q.GetTail()-1])
	if q.GetQueue()[q.GetHead()] == 0 && q.GetQueue()[q.GetTail()-1] == 3 {
		t.Log("入队成功")
	} else {
		t.Error("入队失败")
	}
}

func TestCircularQueue_DeQueue(t *testing.T) {
	q := NewCircularQueue(6)
	for i := 0; i < 5; i++ {
		_ = q.EnQueue(i)
	}
	ret := q.DeQueue()
	if ret == 0 {
		t.Log("出队成功")
	} else {
		t.Error("出队失败")
	}
}

func TestCircularQueue_ToString(t *testing.T) {
	q := NewCircularQueue(6)
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
