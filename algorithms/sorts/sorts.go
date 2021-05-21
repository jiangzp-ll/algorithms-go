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
