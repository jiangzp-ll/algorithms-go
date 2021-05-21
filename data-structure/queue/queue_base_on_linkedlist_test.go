package queue

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var lq = NewLinkedListQueue()

func Test_LinkedListQueue_Add(t *testing.T) {
	defer lq.Clear()
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	if lq.Len() == len(in) {
		t.Log("add element to queue is success")
	} else {
		t.Error("add element to queue is failed")
	}
}

func Test_LinkedListQueue_Add_With_Value_Is_Nil(t *testing.T) {
	if err := lq.Add(nil); err != nil {
		if errors.Is(err, errors2.InputValueCannotBeNilError) {
			t.Log("input value cannot be nil")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Add hsa bug")
	}
}

func Test_LinkedListQueue_Clear(t *testing.T) {
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	lq.Clear()
	if lq.Len() == 0 {
		t.Log("clear queue is success")
	} else {
		t.Error("clear queue is failed")
	}
}

func Test_LinkedListQueue_Contain(t *testing.T) {
	defer lq.Clear()
	target := 2
	in := []int{1, target, 3}
	addElementToLinkedListQueue(in)
	isContain := lq.Contain(target)
	if isContain {
		t.Log("queue has the value")
	} else {
		t.Error("queue not has the value")
	}
}

func Test_LinkedListQueue_Contain_With_Queue_Not_Exist_The_Value(t *testing.T) {
	defer lq.Clear()
	target := "a"
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	isContain := lq.Contain(target)
	if !isContain {
		t.Log("queue not has the value")
	} else {
		t.Error("function Contain has bug")
	}
}

func Test_LinkedListQueue_Contain_With_Queue_Is_Empty(t *testing.T) {
	target := "a"
	isContain := lq.Contain(target)
	if !isContain {
		t.Log("queue not has the value")
	} else {
		t.Error("function Contain has bug")
	}
}

func Test_LinkedListQueue_IsEmpty(t *testing.T) {
	isEmpty := lq.IsEmpty()
	if isEmpty {
		t.Log("queue is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func Test_LinkedListQueue_IsEmpty_With_Queue_Not_Empty(t *testing.T) {
	defer lq.Clear()
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	isEmpty := lq.IsEmpty()
	if !isEmpty {
		t.Log("queue not is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func Test_LinkedListQueue_Peek(t *testing.T) {
	defer lq.Clear()
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	ret, err := lq.Peek()
	if err != nil {
		t.Errorf("peek first value has error! error: %v \n", err)
		return
	}
	if ret == in[0] && lq.Len() == len(in) {
		t.Log("peek first value is success")
	} else {
		t.Error("peek first value is failed")
	}
}

func Test_LinkedListQueue_Peek_With_Queue_Is_Empty(t *testing.T) {
	_, err := lq.Peek()
	if err != nil {
		if errors.Is(err, errors2.QueueIsEmptyError) {
			t.Log("queue is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Peek hsa bug")
	}
}

func Test_LinkedListQueue_Remove(t *testing.T) {
	defer lq.Clear()
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	ret, err := lq.Remove()
	if err != nil {
		t.Errorf("remove first value has error! error: %v \n", err)
		return
	}
	if ret == in[0] && lq.Len() == len(in)-1 {
		t.Log("remove first value is success")
	} else {
		t.Error("remove first value is failed")
	}
}

func Test_LinkedListQueue_Remove_To_Queue_Is_Empty(t *testing.T) {
	defer lq.Clear()
	in := []int{1, 2, 3}
	addElementToLinkedListQueue(in)
	for i := 0; i < len(in); i++ {
		ret, err := lq.Remove()
		if err != nil {
			t.Errorf("remove first value has error! error: %v \n", err)
			return
		}
		if ret != in[i] {
			t.Error("remove first value is failed")
		}
	}
	if lq.IsEmpty() {
		t.Log("remove all element is success")
	} else {
		t.Error("remove all element is failed")
	}
}

func Test_LinkedListQueue_Remove_With_Queue_Is_Empty(t *testing.T) {
	_, err := lq.Remove()
	if err != nil {
		if errors.Is(err, errors2.QueueIsEmptyError) {
			t.Log("queue is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Remove hsa bug")
	}
}

// addElementToLinkedListQueue .
func addElementToLinkedListQueue(in []int) {
	for _, i := range in {
		_ = lq.Add(i)
	}
}
