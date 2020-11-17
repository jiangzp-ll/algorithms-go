package sort

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{1, 6, 3, 5, 8, 6, 4}
	ret := CountingSortTwo(arr)
	if ret[0] == 1 && ret[len(arr)-1] == 8 {
		t.Log("计数排序成功")
	} else {
		t.Error("计数排序失败")
	}
}