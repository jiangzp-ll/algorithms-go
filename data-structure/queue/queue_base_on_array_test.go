package queue

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var q = NewArrayQueue(3)

func Test_ArrayQueue_Add(t *testing.T) {
	defer q.Clear()
	flag := true
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	for i := 0; i < len(in); i++ {
		if in[i] != q.data[i] {
			flag = false
			break
		}
	}
	if q.Len() == len(in) && flag {
		t.Log("add element to queue is success")
	} else {
		t.Error("add element to queue is failed")
	}
}

func Test_ArrayQueue_Add_With_Queue_Is_Full(t *testing.T) {
	defer q.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	if err := q.Add(4); err != nil {
		if errors.Is(err, errors2.QueueIsFullError) {
			t.Log("queue is full")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Add hsa bug")
	}
}

func Test_ArrayQueue_Clear(t *testing.T) {
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	q.Clear()
	if q.Len() == 0 {
		t.Log("clear queue is success")
	} else {
		t.Error("clear queue is failed")
	}
}

func Test_ArrayQueue_Contain(t *testing.T) {
	defer q.Clear()
	target := 2
	in := []int{1, target, 3}
	addElementToArrayQueue(in)
	isContain := q.Contain(target)
	if isContain {
		t.Log("queue has the value")
	} else {
		t.Error("queue not has the value")
	}
}

func Test_ArrayQueue_Contain_With_Queue_Not_Exist_The_Value(t *testing.T) {
	defer q.Clear()
	target := "a"
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	isContain := q.Contain(target)
	if !isContain {
		t.Log("queue not has the value")
	} else {
		t.Error("function Contain has bug")
	}
}

func Test_ArrayQueue_Contain_With_Queue_Is_Empty(t *testing.T) {
	target := "a"
	isContain := q.Contain(target)
	if !isContain {
		t.Log("queue not has the value")
	} else {
		t.Error("function Contain has bug")
	}
}

func Test_ArrayQueue_IsEmpty(t *testing.T) {
	isEmpty := q.IsEmpty()
	if isEmpty {
		t.Log("queue is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func Test_ArrayQueue_IsEmpty_With_Queue_Not_Empty(t *testing.T) {
	defer q.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	isEmpty := q.IsEmpty()
	if !isEmpty {
		t.Log("queue not is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func Test_ArrayQueue_Peek(t *testing.T) {
	defer q.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	ret, err := q.Peek()
	if err != nil {
		t.Errorf("peek first value has error! error: %v \n", err)
		return
	}
	if ret == in[0] && q.Len() == len(in) {
		t.Log("peek first value is success")
	} else {
		t.Error("peek first value is failed")
	}
}

func Test_ArrayQueue_Peek_With_Queue_Is_Empty(t *testing.T) {
	_, err := q.Peek()
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

func Test_ArrayQueue_Remove(t *testing.T) {
	defer q.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	ret, err := q.Remove()
	if err != nil {
		t.Errorf("remove first value has error! error: %v \n", err)
		return
	}
	if ret == in[0] && q.Len() == len(in)-1 {
		t.Log("remove first value is success")
	} else {
		t.Error("remove first value is failed")
	}
}

func Test_ArrayQueue_Remove_To_Queue_Is_Empty(t *testing.T) {
	defer q.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	for i := 0; i < len(in); i++ {
		ret, err := q.Remove()
		if err != nil {
			t.Errorf("remove first value has error! error: %v \n", err)
			return
		}
		if ret != in[i] {
			t.Error("remove first value is failed")
		}
	}
	if q.IsEmpty() {
		t.Log("remove all element is success")
	} else {
		t.Error("remove all element is failed")
	}
}

func Test_ArrayQueue_Remove_With_Queue_Is_Empty(t *testing.T) {
	_, err := q.Remove()
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

// addElementToArrayQueue .
func addElementToArrayQueue(in []int) {
	for _, i := range in {
		_ = q.Add(i)
	}
}
