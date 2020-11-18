package sort

// QuickSort 快速排序方式一:双指针法
func QuickSort(arr []int, start, end int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if start < end {
		//选定基准
		pivot := arr[start]
		// 双边循环
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
		arr[start] = arr[i]
		arr[i] = pivot

		//递归调用
		QuickSort(arr, start, i-1)
		QuickSort(arr, i+1, end)
	}
	return arr
}

// QuickSortTwo 快速排序方法二：单指针法
func QuickSortTwo(arr []int, start, end int) []int {
	if start >= end {
		return nil
	}
	i := partition(arr, start, end)
	QuickSortTwo(arr, start, i-1)
	QuickSortTwo(arr, i+1, end)
	return arr
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
