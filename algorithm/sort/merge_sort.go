package sort

// MergeSort 归并排序
func MergeSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	mid := len(in) >> 1
	left := in[:mid]
	right := in[mid:]

	l := MergeSort(left)
	r := MergeSort(right)

	result := merge(l, r)
	return result
}

// merge 合并两个切片
func merge(left []int, right []int) []int {
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
