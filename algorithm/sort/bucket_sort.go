package sort

// 桶排序

// BucketSort 桶排序
func BucketSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	// 创建桶
	max := getMax(arr)
	buckets := make([][]int, length)

	// 将数据分到桶中
	index := 0
	for i := 0; i < length; i++ {
		index = arr[i] * (length - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}

	// 桶内使用快速排序
	tmpPos := 0
	for i := 0; i < length; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			buckets[i] = QuickSort(buckets[i], 0, len(buckets[i])-1)
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return arr
}

// getMax 获取待排序数组中最大值
func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// BucketSortSimple 桶排序简单实现
func BucketSortSimple(source []int) []int {
	if len(source) <= 1 {
		return source
	}
	array := make([]int, getMax(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]]++
	}
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(source, c)
	return source
}
