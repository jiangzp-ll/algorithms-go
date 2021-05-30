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
		for j := 0; j < length-1; j++ {
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
		for j := 0; j < length-1; j++ {
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
	for i := 0; i < length; i++ {
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
	for i := 0; i < length; i++ {
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