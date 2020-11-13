package sort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
