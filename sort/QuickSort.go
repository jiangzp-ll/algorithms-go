package sort

func QuickSort(arr []int, start, end int) []int {
	if start < end {
		//选定基准
		pivot := arr[start]
		// 双边循环
		i,j := start, end
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
		QuickSort(arr, start, i - 1)
		QuickSort(arr, i + 1, end)
	}
	return arr
}
