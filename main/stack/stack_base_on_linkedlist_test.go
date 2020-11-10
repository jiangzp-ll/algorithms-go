package stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	stack := NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	if stack.topNode.value == 4 {
		t.Log("压栈成功")
	} else {
		t.Error("压栈失败")
	}
}

func TestLinkedListStack_Pop(t *testing.T) {
	stack := NewLinkedListStack()
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
	stack := NewLinkedListStack()
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