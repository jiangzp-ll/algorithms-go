package array

import (
	"errors"
)

var IndexOutOfBoundsError = errors.New("array index out of bounds error")

// Array ,like Java ArrayList
type Array struct {
	data []string
	len  int
}

// NewArray ,init Array
func NewArray(capacity int) (*Array, error) {
	if capacity < 0 {
		return nil, errors.New("capacity must be greater than or equal to zero")
	}
	return &Array{
		data: make([]string, capacity),
		len:  0,
	}, nil
}

// Add ,
func (arr *Array) Add(val string) {
	if len(arr.data) == arr.Len() {
		arr.expansion()
	}
	arr.data[arr.len] = val
	arr.len++
}

// Clear ，make Array to empty
func (arr *Array) Clear() {
	for i := 0; i < len(arr.data); i++ {
		arr.data[i] = ""
	}
}

// Contain, determine whether the value is included in the Array
func (arr *Array) Contain(val string) bool {
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

// expansion , Array expansion
// When the number of elements in the array is less than 1024, the capacity becomes twice of the original array;
// When the number of elements in the array is greater than 1024, the capacity becomes 1.25 times of the original array.
func (arr *Array) expansion() {
	if arr.len < 1024 {
		if arr.len == 0 {
			arr.data = make([]string, 1)
			return
		}
		newData := make([]string, 2*arr.len)
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
func (arr *Array) Get(index int) (string, error) {
	if arr.isIndexOutOfRange(index) {
		return "", IndexOutOfBoundsError
	}
	return arr.data[index], nil
}

// IndexOf ,return the index of the value in the Array
func (arr *Array) IndexOf(val string) (int, error) {
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] == val {
			return i, nil
		}
	}
	return -1, errors.New("not exist the value in the Array")
}

// Insert, insert the value at the specified index
func (arr *Array) Insert(index int, val string) error {
	if len(arr.data) == arr.Len() {
		arr.expansion()
	}
	if index != arr.len && arr.isIndexOutOfRange(index) {
		return IndexOutOfBoundsError
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

// isIndexOutOfRange ,determine whether the index of the Array is out of bounds
func (arr *Array) isIndexOutOfRange(index int) bool {
	return index >= len(arr.data) || index < 0
}

// Len , get Array element number
func (arr *Array) Len() int {
	return arr.len
}

// MergeArray , merge two Array
func (arr *Array) MergeArray(other *Array) *Array {
	//al, ol := arr.len, other.len
	//result, _ := NewArray(al + ol)
	//var l, r, k int
	//for l < al && r < ol {
	//	if arr.data[l] < other.data[r] {
	//		_ = result.Insert(k, arr.data[l])
	//		l++
	//	} else {
	//		_ = result.Insert(k, other.data[r])
	//		r++
	//	}
	//	k++
	//}
	//if l < al {
	//	for i := l; i < al; i++ {
	//		_ = result.Insert(k, arr.data[i])
	//		k++
	//	}
	//}
	//if r < ol {
	//	for i := r; i < ol; i++ {
	//		_ = result.Insert(k, other.data[i])
	//		k++
	//	}
	//}
	//return result
	if other.len < 1 {
		return arr
	}
	for i := 0; i < other.len; i++ {
		arr.Add(other.data[i])
	}
	return arr
}

// Remove , remove and return the element with the specified index
func (arr *Array) Remove(index int) (string, error) {
	if arr.isIndexOutOfRange(index) {
		return "", IndexOutOfBoundsError
	}
	val := arr.data[index]
	for i := index; i < len(arr.data)-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.len--
	return val, nil
}

// Replace ,replace and return the new value with the specified index
func (arr *Array) Replace(index int, val string) (string, error) {
	if arr.isIndexOutOfRange(index) {
		return "", IndexOutOfBoundsError
	}
	old := arr.data[index]
	arr.data[index] = val
	return old, nil
}

// Set , set a value for the specified index
func (arr *Array) Set(index int, val string) error {
	if arr.isIndexOutOfRange(index) {
		return IndexOutOfBoundsError
	}
	arr.data[index] = val
	return nil
}

// SubArray ,get sub Array
func (arr *Array) SubArray(start, end int) (*Array, error) {
	if start > end {
		return nil, errors.New("start must be less than or equal to end")
	}
	if arr.isIndexOutOfRange(start) || arr.isIndexOutOfRange(end) {
		return nil, IndexOutOfBoundsError
	}
	return &Array{
		data: arr.data[start:end],
		len:  end - start,
	}, nil
}

// ToString ,convert array to string separated by ","
func (arr *Array) ToString() string {
	ret := "["
	for i := 0; i < len(arr.data); i++ {
		if i == len(arr.data)-1 {
			ret = ret + arr.data[i]
		} else {
			ret = ret + arr.data[i] + ", "
		}
	}
	return ret + "]"
}
