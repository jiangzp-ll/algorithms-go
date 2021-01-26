package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length int
}

// NewArray 数组初始化
func NewArray(cap int) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

// Len 数组长度
func (arr *Array) Len() int {
	return arr.length
}

// Data 元素
func (arr *Array) Data() []int {
	return arr.data
}

// IsIndexOutOfRange 判断数组是否越界
func (arr *Array) IsIndexOutOfRange(index int) bool {
	if index >= cap(arr.data) {
		return true
	}
	return false
}

// Find 通过索引查找数组，索引范围[0,n-1]
func (arr *Array) Find(index int) (int, error) {
	if arr.IsIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return arr.data[index], nil
}

// Insert 插入元素
func (arr *Array) Insert(index, val int) error {
	if len(arr.data) == arr.Len() {
		arr.Expansion()
	}
	if index != arr.length && arr.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := arr.length; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = val
	arr.length++
	return nil
}

// InsertToTail 从尾部插入元素
func (arr *Array) InsertToTail(val int) error {
	return arr.Insert(arr.Len(), val)
}

// Delete 从数组中删除下标是index的元素并返回
func (arr *Array) Delete(index int) (int, error) {
	if arr.IsIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	val := arr.data[index]
	for i := index; i < arr.Len(); i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.length--
	return val, nil
}

// Print 打印数组
func (arr *Array) Print() {
	var format string
	for i := 0; i < arr.Len(); i++ {
		format += fmt.Sprintf("|%+v", arr.data[i])
	}
	fmt.Println(format)
}

// Expansion 扩容策略
func (arr *Array) Expansion() {
	if len(arr.data) < 1024 {
		newData := make([]int, 2*len(arr.data))
		for i := 0; i < arr.Len(); i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData
	} else {
		nl := float64(arr.Len()) * 1.25
		newData := make([]int, int(nl))
		for i := 0; i < arr.Len(); i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData
	}
}

// MergeArray 合并两个数组成一个有序数组
func (arr *Array) MergeArray(other *Array) *Array {
	al := arr.Len()
	ol := other.Len()
	result := NewArray(al + ol)
	var l, r, k int
	for l < al && r < ol {
		if arr.data[l] < other.data[r] {
			_ = result.Insert(k, arr.data[l])
			l++
		} else {
			_ = result.Insert(k, other.data[r])
			r++
		}
		k++
	}
	if l < al {
		for i := l; i < al; i++ {
			_ = result.Insert(k, arr.data[i])
			k++
		}
	}
	if r < ol {
		for i := r; i < ol; i++ {
			_ = result.Insert(k, other.data[i])
			k++
		}
	}
	return result
}
