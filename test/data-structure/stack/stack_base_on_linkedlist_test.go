package stack

import (
	stack2 "github.com/zepeng-jiang/go-basic-demo/main/data-structure/stack"
	"testing"
)

func TestLinkedListStack_Push(t *testing.T) {
	stack := stack2.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	if stack.GetTopNode().GetValue() == 4 {
		t.Log("压栈成功")
	} else {
		t.Error("压栈失败")
	}
}

func TestLinkedListStack_Pop(t *testing.T) {
	stack := stack2.NewLinkedListStack()
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

func TestLinkedListStack_Top(t *testing.T) {
	stack := stack2.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	e := stack.Top().(int)
	if e == 4 {
		t.Log("返回栈顶元素成功")
	} else {
		t.Error("返回栈顶元素失败")
	}
}

func TestLinkedListStack_Flush(t *testing.T) {
	stack := stack2.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	stack.Flush()
	if nil == stack.GetTopNode() {
		t.Log("清空栈成功")
	} else {
		t.Error("清空栈失败")
	}
}