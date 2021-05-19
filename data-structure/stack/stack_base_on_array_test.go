package stack

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var as = NewArrayStack()

func Test_ArrayStack_Flush(t *testing.T) {
	for i := 1; i <= 5; i++ {
		as.Push(i)
	}
	as.Flush()
	if as.IsEmpty() {
		t.Log("as flush is success")
	} else {
		t.Error("as flush is failed")
	}
}

func Test_ArrayStack_IsEmpty(t *testing.T) {
	defer as.Flush()
	for i := 1; i <= 5; i++ {
		as.Push(i)
	}
	if !as.IsEmpty() {
		t.Log("as is not empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func Test_ArrayStack_Peek(t *testing.T) {
	defer as.Flush()
	elements := []int{1, 2, 3}
	pushValueToArrayStack(elements)
	ret, err := as.Peek()
	if err != nil {
		t.Errorf("peek top has error! error: %v \n", err)
		return
	}
	if ret == elements[len(elements)-1] && len(as.data) == len(elements) {
		t.Log("peek top is success")
	} else {
		t.Error("peek top is failed")
	}
}

func Test_ArrayStack_Peek_With_Stack_Is_Empty(t *testing.T) {
	_, err := as.Peek()
	if err != nil {
		if errors.Is(err, errors2.StackIsEmptyError) {
			t.Log("as is empty")
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
	defer as.Flush()
	elements := []int{1, 2, 3}
	pushValueToArrayStack(elements)
	ret, err := as.Pop()
	if err != nil {
		t.Errorf("pop the top has error! error: %v \n", err)
		return
	}
	if ret == elements[len(elements)-1] && len(as.data) == len(elements)-1 {
		t.Log("pop the top is success")
	} else {
		t.Error("pop the top is failed")
	}
}

func Test_ArrayStack_Pop_With_Stack_Is_Empty(t *testing.T) {
	_, err := as.Pop()
	if err != nil {
		if errors.Is(err, errors2.StackIsEmptyError) {
			t.Log("as is empty")
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
	defer as.Flush()
	flag := true
	elements := []int{1, 2, 3}
	pushValueToArrayStack(elements)
	for i := 0; i < len(as.data); i++ {
		if as.data[i] != elements[i] {
			flag = false
			break
		}
	}
	if !as.IsEmpty() && flag && cap(as.data) == BasicCapacity {
		t.Log("push element to as is success")
	} else {
		t.Error("push element to as is failed")
	}
}

func Test_ArrayStack_Push_With_Stack_Is_Full(t *testing.T) {
	defer as.Flush()
	var elements []int
	for i := 1; i < 20; i++ {
		elements = append(elements, i)
	}
	pushValueToArrayStack(elements)
	if !as.IsEmpty() && cap(as.data) == BasicCapacity*2 {
		t.Log("push element to as is success")
	} else {
		t.Error("push element to as is failed")
	}
}

func Test_ArrayStack_Push_With_Stack_Is_Full_And_Number_Greater_Than_1024(t *testing.T) {
	defer as.Flush()
	var elements []int
	for i := 1; i < 1026; i++ {
		elements = append(elements, i)
	}
	newCap := int(float64(len(elements)-1) * 1.25)
	pushValueToArrayStack(elements)
	if !as.IsEmpty() && cap(as.data) == newCap {
		t.Log("push element to as is success")
	} else {
		t.Error("push element to as is failed")
	}
}

func Test_ArrayStack_Search(t *testing.T) {
	defer as.Flush()
	target := 2
	elements := []int{1, target, 3}
	pushValueToArrayStack(elements)
	ret, err := as.Search(target)
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
	_, err := as.Search(1)
	if err != nil {
		if errors.Is(err, errors2.StackIsEmptyError) {
			t.Log("as is empty")
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
	defer as.Flush()
	elements := []int{1, 2, 3}
	pushValueToArrayStack(elements)
	_, err := as.Search("a")
	if err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Log("as is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Peek has bug")
	}
}

// pushValueToArrayStack .
func pushValueToArrayStack(elements []int) {
	for _, e := range elements {
		as.Push(e)
	}
}
