package arraylist

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"math/rand"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	typeOf := "int"
	arr, err := NewArrayList(typeOf)
	if err != nil {
		t.Errorf("unkown error! error: %v ", err)
	}
	if arr.Len() == 0 && cap(arr.Data()) == DefaultCapacity {
		t.Log("init a ArrayList is success")
	} else {
		t.Error("function NewArrayList has bug")
	}
}

func TestNewArrayList_With_TypeOf_Is_Empty(t *testing.T) {
	_, err := NewArrayList("")
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be not empty")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Errorf("function NewArrayList has bug")
	}
}

func TestNewArrayListWithCap_With_Capacity_Greater_Than_Zero(t *testing.T) {
	capacity := 2
	typeOf := "string"
	arr, err := NewArrayListWithCap(typeOf, capacity)
	if err != nil {
		t.Errorf("init ArrayList has error! error: %v ", err)
		return
	}
	if arr.Len() == 0 && len(arr.Data()) == capacity {
		t.Log("init ArrayList success")
	} else {
		t.Error("init ArrayList failed")
	}
}

func TestNewArrayListWithCap_With_Capacity_Equal_Zero(t *testing.T) {
	capacity := 0
	typeOf := "string"
	arr, err := NewArrayListWithCap(typeOf, capacity)
	if err != nil {
		t.Errorf("init ArrayList has error! error: %v ", err)
		return
	}
	if arr.Len() == 0 && len(arr.Data()) == capacity {
		t.Log("init ArrayList success")
	} else {
		t.Error("init ArrayList failed")
	}
}

func TestNewArrayListWithCap_With_TypeOf_Is_Empty(t *testing.T) {
	capacity := 0
	typeOf := ""
	_, err := NewArrayListWithCap(typeOf, capacity)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be not empty")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Errorf("function NewArrayListWithCap has bug")
	}
}

