package array

import (
	"errors"
	"fmt"
)

// Array ,like Java ArrayList
type Array struct {
	data []string
	len  int
}

// NewArray ,init Array
func NewArray(capacity int) (*Array, error) {
	if capacity < 0 {
		return nil, errors.New("capacity must be greater than zero")
	}
	return &Array{
		data: make([]string, capacity),
		len:  0,
	}, nil
}

// Clear ，make Array to empty
func (arr *Array) Clear() {
	for i := 0; i < len(arr.data); i++ {
		arr.data[i] = ""
	}
}

// Contains, determine whether the value is included in the Array
func (arr *Array) Contains(val string) bool {
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] == val {
			return true
		}
	}
	return false
}

// Data ,get Array all element
func (arr *Array) Data() []string {
	return arr.data
}

// Expansion , Array expansion
// When the number of elements in the array is less than 1024, the capacity becomes twice of the original array;
// When the number of elements in the array is greater than 1024, the capacity becomes 1.25 times of the original array.
func (arr *Array) Expansion() {
	if len(arr.data) < 1024 {
		newData := make([]string, 2*len(arr.data))
		for i := 0; i < arr.Len(); i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData
	} else {
		nl := float64(arr.Len()) * 1.25
		newData := make([]string, int(nl))
		for i := 0; i < arr.Len(); i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData
	}
}

// Get ,get the element at the specified index
// TODO 全局错误码
func (arr *Array) Get(index int) (string, error) {
	if arr.IsIndexOutOfRange(index) {
		return "", errors.New("out of index range")
	}
	return arr.data[index], nil
}

// IndexOf ,return the index of the value in the Array
// TODO 全局错误码
func (arr *Array) IndexOf(val string) int {
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] == val {
			return i
		}
	}
	return -1
}

// Insert, insert the value at the specified index
func (arr *Array) Insert(index int, val string) error {
	if len(arr.data) == arr.Len() {
		arr.Expansion()
	}
	if index != arr.len && arr.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := arr.len; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = val
	arr.len++
	return nil
}

// IsEmpty ,determine whether the Array is an empty Array
func (arr *Array) IsEmpty() bool {
	return arr.len == 0
}

// IsIndexOutOfRange ,determine whether the index of the Array is out of bounds
func (arr *Array) IsIndexOutOfRange(index int) bool {
	return index >= len(arr.data)
}

// Len , get Array element number
func (arr *Array) Len() int {
	return arr.len
}

// MergeArray , merge two Array
func (arr *Array) MergeArray(other *Array) *Array {
	al, ol := arr.len, other.len
	result, _ := NewArray(al + ol)
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

// Remove , remove and return the element with the specified index
func (arr *Array) Remove(index int) (string, error) {
	if arr.IsIndexOutOfRange(index) {
		return "", errors.New("out of index range")
	}
	val := arr.data[index]
	for i := index; i < len(arr.data); i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.len--
	return val, nil
}

// Replace ,replace and return the new value with the specified index
func (arr *Array) Replace(index int, val string) (string, error) {
	if arr.IsIndexOutOfRange(index) {
		return "", errors.New("out of index range")
	}
	old := arr.data[index]
	arr.data[index] = val
	return old, nil
}

// Set , set a value for the specified index
func (arr *Array) Set(index int, val string) error {
	if arr.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	arr.data[index] = val
	return nil
}

// Find 通过索引查找数组，索引范围[0,n-1]
func (arr *Array) Find(index int) (string, error) {
	if arr.IsIndexOutOfRange(index) {
		return "", errors.New("out of index range")
	}
	return arr.data[index], nil
}

// InsertToTail 从尾部插入元素
func (arr *Array) InsertToTail(val string) error {
	return arr.Insert(arr.Len(), val)
}

// Print 打印数组
func (arr *Array) Print() {
	var format string
	for i := 0; i < arr.Len(); i++ {
		format += fmt.Sprintf("|%+v", arr.data[i])
	}
	fmt.Println(format)
}
