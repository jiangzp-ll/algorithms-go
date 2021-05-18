package array

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"math/rand"
	"strconv"
	"testing"
)

func TestArray_NewArray_With_Capacity_Greater_Than_Zero(t *testing.T) {
	capacity := 2
	arr := initArray(capacity, t)
	if arr.Len() == 0 && len(arr.Data()) == capacity {
		t.Log("init Array success")
	} else {
		t.Error("init Array failed")
	}
}

func TestArray_NewArray_With_Capacity_Equal_Zero(t *testing.T) {
	capacity := 0
	arr := initArray(capacity, t)
	if arr.Len() == 0 && len(arr.Data()) == capacity {
		t.Log("init Array success")
	} else {
		t.Error("init Array failed")
	}
}

func TestArray_NewArray_With_Capacity_Less_Than_Zero(t *testing.T) {
	capacity := -1
	_, err := NewArray(capacity)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("init Array has bug")
	}
}

func TestArray_Add_With_Not_Need_Expansion(t *testing.T) {
	in := "a"
	capacity := 2
	arr := initArray(capacity, t)
	arr.Add(in)
	if len(arr.data) == capacity && arr.len == 1 && arr.data[0] == in {
		t.Logf("add %s to Array is success", in)
	} else {
		t.Errorf("add %s to Array is failed", in)
	}
}

func TestArray_Add_With_Need_Expansion_And_Array_Length_Less_Than_1024(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 2
	flag := true
	arr := initArrayAndAddValue(capacity, t, in)
	for i := 0; i < len(in); i++ {
		if in[i] != arr.data[i] {
			flag = false
			break
		}
	}
	if len(arr.data) == capacity*2 && arr.len == 3 && flag {
		t.Logf("add %v to Array is success", in)
	} else {
		t.Errorf("add %v to Array is failed", in)
	}
}

func TestArray_Add_With_Need_Expansion_And_Array_Length_Greater_Than_1024(t *testing.T) {
	capacity := 1024
	arr := initArray(capacity, t)
	for i := 0; i <= capacity; i++ {
		arr.Add(strconv.Itoa(i))
	}
	if len(arr.data) == int(float64(capacity)*1.25) && arr.len == 1025 {
		t.Log("function Add is good")
	} else {
		t.Error("function Add has bug")
	}
}

func TestArray_Add_With_Need_Expansion_And_Array_Is_Empty(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 0
	flag := true
	arr := initArrayAndAddValue(capacity, t, in)
	for i := 0; i < len(in); i++ {
		if in[i] != arr.data[i] {
			flag = false
			break
		}
	}
	t.Logf("array: %#v \n", arr)
	if len(arr.data) == 4 && arr.len == 3 && flag {
		t.Logf("add %v to Array is success", in)
	} else {
		t.Errorf("add %v to Array is failed", in)
	}
}

func TestArray_Clear(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	flag := true
	arr := initArrayAndAddValue(capacity, t, in)
	arr.Clear()
	for i := 0; i < len(arr.data); i++ {
		if "" != arr.data[i] {
			flag = false
			break
		}
	}
	if arr.len == capacity && flag {
		t.Log("clear Array is success")
	} else {
		t.Error("clear Array is failed")
	}
}

func TestArray_Contains_With_Contain_The_Value(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	target := "a"
	arr := initArrayAndAddValue(capacity, t, in)
	ret := arr.Contain(target)
	if ret {
		t.Logf("Array contain %s \n", target)
	} else {
		t.Errorf("Array not contain %s \n", target)
	}
}

func TestArray_Contains_With_Not_Contain_The_Value(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	target := "A"
	arr := initArrayAndAddValue(capacity, t, in)
	ret := arr.Contain(target)
	if !ret {
		t.Logf("Array not contain %s \n", target)
	} else {
		t.Error("function Contain has bug")
	}
}

