package array

import (
	"fmt"
	"testing"
)

var arr = NewArray(10)

func TestNewArray(t *testing.T) {
	cap := 10
	arr := NewArray(cap)
	if arr.Len() == 0 && len(arr.Data()) == 10 {
		t.Logf("成功初始化长度%d的数组！\n", cap)
	} else {
		t.Error("初始化失败！")
	}
}

func TestIsIndexOutOfRange(t *testing.T) {
	if arr.IsIndexOutOfRange(5) {
		t.Error("数组角标越界了")
	} else {
		t.Log("数组角标在范围之内")
	}
}

func TestArray_Insert(t *testing.T) {
	err := arr.Insert(0, 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("下标 %d 处，插入了 %d \n", 0, 1)
}

func TestArray_Find(t *testing.T) {
	err := arr.Insert(0, 1)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := arr.Find(0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("下标 %d 的元素是 %d \n", 0, result)
}

func TestArray_InsertToTail(t *testing.T) {
	for i := 0; i < len(arr.Data())-1; i++ {
		err := arr.Insert(i, i+1)
		if err != nil {
			t.Error(err)
			return
		}
	}
	err := arr.InsertToTail(100)
	if err != nil {
		t.Error(err)
		return
	}
	val, err := arr.Find(arr.Len() - 1)
	if err != nil {
		t.Error(err)
		return
	}
	if val == 100 {
		t.Logf("数组最后一个元素是 %d \n", val)
	} else {
		t.Error("在数组末尾添加元素失败！")
	}
}

func TestArray_Delete(t *testing.T) {
	err := arr.Insert(0, 1)
	if err != nil {
		t.Error(err)
		return
	}
	reslut, err := arr.Delete(0)
	if err != nil {
		t.Error("delete element failed!")
		return
	}
	t.Logf("从数组中删除了 %d \n", reslut)
}

func TestArray_Print(t *testing.T) {
	for i := 0; i <= 10; i++ {
		_ = arr.InsertToTail(i + 1)
	}
	fmt.Println("A")
	arr.Print()
}

// 数组动态扩容
func TestArray_Expansion(t *testing.T) {
	array := NewArray(2)
	fmt.Printf("before array address: %v \n", &array)
	for i := 0; i < 1025; i++ {
		_ = array.Insert(i, i+1)
	}
	fmt.Printf("after array address: %v \n", &array)
	fmt.Printf("cap: %d, len: %d \n", cap(array.Data()), len(array.Data()))
	fmt.Println(array.Len())
	array.Print()
}

// 将两个有序书组合合并成一个有序数组
func TestMergeArray(t *testing.T) {
	arr := NewArray(2)
	other := NewArray(2)
	_ = arr.Insert(0, 1)
	_ = arr.Insert(1, 3)
	_ = other.Insert(0, 2)
	_ = other.Insert(1, 4)
	result := arr.MergeArray(other)
	fmt.Println(*result)
}