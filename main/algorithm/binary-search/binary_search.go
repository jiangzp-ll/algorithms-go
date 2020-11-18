package binary_search

// BinarySearch 二分查找法
func BinarySearch(arr []int, t int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if arr[0] == t {
			return 0
		} else {
			return -1
		}
	}
	var mid int
	low := 0
	high := length - 1
	for low <= high {
		mid = low + ((high - low) >> 2)
		if arr[mid] == t {
			return mid
		}
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
	}
	return -1
}

// BinarySearchFindFirst 查找第一个等于给定值的元素
func BinarySearchFindFirst(arr []int, t int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if arr[0] == t {
			return 0
		} else {
			return -1
		}
	}
	var mid int
	low := 0
	high := length - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == t {
			if mid == 0 || arr[mid-1] != t {
				return mid
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
	return -1
}

// BinarySearchFindLast 查找最后一个等于给定值的元素
func BinarySearchFindLast(arr []int, t int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if arr[0] == t {
			return 0
		} else {
			return -1
		}
	}
	var mid int
	low := 0
	high := length - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == t {
			if mid == length-1 || arr[mid+1] != t {
				return mid
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
	return -1
}

// BinarySearchFirstGT 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(arr []int, t int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if arr[0] == t {
			return 0
		} else {
			return -1
		}
	}
	var mid int
	low := 0
	high := length - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
		if mid != length-1 && arr[mid+1] > t {
			return mid + 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// BinarySearchLastLT 查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, t int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if arr[0] == t {
			return 0
		} else {
			return -1
		}
	}
	var mid int
	low := 0
	high := length - 1
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] < t {
			low = mid + 1
		}
		if arr[mid] > t {
			high = mid - 1
		}
		if mid == 0 || arr[mid] < t {
			return mid
		} else {
			high = mid - 1
		}
	}
	return -1
}