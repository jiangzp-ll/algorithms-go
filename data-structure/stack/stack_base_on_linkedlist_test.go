package stack

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var ls = initLinkedList("int")

func TestNewLinkedListStack_With_TypeOf_Is_Empty(t *testing.T) {
	_, err := NewLinkedListStack("")
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be not emtpy")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function NewLinkedListStack hsa bug")
	}
}

func TestLinkedListStack_Flush(t *testing.T) {
	for i := 1; i <= 5; i++ {
		_ = ls.Push(i)
	}
	ls.Flush()
	if ls.IsEmpty() {
		t.Log("stack flush is success")
	} else {
		t.Error("stack flush is failed")
	}
}

func TestLinkedListStack_IsEmpty(t *testing.T) {
	defer ls.Flush()
	for i := 1; i <= 5; i++ {
		_ = ls.Push(i)
	}
	if !ls.IsEmpty() {
		t.Log("stack is not empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestLinkedListStack_Len(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	for _, e := range elements {
		_ = ls.Push(e)
	}
	if ls.Len() == len(elements) {
		t.Log("get LinkedListStack len is success")
	} else {
		t.Error("function Len has bug")
	}
}

func TestLinkedListStack_Peek(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	ret, err := ls.Peek()
	if err != nil {
		t.Errorf("peek top has error! error: %v \n", err)
		return
	}
	if ret == elements[len(elements)-1] {
		t.Log("peek top is success")
	} else {
		t.Error("peek top is failed")
	}
}

func TestLinkedListStack_Peek_With_Stack_Is_Empty(t *testing.T) {
	_, err := ls.Peek()
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

func TestLinkedListStack_Pop(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	ret, err := ls.Pop()
	if err != nil {
		t.Errorf("pop the top has error! error: %v \n", err)
		return
	}
	if ret == elements[len(elements)-1] {
		t.Log("pop the top is success")
	} else {
		t.Error("pop the top is failed")
	}
}

func TestLinkedListStack_Pop_With_Stack_Is_Empty(t *testing.T) {
	_, err := ls.Pop()
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

func TestLinkedListStack_Push(t *testing.T) {
	defer ls.Flush()
	flag := true
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	cur := ls.top
	for i := 0; i < len(elements); i++ {
		if elements[i] != cur.value {
			flag = false
			break
		}
		cur = cur.next
	}
	if !ls.IsEmpty() && flag {
		t.Log("push element to stack is success")
	} else {
		t.Error("push element to stack is failed")
	}
}

func TestLinkedListStack_Push_With_Stack_Is_Full(t *testing.T) {
	defer ls.Flush()
	var elements []int
	for i := 1; i < 20; i++ {
		elements = append(elements, i)
	}
	pushValueToLinkedListStack(elements)
	if !ls.IsEmpty() {
		t.Log("push element to stack is success")
	} else {
		t.Error("push element to stack is failed")
	}
}

func TestLinkedListStack_Push_With_Stack_Is_Full_And_Number_Greater_Than_1024(t *testing.T) {
	defer ls.Flush()
	var elements []int
	for i := 1; i < 1026; i++ {
		elements = append(elements, i)
	}
	pushValueToLinkedListStack(elements)
	if !ls.IsEmpty() {
		t.Log("push element to stack is success")
	} else {
		t.Error("push element to stack is failed")
	}
}

func TestLinkedListStack_Push_With_Value_Is_Nil(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	if err := ls.Push(nil); err != nil {
		if errors.Is(err, errors2.InputValueCannotBeNilError) {
			t.Logf("input value can not be nil")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Push has bug")
	}
}

func TestLinkedListStack_Push_With_Different_Type_Value(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	if err := ls.Push("a"); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Push has bug")
	}
}

func TestLinkedListStack_Search(t *testing.T) {
	defer ls.Flush()
	target := 2
	elements := []int{1, target, 3}
	pushValueToLinkedListStack(elements)
	ret, err := ls.Search(target)
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

func TestLinkedListStack_Search_With_Stack_Only_One_Element(t *testing.T) {
	defer ls.Flush()
	target := 1
	elements := []int{target}
	pushValueToLinkedListStack(elements)
	ret, err := ls.Search(target)
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

func TestLinkedListStack_Search_With_Stack_Only_One_Element_And_Not_In_Stack(t *testing.T) {
	defer ls.Flush()
	target := 2
	elements := []int{1}
	pushValueToLinkedListStack(elements)
	_, err := ls.Search(target)
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
func TestLinkedListStack_Search_With_Stack_Is_Empty(t *testing.T) {
	_, err := ls.Search(1)
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

func TestLinkedListStack_Search_With_The_Value_Not_Exist(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	_, err := ls.Search(4)
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

func TestLinkedListStack_Search_With_Value_Is_Nil(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	_, err := ls.Search(nil)
	if err != nil {
		if errors.Is(err, errors2.InputValueCannotBeNilError) {
			t.Logf("input value can not be nil")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Push has bug")
	}
}

func TestLinkedListStack_Search_With_Different_Type_Value(t *testing.T) {
	defer ls.Flush()
	elements := []int{1, 2, 3}
	pushValueToLinkedListStack(elements)
	_, err := ls.Search("a")
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Peek has bug")
	}
}

// initLinkedList .
func initLinkedList(typeOf string) *LinkedListStack {
	s, _ := NewLinkedListStack(typeOf)
	return s
}

// pushValueToLinkedListStack .
func pushValueToLinkedListStack(elements []int) {
	for _, e := range elements {
		_ = ls.Push(e)
	}
}