func TestArray_Data(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	flag := true
	arr := initArrayAndAddValue(capacity, t, in)
	data := arr.Data()
	for i := 0; i < len(data); i++ {
		if data[i] != in[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("Array is %v \n", data)
	} else {
		t.Error("func Data has bug")
	}
}

func TestArray_Get_With_Index_In_The_Array(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	index := rand.Intn(3)
	arr := initArrayAndAddValue(capacity, t, in)
	val, err1 := arr.Get(index)
	if err1 != nil {
		t.Errorf("get value is failed! error: %v", err1)
		return
	}
	if val == in[index] {
		t.Logf("get the index is %d elements in the Array success", index)
	} else {
		t.Errorf("get the index is %d elements in the Array failed", index)
	}
}

func TestArray_Get_With_Index_Not_In_The_Array(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	index := capacity + 1
	arr := initArrayAndAddValue(capacity, t, in)
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

func TestArray_IndexOf_With_Array_Has_Existed_The_Value(t *testing.T) {
	target := "c"
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	index, err1 := arr.IndexOf(target)
	if err1 != nil {
		t.Errorf("get index failed! error: %v \n", err1)
		return
	}
	if index == arr.len-1 {
		t.Logf("the element with index %d in the array is %s \n", index, target)
	} else {
		t.Error("function IndexOf has bug")
	}
}

func TestArray_IndexOf_With_Array_Not_Exist_The_Value(t *testing.T) {
	target := "C"
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	_, err1 := arr.IndexOf(target)
	if err1 != nil {
		t.Logf("Array not exist %s \n", target)
	} else {
		t.Error("function IndexOf has bug")
	}
}

func TestArray_Insert_With_Index_In_The_Array(t *testing.T) {
	target := "A"
	index := 0
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		t.Errorf("insert value has error! error: %v", err)
		return
	}
	if arr.data[index] == target {
		t.Logf("insert value %s to Array index %d is success", target, index)
	} else {
		t.Errorf("insert value %s to Array index %d is failed", target, index)
	}
}

func TestArray_Insert_With_Index_Not_In_The_Array_And_Array_Not_Full(t *testing.T) {
	target := "d"
	index := 5
	in := []string{"a", "b", "c"}
	capacity := 4
	arr := initArrayAndAddValue(capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		if errors.Is(err, errors2.IndexOutOfBoundsError) {
			t.Logf("Array not full and index out of bounds")
		} else {
			t.Errorf("unknown error! error: %v", err)
		}
	} else {
		t.Error("function Insert has bug")
	}
}

func TestArray_Insert_With_Index_Not_In_The_Array_And_Array_Is_Full(t *testing.T) {
	target := "d"
	index := -1
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	if err := arr.Insert(index, target); err != nil {
		if errors.Is(err, errors2.IndexOutOfBoundsError) {
			t.Logf("Array not full and index out of bounds")
		} else {
			t.Errorf("unknown error! error: %v", err)
		}
	} else {
		t.Error("function Insert has bug")
	}
}

func TestArray_IsEmpty_With_Array_Is_Empty(t *testing.T) {
	capacity := 3
	arr := initArray(capacity, t)
	if arr.IsEmpty() {
		t.Log("Array is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArray_IsEmpty_With_Array_Is_Not_Empty(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	if !arr.IsEmpty() {
		t.Log("Array is not empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArray_Len(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	target := len(in)
	if target == arr.Len() {
		t.Logf("Array len is %d \n", target)
	} else {
		t.Error("function Len has bug")
	}
}

func TestArray_MergeArray_With_Not_Empty_Array(t *testing.T) {
	target := []string{"a", "b", "c", "A", "B", "C"}
	flag := true
	arrA := []string{"a", "b", "c"}
	arrB := []string{"A", "B", "C"}
	capacity := 3
	a := initArray(capacity, t)
	for _, v := range arrA {
		a.Add(v)
	}
	b := initArray(capacity, t)
	for _, v := range arrB {
		b.Add(v)
	}
	array := a.MergeArray(b)
	if array.Len() != len(target) {
		t.Error("merge Array failed")
		return
	}
	for i := 0; i < array.len; i++ {
		if array.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("merge Array success")
	} else {
		t.Error("merge Array failed")
	}
}

func TestArray_MergeArray_With_Empty_Array(t *testing.T) {
	target := []string{"a", "b", "c"}
	flag := true
	arrA := []string{"a", "b", "c"}
	capacity := 3
	a := initArray(capacity, t)
	for _, v := range arrA {
		a.Add(v)
	}
	b := initArray(capacity, t)
	array := a.MergeArray(b)
	if array.Len() != len(target) {
		t.Error("merge Array failed")
		return
	}
	for i := 0; i < array.len; i++ {
		if array.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("merge Array success")
	} else {
		t.Error("merge Array failed")
	}
}

func TestArray_Remove_With_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "c"}
	index := 1
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	ret, err1 := arr.Remove(index)
	if err1 != nil {
		t.Errorf("remove is error! error: %v \n", err1)
		return
	}
	for i := 0; i < arr.len; i++ {
		if target[i] != arr.data[i] {
			flag = false
			break
		}
	}
	if arr.len == capacity-1 && ret == in[index] && flag {
		t.Logf("success to remove the element with index %d in the Array \n", index)
	} else {
		t.Error("function Remove has bug")
	}
}

func TestArray_Remove_With_Index_Not_In_The_Array(t *testing.T) {
	index := -1
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
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

func TestArray_Replace_With_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "B", "c"}
	index := 1
	newVal := "B"
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	oldVal, err1 := arr.Replace(index, newVal)
	if err1 != nil {
		t.Errorf("replace value has error! error: %v \n", err1)
		return
	}
	if arr.len != len(target) {
		t.Error("function Replace has bug")
		return
	}
	for i := 0; i < arr.len; i++ {
		if arr.data[i] != target[i] {
			flag = false
			break
		}
	}
	if oldVal == in[index] && flag {
		t.Logf("success to replace the element with index %d in the Array \n", index)
	} else {
		t.Errorf("failed to replace the element with index %d in the Array \n", index)
	}
}

func TestArray_Replace_With_Index_Not_In_The_Array(t *testing.T) {
	index := -1
	newVal := "B"
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
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

func TestArray_Set_With_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "B", "c"}
	index := 1
	newVal := "B"
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	err := arr.Set(index, newVal)
	if err != nil {
		t.Errorf("set value has error! error: %v \n", err)
		return
	}
	if arr.len != len(target) {
		t.Error("function Set has bug")
		return
	}
	for i := 0; i < arr.len; i++ {
		if arr.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("success to set the element with index %d in the Array \n", index)
	} else {
		t.Errorf("failed to set the element with index %d in the Array \n", index)
	}
}

