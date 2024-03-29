package sorts

// This section contains some common sorting algorithms.
// Only array sorting of type int or float is supported

// BubbleSortOfTypeInt , bubble sort with input array type is int
func BubbleSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// BubbleSortOfTypeFloat64 , bubble sort with input array type is float64
func BubbleSortOfTypeFloat64(arr []float64) []float64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// SelectionSortOfTypeInt , selection sort with input array type is int
func SelectionSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[i] {
				index = j
				break
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}

// SelectionSortOfTypeInt , selection sort with input array type is float64
func SelectionSortOfTypeFloat64(arr []float64) []float64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[i] {
				index = j
				break
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}

// InsertionSortOfTypeInt , insertion sort with input array type is int
func InsertionSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

// InsertionSortOfTypeFloat64 , insertion sort with input array type is float64
func InsertionSortOfTypeFloat64(arr []float64) []float64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

// ShellSortOfTypeInt , shell sort with input array type is int
func ShellSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	gap := length >> 1
	mid := gap
	for i := 0; i < mid; i++ {
		for j := 0; j <= mid; j++ {
			if length%2 == 0 && j == mid && gap != 1 {
				break
			}
			if arr[j] > arr[j+gap] {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
		gap = gap >> 1
	}
	return arr
}

// ShellSortOfTypeFloat64 , shell sort with input array type is float64
func ShellSortOfTypeFloat64(arr []float64) []float64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	gap := length >> 1
	mid := gap
	for i := 0; i < mid; i++ {
		for j := 0; j <= mid; j++ {
			if length%2 == 0 && j == mid && gap != 1 {
				break
			}
			if arr[j] > arr[j+gap] {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
		gap = gap >> 1
	}
	return arr
}

// MergeSortOfTypeInt , merge sort with input array type is int
func MergeSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length >> 1
	left, right := MergeSortOfTypeInt(arr[:mid]), MergeSortOfTypeInt(arr[mid:])
	return mergeOfInt(left, right)
}

// mergeOfInt ,merge two arrays of type int
func mergeOfInt(left []int, right []int) []int {
	ret := make([]int, len(left)+len(right))
	var l, r, x int
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			ret[x] = left[l]
			l++
		} else {
			ret[x] = right[r]
			r++
		}
		x++
	}
	if l < len(left) {
		for i := l; i < len(left); i++ {
			ret[x] = left[i]
			x++
		}
	}
	if r < len(right) {
		for i := r; i < len(right); i++ {
			ret[x] = right[i]
			x++
		}
	}
	return ret
}

// MergeSortOfTypeFloat64 , merge sort with input array type is float64
func MergeSortOfTypeFloat64(arr []float64) []float64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length >> 1
	left, right := MergeSortOfTypeFloat64(arr[:mid]), MergeSortOfTypeFloat64(arr[mid:])
	return mergeOfFloat64(left, right)
}

// mergeOfFloat64 ,merge two arrays of type float64
func mergeOfFloat64(left []float64, right []float64) []float64 {
	ret := make([]float64, len(left)+len(right))
	var l, r, x int
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			ret[x] = left[l]
			l++
		} else {
			ret[x] = right[r]
			r++
		}
		x++
	}
	if l < len(left) {
		for i := l; i < len(left); i++ {
			ret[x] = left[i]
			x++
		}
	}
	if r < len(right) {
		for i := r; i < len(right); i++ {
			ret[x] = right[i]
			x++
		}
	}
	return ret
}

// QuickSortOfTypeInt , quick sort with input array type is int
// start and end represent the starting positions of the two pointers in the array.
// The value of start is 0, and the value of end is len(arr)-1.
func QuickSortOfTypeInt(arr []int, start, end int) []int {
	if start >= end {
		return arr
	}
	pivot := arr[start]
	i, j := start, end
	for {
		for arr[j] >= pivot && i < j {
			j--
		}
		for arr[i] <= pivot && i < j {
			i++
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[start], arr[i] = arr[i], pivot
	QuickSortOfTypeInt(arr, 0, i-1)
	QuickSortOfTypeInt(arr, i+1, end)
	return arr
}

// QuickSortOfTypeFloat64 , quick sort with input array type is float64
// start and end represent the starting positions of the two pointers in the array.
// The value of start is 0, and the value of end is len(arr)-1.
func QuickSortOfTypeFloat64(arr []float64, start, end int) []float64 {
	if start >= end {
		return arr
	}
	pivot := arr[start]
	i, j := start, end
	for {
		for arr[j] >= pivot && i < j {
			j--
		}
		for arr[i] <= pivot && i < j {
			i++
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[start], arr[i] = arr[i], pivot
	QuickSortOfTypeFloat64(arr, 0, i-1)
	QuickSortOfTypeFloat64(arr, i+1, end)
	return arr
}

// HeapSortOfTypeInt , heap sort with input array type is int
func HeapSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		a := arr[i:]
		minHeapOfInt(a)
	}
	return arr
}

// minHeapTypeOfInt ,create a minimum heap of type int
func minHeapOfInt(arr []int) {
	length := len(arr)
	floor := length/2 - 1
	for i := floor; i >= 0; i-- {
		// 'fa' represents the parent node
		fa := i
		left, right := 2*i+1, 2*i+2
		if right < length && arr[fa] > arr[right] {
			arr[fa], arr[right] = arr[right], arr[fa]
		}
		if left < length && arr[fa] > arr[left] {
			arr[fa], arr[left] = arr[left], arr[fa]
		}
	}
}

// HeapSortOfTypeFloat64 , heap sort with input array type is float64
func HeapSortOfTypeFloat64(arr []float64) []float64 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		a := arr[i:]
		minHeapOfFloat64(a)
	}
	return arr
}

// minHeapTypeOfFloat64 ,create a minimum heap of type float64
func minHeapOfFloat64(arr []float64) {
	length := len(arr)
	floor := length/2 - 1
	for i := floor; i >= 0; i-- {
		// 'fa' represents the parent node
		fa := i
		left, right := 2*i+1, 2*i+2
		if right < length && arr[fa] > arr[right] {
			arr[fa], arr[right] = arr[right], arr[fa]
		}
		if left < length && arr[fa] > arr[left] {
			arr[fa], arr[left] = arr[left], arr[fa]
		}
	}
}

// CountingSortOfTypeInt , counting sort with input array type is int
func CountingSortOfTypeInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	// get min and max value in the Array
	min, max := getMinAndMax(arr, length)
	// fill counting array
	ca := make([]int, max-min+1)
	for i := 0; i < length; i++ {
		index := arr[i] - min
		ca[index] = ca[index] + 1
	}
	// replace the original array
	index := 0
	for i := 0; i < len(ca); i++ {
		for j := 0; j < ca[i]; j++ {
			arr[index] = i + min
			index++
		}
	}
	return arr
}

// getMinAndMax ,get min and max value in the Array type of int
func getMinAndMax(arr []int, length int) (int, int) {
	min, max := arr[0], arr[0]
	for i := 0; i < length; i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	return min, max
}
