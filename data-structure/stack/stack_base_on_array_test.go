package stack

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var stack = NewArrayStack()

func Test_ArrayStack_Flush(t *testing.T) {
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}
	stack.Flush()
	if stack.IsEmpty() {
		t.Log("stack flush is success")
	} else {
		t.Error("stack flush is failed")
	}
}

func Test_ArrayStack_IsEmpty(t *testing.T) {
	defer stack.Flush()
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}
	if !stack.IsEmpty() {
		t.Log("stack is not empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func Test_ArrayStack_Peek(t *testing.T) {
	defer stack.Flush()
	elements := []int{1, 2, 3}
	pushValueToStack(elements)
	ret, err := stack.Peek()
	if err != nil {
		t.Errorf("peek top has error! error: %v \n", err)
		return
	}
	if ret == elements[len(elements)-1] && len(stack.data) == len(elements) {
		t.Log("peek top is success")
	} else {
		t.Error("peek top is failed")
	}
}

func Test_ArrayStack_Peek_With_Stack_Is_Empty(t *testing.T) {
	_, err := stack.Peek()
	if err != nil {
		if errors.Is(err, errors2.StackIsEmptyError) {
			t.Log("stack is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Peek has bug")
	}
}

func Test_ArrayStack_Pop(t *testing.T) {
	defer stack.Flush()
	elements := []int{1, 2, 3}
	pushValueToStack(elements)
	ret, err := stack.Pop()
	if err != nil {
		t.Errorf("pop the top has error! error: %v \n", err)
		return
	}
	if ret == elements[len(elements)-1] && len(stack.data) == len(elements)-1 {
		t.Log("pop the top is success")
	} else {
		t.Error("pop the top is failed")
	}
}

func Test_ArrayStack_Pop_With_Stack_Is_Empty(t *testing.T) {
	_, err := stack.Pop()
	if err != nil {
		if errors.Is(err, errors2.StackIsEmptyError) {
			t.Log("stack is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Peek has bug")
	}
}

func Test_ArrayStack_Push(t *testing.T) {
	defer stack.Flush()
	flag := true
	elements := []int{1, 2, 3}
	pushValueToStack(elements)
	for i := 0; i < len(stack.data); i++ {
		if stack.data[i] != elements[i] {
			flag = false
			break
		}
	}
	if !stack.IsEmpty() && flag && cap(stack.data) == BasicCapacity {
		t.Log("push element to stack is success")
	} else {
		t.Error("push element to stack is failed")
	}
}

func Test_ArrayStack_Push_With_Stack_Is_Full(t *testing.T) {
	defer stack.Flush()
	var elements []int
	for i := 1; i < 20; i++ {
		elements = append(elements, i)
	}
	pushValueToStack(elements)
	if !stack.IsEmpty() && cap(stack.data) == BasicCapacity*2 {
		t.Log("push element to stack is success")
	} else {
		t.Error("push element to stack is failed")
	}
}

func Test_ArrayStack_Push_With_Stack_Is_Full_And_Number_Greater_Than_1024(t *testing.T) {
	defer stack.Flush()
	var elements []int
	for i := 1; i < 1026; i++ {
		elements = append(elements, i)
	}
	newCap := int(float64(len(elements)-1) * 1.25)
	pushValueToStack(elements)
	if !stack.IsEmpty() && cap(stack.data) == newCap {
		t.Log("push element to stack is success")
	} else {
		t.Error("push element to stack is failed")
	}
}

func Test_ArrayStack_Search(t *testing.T) {
	defer stack.Flush()
	target := 2
	elements := []int{1, target, 3}
	pushValueToStack(elements)
	ret, err := stack.Search(target)
	if err != nil {
		t.Errorf("search the value has error! error: %v \n", err)
		return
	}
	if ret == target {
		t.Log("search the value is success")
	} else {
		t.Errorf("search the value is failed")
	}
}

func Test_ArrayStack_Search_With_Stack_Is_Empty(t *testing.T) {
	_, err := stack.Search(1)
	if err != nil {
		if errors.Is(err, errors2.StackIsEmptyError) {
			t.Log("stack is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Peek has bug")
	}
}

func Test_ArrayStack_Search_With_The_Value_Not_Exist(t *testing.T) {
	defer stack.Flush()
	elements := []int{1, 2, 3}
	pushValueToStack(elements)
	_, err := stack.Search("a")
	if err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Log("stack is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Peek has bug")
	}
}

// pushValueToStack .
func pushValueToStack(elements []int) {
	for _, e := range elements {
		stack.Push(e)
	}
}
