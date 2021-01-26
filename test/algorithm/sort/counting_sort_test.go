package sort

import (
	"github.com/zepeng-jiang/go-basic-demo/algorithm/sort"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{1, 6, 3, 5, 8, 6, 4}
	ret := sort.CountingSort(arr)
	if ret[0] == 1 && ret[len(arr)-1] == 8 {
		t.Log("计数排序成功")
	} else {
		t.Error("计数排序失败")
	}
}

func TestCountingSortTwo(t *testing.T) {
	arr := []int{1, 6, 3, 5, 8, 6, 4}
	ret := sort.CountingSortTwo(arr)
	if ret[0] == 1 && ret[len(arr)-1] == 8 {
		t.Log("计数排序成功")
	} else {
		t.Error("计数排序失败")
	}
}