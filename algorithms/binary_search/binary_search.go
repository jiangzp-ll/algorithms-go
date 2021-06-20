package binary_search

import (
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
)

// BinarySearchOfTypeInt ,slice binary search of type int
func BinarySearchOfTypeInt(arr []int, t int) (int, error) {
	if len(arr) == 0 {
		return -1, errors2.InputSliceIsEmptyError
	}
	var mid int
	low, high := 0, len(arr)-1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == t {
			return mid, nil
		}
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
	}
	return -1, errors2.NotExistError
}

// BinarySearchOfTypeFloat64 ,slice binary search of type float64
func BinarySearchOfTypeFloat64(arr []float64, t float64) (int, error) {
	if len(arr) == 0 {
		return -1, errors2.InputSliceIsEmptyError
	}
	var mid int
	low, high := 0, len(arr)-1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == t {
			return mid, nil
		}
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
	}
	return -1, errors2.NotExistError
}

// BinarySearchFindFirst .
// Finds the first element equal to the given value.
func BinarySearchFindFirst(arr []int, t int) (int, error) {
	if len(arr) == 0 {
		return -1, errors2.InputSliceIsEmptyError
	}
	var mid int
	low, high := 0, len(arr)-1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == t {
			if mid == 0 || arr[mid-1] != t {
				return mid, nil
			} else {
				high = mid - 1
			}
		}
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
	}
	return -1, errors2.NotExistError
}

// BinarySearchFindLast .
// Finds the last element equal to the given value.
func BinarySearchFindLast(arr []int, t int) (int, error) {
	if len(arr) == 0 {
		return -1, errors2.InputSliceIsEmptyError
	}
	var mid int
	low, high := 0, len(arr)-1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == t {
			if mid == len(arr)-1 || arr[mid+1] != t {
				return mid, nil
			} else {
				low = mid + 1
			}
		}
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
	}
	return -1, errors2.NotExistError
}

// BinarySearchFirstGT .
// Find the first element greater than to the given value
func BinarySearchFirstGT(arr []int, t int) (int, error) {
	if len(arr) == 0 {
		return -1, errors2.InputSliceIsEmptyError
	}
	var mid int
	low, high := 0, len(arr)-1
	if arr[high] <= t {
		return -1, errors2.InvalidInputValueError
	}
	if arr[low] >= t {
		return low, nil
	}
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] < t {
			low = mid + 1
		}
		if mid != len(arr)-1 && arr[mid+1] > t {
			mid = mid + 1
			break
		} else {
			low = mid + 1
		}
	}
	return mid, nil
}

// BinarySearchLastLT .
// Find the last element less than to the given value
func BinarySearchLastLT(arr []int, t int) (int, error) {
	if len(arr) == 0 {
		return -1, errors2.InputSliceIsEmptyError
	}
	var mid int
	low, high := 0, len(arr)-1
	if arr[low] >= t {
		return -1, errors2.InvalidInputValueError
	}
	if arr[high] <= t {
		return high, nil
	}
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] >= t {
			high = mid - 1
		}
		if mid == 0 || (arr[mid] < t && arr[mid+1] >= t) {
			break
		}
	}
	return mid, nil
}
