package stack

import (
	"fmt"
	stack2 "github.com/zepeng-jiang/go-basic-demo/main/data-structure/stack"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	stack := stack2.NewArrayStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	if stack.GetTop() == 4 && stack.GetData()[stack.GetTop()] == 4{
		t.Log("压栈成功")
	} else {
		t.Error("压栈失败")
	}
}

func TestArrayStack_Pop(t *testing.T) {
	stack := stack2.NewArrayStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	var e int
	for i := 0; i < 5; i++ {
		e = stack.Pop().(int)
	}
	if e == 0 {
		t.Log("出栈成功")
	} else {
		t.Error("出栈失败")
	}
}

func TestArrayStack_Top(t *testing.T) {
	stack := stack2.NewArrayStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	e := stack.Top().(int)
	if e == 4 {
		t.Log("获取栈顶元素成功")
	} else {
		t.Error("获取栈顶元素失败")
	}
}

func TestArrayStack_Flush(t *testing.T) {
	stack := stack2.NewArrayStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	stack.Flush()
	fmt.Println("after:", stack.GetData())
	if stack.GetTop() == -1 && len(stack.GetData()) == 0{
		t.Log("清空栈成功")
	} else {
		t.Error("清空栈失败")
	}
}