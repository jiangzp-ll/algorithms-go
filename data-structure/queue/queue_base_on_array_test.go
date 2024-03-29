package queue

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var aq = initArrayQueue("int", 3)

func TestNewArrayQueue_With_TypeOf_Is_Empty(t *testing.T) {
	_, err := NewArrayQueue("", 3)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be not emtpy")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function NewArrayQueue hsa bug")
	}
}

func TestArrayQueue_Add(t *testing.T) {
	defer aq.Clear()
	flag := true
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	for i := 0; i < len(in); i++ {
		if in[i] != aq.data[i] {
			flag = false
			break
		}
	}
	if aq.Len() == len(in) && flag {
		t.Log("add element to queue is success")
	} else {
		t.Error("add element to queue is failed")
	}
}

func TestArrayQueue_Add_With_Queue_Is_Full(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	if err := aq.Add(4); err != nil {
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

func TestArrayQueue_Add_With_Value_Is_Nil(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	if err := aq.Add(nil); err != nil {
		if errors.Is(err, errors2.InputValueCannotBeNilError) {
			t.Log("input value can not be nil")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Add hsa bug")
	}
}

func TestArrayQueue_Add_With_Different_Type_Value(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	if err := aq.Add("q"); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Add hsa bug")
	}
}

func TestArrayQueue_Check_With_Input_Is_Nil(t *testing.T) {
	defer aq.Clear()
	_ = aq.Add("a")
	if err := aq.Check(nil); err != nil {
		t.Log("input value type not same with ArrayQueue type.")
	} else {
		t.Error("function Check has bug")
	}
}

func TestArrayQueue_Clear(t *testing.T) {
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	aq.Clear()
	if aq.Len() == 0 {
		t.Log("clear queue is success")
	} else {
		t.Error("clear queue is failed")
	}
}

func TestArrayQueue_Contain(t *testing.T) {
	defer aq.Clear()
	target := 2
	in := []int{1, target, 3}
	addElementToArrayQueue(in)
	isContain := aq.Contain(target)
	if isContain {
		t.Log("queue has the value")
	} else {
		t.Error("queue not has the value")
	}
}

func TestArrayQueue_Contain_With_Queue_Not_Exist_The_Value(t *testing.T) {
	defer aq.Clear()
	target := 4
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	isContain := aq.Contain(target)
	if !isContain {
		t.Log("queue not has the value")
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArrayQueue_Contain_With_Queue_Is_Empty(t *testing.T) {
	defer aq.Clear()
	target := 1
	isContain := aq.Contain(target)
	if !isContain {
		t.Log("queue is empty")
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArrayQueue_Contain_With_Value_Is_Nil(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	isContain := aq.Contain(nil)
	if !isContain {
		t.Log("input value can not be nil")
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArrayQueue_Contain_With_Different_Type_Value(t *testing.T) {
	defer aq.Clear()
	target := "a"
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	isContain := aq.Contain(target)
	if !isContain {
		t.Log("queue not has the value")
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArrayQueue_IsEmpty(t *testing.T) {
	isEmpty := aq.IsEmpty()
	if isEmpty {
		t.Log("queue is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArrayQueue_IsEmpty_With_Queue_Not_Empty(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	isEmpty := aq.IsEmpty()
	if !isEmpty {
		t.Log("queue not is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArrayQueue_Peek(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	ret, err := aq.Peek()
	if err != nil {
		t.Errorf("peek first value has error! error: %v \n", err)
		return
	}
	if ret == in[0] && aq.Len() == len(in) {
		t.Log("peek first value is success")
	} else {
		t.Error("peek first value is failed")
	}
}

func TestArrayQueue_Peek_With_Queue_Is_Empty(t *testing.T) {
	_, err := aq.Peek()
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

func TestArrayQueue_Remove(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	ret, err := aq.Remove()
	if err != nil {
		t.Errorf("remove first value has error! error: %v \n", err)
		return
	}
	if ret == in[0] && aq.Len() == len(in)-1 {
		t.Log("remove first value is success")
	} else {
		t.Error("remove first value is failed")
	}
}

func TestArrayQueue_Remove_To_Queue_Is_Empty(t *testing.T) {
	defer aq.Clear()
	in := []int{1, 2, 3}
	addElementToArrayQueue(in)
	for i := 0; i < len(in); i++ {
		ret, err := aq.Remove()
		if err != nil {
			t.Errorf("remove first value has error! error: %v \n", err)
			return
		}
		if ret != in[i] {
			t.Error("remove first value is failed")
		}
	}
	if aq.IsEmpty() {
		t.Log("remove all element is success")
	} else {
		t.Error("remove all element is failed")
	}
}

func TestArrayQueue_Remove_With_Queue_Is_Empty(t *testing.T) {
	_, err := aq.Remove()
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

// initArrayQueue .
func initArrayQueue(typeof string, c int) *ArrayQueue {
	queue, _ := NewArrayQueue(typeof, c)
	return queue
}

// addElementToArrayQueue .
func addElementToArrayQueue(in []int) {
	for _, i := range in {
		_ = aq.Add(i)
	}
}
