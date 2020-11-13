package sort

// SelectSort 选择排序
func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