func TestNewArrayListWithCap_With_Capacity_Less_Than_Zero(t *testing.T) {
	capacity := -1
	typeOf := "string"
	_, err := NewArrayListWithCap(typeOf, capacity)
	if err != nil {
		if errors.Is(err, errors2.InvalidCapacityError) {
			t.Logf("capacity must be greater than zero")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Errorf("function NewArrayListWithCap has bug")
	}
}

func TestArrayList_Add_With_Not_Need_Expansion(t *testing.T) {
	in := 1
	capacity := 2
	typeOf := "int"
	arr := initArrayList(typeOf, capacity, t)
	if err := arr.Add(in); err != nil {
		t.Errorf("add value to ArrayList has error! error: %v ", err)
		return
	}
	if len(arr.data) == capacity && arr.Len() == 1 && arr.data[0] == in {
		t.Logf("add %d to ArrayList is success", in)
	} else {
		t.Errorf("add %d to ArrayList is failed", in)
	}
}

func TestArrayList_Add_With_Need_Expansion_And_ArrayList_Length_Less_Than_1024(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 2
	typeOf := "string"
	flag := true
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	for i := 0; i < len(in); i++ {
		if in[i] != arr.data[i] {
			flag = false
			break
		}
	}
	if len(arr.data) == capacity*2 && arr.Len() == 3 && flag {
		t.Logf("add %v to ArrayList is success", in)
	} else {
		t.Errorf("add %v to ArrayList is failed", in)
	}
}

func TestArrayList_Add_With_Need_Expansion_And_ArrayList_Length_Greater_Than_1024(t *testing.T) {
	capacity := 1024
	typeOf := "int"
	arr := initArrayList(typeOf, capacity, t)
	for i := 0; i <= capacity; i++ {
		_ = arr.Add(i)
	}
	if len(arr.data) == int(float64(capacity)*1.25) && arr.Len() == 1025 {
		t.Log("function Add is good")
	} else {
		t.Error("function Add has bug")
	}
}

func TestArrayList_Add_With_Need_Expansion_And_ArrayList_Is_Empty(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 0
	typeOf := "string"
	flag := true
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	for i := 0; i < len(in); i++ {
		if in[i] != arr.data[i] {
			flag = false
			break
		}
	}
	if len(arr.data) == 4 && arr.Len() == 3 && flag {
		t.Logf("add %v to ArrayList is success", in)
	} else {
		t.Errorf("add %v to ArrayList is failed", in)
	}
}

func TestArrayList_Add_With_Different_Type(t *testing.T) {
	in := "a"
	capacity := 2
	typeOf := "int"
	arr := initArrayList(typeOf, capacity, t)
	if err := arr.Add(in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("The type of element added must be the same as the type of ArrayList")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Errorf("function Add has bug")
	}
}

func TestArrayList_Check_With_Input_Is_Nil(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	if err := arr.Check(nil); err != nil {
		t.Log("input value type not same with Array type.")
	} else {
		t.Error("function Check has bug")
	}
}

func TestArrayList_Clear(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	flag := true
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	arr.Clear()
	for i := 0; i < len(arr.data); i++ {
		if nil != arr.data[i] {
			flag = false
			break
		}
	}
	if arr.Len() == capacity && flag {
		t.Log("clear ArrayList is success")
	} else {
		t.Error("clear ArrayList is failed")
	}
}

func TestArrayList_Contains_With_Contain_The_Value(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	target := "a"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	ret := arr.Contain(target)
	if ret {
		t.Logf("ArrayList contain %s \n", target)
	} else {
		t.Errorf("ArrayList not contain %s \n", target)
	}
}

func TestArrayList_Contains_With_Not_Contain_The_Value(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	target := "A"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	ret := arr.Contain(target)
	if !ret {
		t.Logf("ArrayList not contain %s \n", target)
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArrayList_Contains_With_Different_Type_Value(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	target := 1
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	ret := arr.Contain(target)
	if !ret {
		t.Logf("ArrayList not contain %d \n", target)
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArrayList_Data(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	flag := true
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	data := arr.Data()
	for i := 0; i < len(data); i++ {
		if data[i] != in[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("ArrayList is %v \n", data)
	} else {
		t.Error("func Data has bug")
	}
}

func TestArrayList_Get_With_Index_In_The_Array(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	index := rand.Intn(3)
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	val, err1 := arr.Get(index)
	if err1 != nil {
		t.Errorf("get value is failed! error: %v", err1)
		return
	}
	if val == in[index] {
		t.Logf("get the index is %d elements in the ArrayList success", index)
	} else {
		t.Errorf("get the index is %d elements in the ArrayList failed", index)
	}
}

func TestArrayList_Get_With_Index_Not_In_The_Array(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	index := capacity + 1
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.Get(index)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Log("invalid index")
		} else {
			t.Errorf("unknown error! error: %v \n", err1)
		}
	} else {
		t.Error("function Get has bug")
	}
}

func TestArrayList_IndexOf_With_ArrayList_Has_Existed_The_Value(t *testing.T) {
	target := "c"
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	index, err1 := arr.IndexOf(target)
	if err1 != nil {
		t.Errorf("get index failed! error: %v \n", err1)
		return
	}
	if index == arr.Len()-1 {
		t.Logf("the element with index %d in the ArrayList is %s \n", index, target)
	} else {
		t.Error("function IndexOf has bug")
	}
}

func TestArrayList_IndexOf_With_ArrayList_Not_Exist_The_Value(t *testing.T) {
	target := "C"
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.IndexOf(target)
	if err1 != nil {
		t.Logf("ArrayList not exist %s \n", target)
	} else {
		t.Error("function IndexOf has bug")
	}
}

func TestArrayList_IndexOf_With_Different_Type_Value(t *testing.T) {
	target := 1
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err := arr.IndexOf(target)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Errorf("function IndexOf has bug")
	}
}

func TestArrayList_Insert_With_Index_In_The_Array(t *testing.T) {
	target := "A"
	index := 0
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		t.Errorf("insert value has error! error: %v", err)
		return
	}
	if arr.data[index] == target {
		t.Logf("insert value %s to ArrayList index %d is success", target, index)
	} else {
		t.Errorf("insert value %s to ArrayList index %d is failed", target, index)
	}
}

func TestArrayList_Insert_With_Index_Not_In_The_ArrayList_And_ArrayList_Not_Full(t *testing.T) {
	target := "d"
	index := 5
	in := []string{"a", "b", "c"}
	capacity := 4
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		if errors.Is(err, errors2.IndexOutOfBoundsError) {
			t.Logf("ArrayList not full and index out of bounds")
		} else {
			t.Errorf("unknown error! error: %v", err)
		}
	} else {
		t.Error("function Insert has bug")
	}
}

func TestArrayList_Insert_With_Index_Not_In_The_ArrayList_And_ArrayList_Is_Full(t *testing.T) {
	target := "d"
	index := -1
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		if errors.Is(err, errors2.IndexOutOfBoundsError) {
			t.Logf("ArrayList not full and index out of bounds")
		} else {
			t.Errorf("unknown error! error: %v", err)
		}
	} else {
		t.Error("function Insert has bug")
	}
}

