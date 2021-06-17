package arraylist

import (
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"github.com/zepeng-jiang/go-basic-demo/pkg/generic"
)

const DefaultCapacity = 12

// ArrayListList ,like Java ArrayListList
type ArrayList struct {
	// data ,store input vale
	data []interface{}
	// len ,number of elements in ArrayList
	len int
	// typeOf ,ArrayList type
	// Because Go not have Generic
	typeOf string
}

// NewArrayList ,init ArrayList with input type
func NewArrayList(typeOf string) (*ArrayList, error) {
	if typeOf == "" {
		return nil, errors2.InvalidTypeError
	}
	return &ArrayList{
		data:   make([]interface{}, 0, DefaultCapacity),
		len:    0,
		typeOf: typeOf,
	}, nil
}

// NewArrayListWithCap ,init ArrayList with input type and capacity
func NewArrayListWithCap(typeOf string, capacity int) (*ArrayList, error) {
	if typeOf == "" {
		return nil, errors2.InvalidTypeError
	}
	if capacity < 0 {
		return nil, errors2.InvalidCapacityError
	}
	return &ArrayList{
		data:   make([]interface{}, capacity),
		len:    0,
		typeOf: typeOf,
	}, nil
}

// Add ,add a value to tail
func (arr *ArrayList) Add(val interface{}) error {
	if err := arr.Check(val); err != nil {
		return err
	}
	if len(arr.data) == arr.Len() {
		arr.expansion()
	}
	arr.data[arr.len] = val
	arr.len++
	return nil
}

// Check ,check input value type
func (arr *ArrayList) Check(val interface{}) error {
	if nil == val {
		return errors2.InputValueCannotBeNilError
	}
	if err := generic.CheckType(arr.typeOf, val); err != nil {
		return err
	}
	return nil
}

// Clear ,make ArrayList to empty
func (arr *ArrayList) Clear() {
	for i := 0; i < len(arr.data); i++ {
		arr.data[i] = nil
	}
}

// Contain ,determine whether the value is included in the ArrayList
func (arr *ArrayList) Contain(val interface{}) bool {
	if err := arr.Check(val); err != nil {
		return false
	}
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] == val {
			return true
		}
	}
	return false
}

// Data ,get ArrayList all element
func (arr *ArrayList) Data() []interface{} {
	return arr.data
}

// expansion ,ArrayList expansion
// When the number of elements in the ArrayList is less than 1024, the capacity becomes twice of the original ArrayList;
// When the number of elements in the ArrayList is greater than 1024, the capacity becomes 1.25 times of the original ArrayList.
func (arr *ArrayList) expansion() {
	if arr.len < 1024 {
		if arr.len == 0 {
			arr.data = make([]interface{}, 1)
			return
		}
		newData := make([]interface{}, 2*arr.len)
		for i := 0; i < arr.Len(); i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData
	} else {
		nl := float64(arr.Len()) * 1.25
		newData := make([]interface{}, int(nl))
		for i := 0; i < arr.Len(); i++ {
			newData[i] = arr.data[i]
		}
		arr.data = newData
	}
}

// Get ,get the element at the specified index
func (arr *ArrayList) Get(index int) (interface{}, error) {
	if arr.isIndexOutOfRange(index) {
		return "", errors2.IndexOutOfBoundsError
	}
	return arr.data[index], nil
}

// IndexOf ,return the index of the value in the ArrayList
func (arr *ArrayList) IndexOf(val interface{}) (int, error) {
	if err := arr.Check(val); err != nil {
		return -1, err
	}
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] == val {
			return i, nil
		}
	}
	return -1, errors2.NotExistError
}

// Insert ,insert the value at the specified index
func (arr *ArrayList) Insert(index int, val interface{}) error {
	if err := arr.Check(val); err != nil {
		return err
	}
	if len(arr.data) == arr.Len() {
		arr.expansion()
	}
	if index != arr.len && arr.isIndexOutOfRange(index) {
		return errors2.IndexOutOfBoundsError
	}
	for i := arr.len; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = val
	arr.len++
	return nil
}

// IsEmpty ,determine whether the ArrayList is an empty ArrayList
func (arr *ArrayList) IsEmpty() bool {
	return arr.len == 0
}

// isIndexOutOfRange ,determine whether the index of the ArrayList is out of bounds
func (arr *ArrayList) isIndexOutOfRange(index int) bool {
	return index >= len(arr.data) || index < 0
}

// Len ,get ArrayList element number
func (arr *ArrayList) Len() int {
	return arr.len
}

// MergeArrayList ,merge two ArrayList
// only append other ArrayList to this ArrayList tail
func (arr *ArrayList) MergeArrayList(other *ArrayList) *ArrayList {
	if other.len < 1 {
		return arr
	}
	for i := 0; i < other.len; i++ {
		_ = arr.Add(other.data[i])
	}
	return arr
}

// Remove ,remove and return the element with the specified index
func (arr *ArrayList) Remove(index int) (interface{}, error) {
	if arr.isIndexOutOfRange(index) {
		return "", errors2.IndexOutOfBoundsError
	}
	val := arr.data[index]
	for i := index; i < len(arr.data)-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.len--
	return val, nil
}

// Replace ,replace and return the new value with the specified index
func (arr *ArrayList) Replace(index int, val interface{}) (interface{}, error) {
	if err := arr.Check(val); err != nil {
		return nil, err
	}
	if arr.isIndexOutOfRange(index) {
		return "", errors2.IndexOutOfBoundsError
	}
	old := arr.data[index]
	arr.data[index] = val
	return old, nil
}

// Set ,set a value for the specified index
func (arr *ArrayList) Set(index int, val interface{}) error {
	if err := arr.Check(val); err != nil {
		return err
	}
	if arr.isIndexOutOfRange(index) {
		return errors2.IndexOutOfBoundsError
	}
	arr.data[index] = val
	return nil
}

// SubArrayList ,get sub ArrayList
func (arr *ArrayList) SubArrayList(start, end int) (*ArrayList, error) {
	if start > end {
		return nil, errors2.SubArrayListWithIndexError
	}
	if arr.isIndexOutOfRange(start) || arr.isIndexOutOfRange(end) {
		return nil, errors2.IndexOutOfBoundsError
	}
	return &ArrayList{
		data: arr.data[start:end],
		len:  end - start,
	}, nil
}
