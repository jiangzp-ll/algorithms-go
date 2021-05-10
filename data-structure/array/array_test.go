package array

import (
	"errors"
	"math/rand"
	"testing"
)

//var arr = NewArray(10)

func TestArray_NewArray_With_Capacity_Greater_Than_Zero(t *testing.T) {
	capacity := 2
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	if arr.Len() == 0 && len(arr.Data()) == capacity {
		t.Log("init Array success")
	} else {
		t.Error("init Array failed")
	}
}

func TestArray_NewArray_With_Capacity_Equal_Zero(t *testing.T) {
	capacity := 0
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	arr.Add(in)
	if len(arr.data) == capacity && arr.len == 1 && arr.data[0] == in {
		t.Logf("add %s to Array is success", in)
	} else {
		t.Errorf("add %s to Array is failed", in)
	}
}

func TestArray_Add_With_Need_Expansion(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 2
	flag := true
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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

func TestArray_Clear(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	flag := true
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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

func TestArray_Get_Index_In_The_Array(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	index := rand.Intn(3)
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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

func TestArray_Get_Index_Not_In_The_Array(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	index := capacity + 1
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	_, err1 := arr.Get(index)
	if err1 != nil {
		if errors.Is(err1, IndexOutOfBoundsError) {
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	_, err1 := arr.IndexOf(target)
	if err1 != nil {
		t.Logf("Array not exist %s \n", target)
	} else {
		t.Error("function IndexOf has bug")
	}
}

func TestArray_Insert_Index_In_The_Array(t *testing.T) {
	target := "A"
	index := 0
	in := []string{"a", "b", "c"}
	capacity := 3
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	if err = arr.Insert(index, target); err != nil {
		t.Errorf("insert value has error! error: %v", err)
		return
	}
	if arr.data[index] == target {
		t.Logf("insert value %s to Array index %d is success", target, index)
	} else {
		t.Errorf("insert value %s to Array index %d is failed", target, index)
	}
}

func TestArray_Insert_Index_Not_In_The_Array_And_Array_Not_Full(t *testing.T) {
	target := "d"
	index := 5
	in := []string{"a", "b", "c"}
	capacity := 4
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	if err = arr.Insert(index, target); err != nil {
		if errors.Is(err, IndexOutOfBoundsError) {
			t.Logf("Array not full and index out of bounds")
		} else {
			t.Errorf("unknown error! error: %v", err)
		}
	} else {
		t.Error("function Insert has bug")
	}
}

func TestArray_Insert_Index_Not_In_The_Array_And_Array_Is_Full(t *testing.T) {
	target := "d"
	index := -1
	in := []string{"a", "b", "c"}
	capacity := 3
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	if err = arr.Insert(index, target); err != nil {
		if errors.Is(err, IndexOutOfBoundsError) {
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
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	if arr.IsEmpty() {
		t.Log("Array is empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArray_IsEmpty_With_Array_Is_Not_Empty(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	if !arr.IsEmpty() {
		t.Log("Array is not empty")
	} else {
		t.Error("function IsEmpty has bug")
	}
}

func TestArray_Len(t *testing.T) {
	in := []string{"a", "b", "c"}
	capacity := 3
	arr, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range in {
		arr.Add(v)
	}
	target := len(in)
	if target == arr.Len() {
		t.Logf("Array len is %d \n", target)
	} else {
		t.Error("function Len has bug")
	}
}

func TestArray_MergeArray(t *testing.T) {
	target := []string{"a", "b", "c", "A", "B", "C"}
	flag := true
	arrA := []string{"a", "b", "c"}
	arrB := []string{"A", "B", "C"}
	capacity := 3
	a, err := NewArray(capacity)
	if err != nil {
		t.Errorf("init Array has error! error: %v \n", err)
		return
	}
	for _, v := range arrA {
		a.Add(v)
	}
	b, err1 := NewArray(capacity)
	if err1 != nil {
		t.Errorf("init Array has error! error: %v \n", err1)
		return
	}
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
	t.Logf("Array: %#v \n", array)
	if flag {
		t.Logf("merge Array success")
	} else {
		t.Error("merge Array failed")
	}
}

//func TestIsIndexOutOfRange(t *testing.T) {
//	if arr.IsIndexOutOfRange(5) {
//		t.Error("数组角标越界了")
//	} else {
//		t.Log("数组角标在范围之内")
//	}
//}
//
//func TestArray_Insert(t *testing.T) {
//	err := arr.Insert(0, 1)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	t.Logf("下标 %d 处，插入了 %d \n", 0, 1)
//}
//
//func TestArray_Find(t *testing.T) {
//	err := arr.Insert(0, 1)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	result, err := arr.Find(0)
//	if err != nil {
//		t.Error(err)
//	}
//	t.Logf("下标 %d 的元素是 %d \n", 0, result)
//}
//
//func TestArray_InsertToTail(t *testing.T) {
//	for i := 0; i < len(arr.Data())-1; i++ {
//		err := arr.Insert(i, i+1)
//		if err != nil {
//			t.Error(err)
//			return
//		}
//	}
//	err := arr.InsertToTail(100)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	in, err := arr.Find(arr.Len() - 1)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	if in == 100 {
//		t.Logf("数组最后一个元素是 %d \n", in)
//	} else {
//		t.Error("在数组末尾添加元素失败！")
//	}
//}
//
//func TestArray_Delete(t *testing.T) {
//	err := arr.Insert(0, 1)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	reslut, err := arr.Delete(0)
//	if err != nil {
//		t.Error("delete element failed!")
//		return
//	}
//	t.Logf("从数组中删除了 %d \n", reslut)
//}
//
//func TestArray_Print(t *testing.T) {
//	for i := 0; i <= 10; i++ {
//		_ = arr.InsertToTail(i + 1)
//	}
//	fmt.Println("A")
//	arr.Print()
//}
//
//// 数组动态扩容
//func TestArray_Expansion(t *testing.T) {
//	array := NewArray(2)
//	fmt.Printf("before array address: %v \n", &array)
//	for i := 0; i < 1025; i++ {
//		_ = array.Insert(i, i+1)
//	}
//	fmt.Printf("after array address: %v \n", &array)
//	fmt.Printf("cap: %d, len: %d \n", cap(array.Data()), len(array.Data()))
//	fmt.Println(array.Len())
//	array.Print()
//}
//
//// 将两个有序书组合合并成一个有序数组
//func TestMergeArray(t *testing.T) {
//	arr := NewArray(2)
//	other := NewArray(2)
//	_ = arr.Insert(0, 1)
//	_ = arr.Insert(1, 3)
//	_ = other.Insert(0, 2)
//	_ = other.Insert(1, 4)
//	result := arr.MergeArray(other)
//	fmt.Println(*result)
//}