func TestArrayList_Insert_With_Different_Type_Value(t *testing.T) {
	target := 1
	index := -1
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Error("function Insert has bug")
	}
}

func TestArrayList_IsEmpty_With_ArrayList_Is_Empty(t *testing.T) {
	capacity := 3
	typeOf := "string"
	arr := initArrayList(typeOf, capacity, t)
	if arr.IsEmpty() {
		t.Log("ArrayList is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArrayList_IsEmpty_With_ArrayList_Is_Not_Empty(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	if !arr.IsEmpty() {
		t.Log("ArrayList is not empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArrayList_Len(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	target := len(in)
	if target == arr.Len() {
		t.Logf("ArrayList len is %d \n", target)
	} else {
		t.Error("function Len has bug")
	}
}

func TestArrayList_MergeArrayList_With_Not_Empty_Array(t *testing.T) {
	target := []string{"a", "b", "c", "A", "B", "C"}
	flag := true
	arrA := []string{"a", "b", "c"}
	arrB := []string{"A", "B", "C"}
	capacity := 3
	typeOf := "string"
	a := initArrayList(typeOf, capacity, t)
	for _, v := range arrA {
		_ = a.Add(v)
	}
	b := initArrayList(typeOf, capacity, t)
	for _, v := range arrB {
		_ = b.Add(v)
	}
	array := a.MergeArrayList(b)
	if array.Len() != len(target) {
		t.Error("merge ArrayList failed")
		return
	}
	for i := 0; i < array.len; i++ {
		if array.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("merge ArrayList success")
	} else {
		t.Error("merge ArrayList failed")
	}
}

func TestArrayList_MergeArrayList_With_Empty_Array(t *testing.T) {
	target := []string{"a", "b", "c"}
	flag := true
	arrA := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	a := initArrayList(typeOf, capacity, t)
	for _, v := range arrA {
		_ = a.Add(v)
	}
	b := initArrayList(typeOf, capacity, t)
	array := a.MergeArrayList(b)
	if array.Len() != len(target) {
		t.Error("merge ArrayList failed")
		return
	}
	for i := 0; i < array.len; i++ {
		if array.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("merge ArrayList success")
	} else {
		t.Error("merge ArrayList failed")
	}
}

func TestArrayList_Remove_With_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "c"}
	index := 1
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	ret, err1 := arr.Remove(index)
	if err1 != nil {
		t.Errorf("remove is error! error: %v \n", err1)
		return
	}
	for i := 0; i < arr.Len(); i++ {
		if target[i] != arr.data[i] {
			flag = false
			break
		}
	}
	if arr.Len() == capacity-1 && ret == in[index] && flag {
		t.Logf("success to remove the element with index %d in the ArrayList \n", index)
	} else {
		t.Error("function Remove has bug")
	}
}

func TestArrayList_Remove_With_Index_Not_In_The_Array(t *testing.T) {
	index := -1
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.Remove(index)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err1)
			return
		}
	} else {
		t.Error("function Remove has bug")
	}
}

func TestArrayList_Replace_With_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "B", "c"}
	index := 1
	newVal := "B"
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	oldVal, err1 := arr.Replace(index, newVal)
	if err1 != nil {
		t.Errorf("replace value has error! error: %v \n", err1)
		return
	}
	if arr.Len() != len(target) {
		t.Error("function Replace has bug")
		return
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.data[i] != target[i] {
			flag = false
			break
		}
	}
	if oldVal == in[index] && flag {
		t.Logf("success to replace the element with index %d in the ArrayList \n", index)
	} else {
		t.Errorf("failed to replace the element with index %d in the ArrayList \n", index)
	}
}