func TestArray_Set_With_Index_Not_In_The_Array(t *testing.T) {
	index := -1
	newVal := "B"
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
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

func TestArray_SubArray_With_Start_And_End_Index_In_The_Array(t *testing.T) {
	target := []string{"a", "b"}
	start, end := 0, 2
	flag := true
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	subArr, err1 := arr.SubArray(start, end)
	if err1 != nil {
		t.Errorf("get sub Array has error! error: %v \n", err1)
		return
	}
	if subArr.len != len(target) {
		t.Error("function Set has bug")
		return
	}
	for i := 0; i < subArr.len; i++ {
		if subArr.data[i] != target[i] {
			flag = false
			break
		}
	}
	if flag {
		t.Logf("sub Array of %d to %d was obtained successfully \n", start, end)
	} else {
		t.Errorf("sub Array of %d to %d was obtained successfully \n", start, end)
	}
}

func TestArray_SubArray_With_End_Less_Than_Start(t *testing.T) {
	start, end := 2, 1
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	_, err1 := arr.SubArray(start, end)
	if err1 != nil {
		if err1.Error() == "start must be less than or equal to end" {
			t.Logf("start must be less than or equal to end")
			return
		} else {
			t.Error("function SubArray has bug")
			return
		}
	} else {
		t.Error("function SubArray has bug")
	}
}

func TestArray_SubArray_With_Start_Less_Than_Zero(t *testing.T) {
	start, end := -1, 2
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	_, err1 := arr.SubArray(start, end)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Error("function SubArray has bug")
			return
		}
	} else {
		t.Error("function SubArray has bug")
	}
}

func TestArray_SubArray_With_End_Greater_Than_Array_Length(t *testing.T) {
	start, end := -1, 4
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	_, err1 := arr.SubArray(start, end)
	if err1 != nil {
		if errors.Is(err1, errors2.IndexOutOfBoundsError) {
			t.Logf("index out of bounds")
			return
		} else {
			t.Error("function SubArray has bug")
			return
		}
	} else {
		t.Error("function SubArray has bug")
	}
}

func TestArray_ToString(t *testing.T) {
	target := "[a, b, c]"
	in := []string{"a", "b", "c"}
	capacity := 3
	arr := initArrayAndAddValue(capacity, t, in)
	ret := arr.ToString()
	if ret == target {
		t.Log("function ToString is good")
	} else {
		t.Error("function ToString has bug")
	}
}

// initArray ,init a empty Array
func initArray(capacity int, t *testing.T) *Array {
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return nil
	}
	return arr
}

// initArrayAndAddValue ,init A empty Array and add value
func initArrayAndAddValue(capacity int, t *testing.T, in []string) *Array {
	arr := initArray(capacity, t)
	for _, v := range in {
		arr.Add(v)
	}
	return arr
}
