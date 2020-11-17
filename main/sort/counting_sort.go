package sort

// 计数排序

// CountingSort 计数排序
func CountingSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	max := getMax(arr)
	ret := make([]int, max+1)
	for i := 0; i < length; i++ {
		ret[arr[i]] += 1
	}
	j := 0
	for i := 0; i < max+1; i++ {
		for ret[i] > 0 {
			arr[j] = i
			j++
			ret[i]--
		}
	}
	return arr
}

// CountingSortTwo 计数排序方式二
func CountingSortTwo(arr []int) []int {
	// 1. 获取待排序数组的最大值
	max := getMax(arr)
	// 2. 构建统计数组
	countArr := make([]int, max+1)
	// 3. 遍历待排序数组，填充统计数组
	for _, a := range arr {
		countArr[a]++
	}
	// 4. 遍历统计数组，输出结果
	index := 0
	ret := make([]int, len(arr))
	for i := 0; i < len(countArr); i++ {
		if countArr[i] > 0 {
			for j := 0; j < countArr[i]; j++ {
				ret[index] = i
				index++
			}
		}
	}
	return ret

}
