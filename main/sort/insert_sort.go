package sort

// InsertSort 插入排序
func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
