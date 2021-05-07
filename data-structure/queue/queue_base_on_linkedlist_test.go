package queue

import (
	"fmt"
	"github.com/zepeng-jiang/go-basic-demo/data-structure/queue"
	"testing"
)

func TestLinkedListQueue_EnQueue(t *testing.T) {
	q := queue.NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	if q.GetHead().GetValue() == 0 && q.GetTail().GetValue() == 4 {
		t.Log("入队成功")
	} else {
		t.Error("入队失败")
	}
}

func TestLinkedListQueue_DeQueue(t *testing.T) {
	q := queue.NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	q.DeQueue()
	val := q.DeQueue()
	if val == 1 {
		t.Log("出队成功")
	} else {
		t.Error("出队失败")
	}
}

func TestLinkedListQueue_ToString(t *testing.T) {
	q := queue.NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}
	expect := "head<-0<-1<-2<-3<-4<-tail"
	result := q.ToString()
	fmt.Println(result)
	if expect == result {
		t.Log("转成字符串成功")
	} else {
		t.Error("转成字符串失败")
	}
}