func TestArrayList_Replace_With_Index_Not_In_The_Array(t *testing.T) {
	index := -1
	newVal := "B"
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.Replace(index, newVal)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err1)
			return
		}
	} else {
		t.Error("function Replace has bug")
	}
}

func TestArrayList_Replace_With_Different_Type_Value(t *testing.T) {
	index := 1
	newVal := 100
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err := arr.Replace(index, newVal)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Error("function Replace has bug")
	}
}

func TestArrayList_Set_With_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "B", "c"}
	index := 1
	newVal := "B"
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	err := arr.Set(index, newVal)
	if err != nil {
		t.Errorf("set value has error! error: %v \n", err)
		return
	}
	if arr.Len() != len(target) {
		t.Error("function Set has bug")
		return
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("success to set the element with index %d in the ArrayList \n", index)
	} else {
		t.Errorf("failed to set the element with index %d in the ArrayList \n", index)
	}
}

func TestArrayList_Set_With_Index_Not_In_The_Array(t *testing.T) {
	index := -1
	newVal := "B"
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	err := arr.Set(index, newVal)
	if err != nil {
		if errors.Is(err, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Set has bug")
	}
}

func TestArrayList_Set_With_Different_Type_Value(t *testing.T) {
	index := 2
	newVal := 100
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	err := arr.Set(index, newVal)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
		}
	} else {
		t.Error("function Set has bug")
	}
}

func TestArrayList_SubArrayList_With_Start_And_End_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "b"}
	start, end := 0, 2
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	subArr, err1 := arr.SubArrayList(start, end)
	if err1 != nil {
		t.Errorf("get sub ArrayList has error! error: %v \n", err1)
		return
	}
	if subArr.Len() != len(target) {
		t.Error("function Set has bug")
		return
	}
	for i := 0; i < subArr.Len(); i++ {
		if subArr.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("sub ArrayList of %d to %d was obtained successfully \n", start, end)
	} else {
		t.Errorf("sub ArrayList of %d to %d was obtained successfully \n", start, end)
	}
}

func TestArrayList_SubArrayList_With_End_Less_Than_Start(t *testing.T) {
	start, end := 2, 1
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.SubArrayList(start, end)
	if err1 != nil {
		if err1.Error() == "start must be less than or equal to end" {
			t.Logf("start must be less than or equal to end")
			return
		} else {
			t.Error("function SubArrayList has bug")
			return
		}
	} else {
		t.Error("function SubArrayList has bug")
	}
}

func TestArrayList_SubArrayList_With_Start_Less_Than_Zero(t *testing.T) {
	start, end := -1, 2
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.SubArrayList(start, end)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Error("function SubArrayList has bug")
			return
		}
	} else {
		t.Error("function SubArrayList has bug")
	}
}

func TestArrayList_SubArrayList_With_End_Greater_Than_ArrayList_Length(t *testing.T) {
	start, end := -1, 4
	in := []string{"a", "b", "c"}
	capacity := 3
	typeOf := "string"
	arr := initArrayListAndAddValue(typeOf, capacity, t, in)
	_, err1 := arr.SubArrayList(start, end)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Error("function SubArrayList has bug")
			return
		}
	} else {
		t.Error("function SubArrayList has bug")
	}
}

// initArrayList ,init a empty Array
func initArrayList(typeOf string, capacity int, t *testing.T) *ArrayList {
	arr, err := NewArrayListWithCap(typeOf, capacity)
	if err != nil {
		t.Errorf("init ArrayList has error! error: %v ", err)
		return nil
	}
	return arr
}

// initArrayListAndAddValue ,init A empty ArrayList and add value
func initArrayListAndAddValue(typeOf string, capacity int, t *testing.T, in []string) *ArrayList {
	arr := initArrayList(typeOf, capacity, t)
	for _, v := range in {
		_ = arr.Add(v)
	}
	return arr
}
