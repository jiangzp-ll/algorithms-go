package sort

import (
	"github.com/zepeng-jiang/go-basic-demo/algorithm/sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	in := []int{9, 4, 8, 7, 3, 1, 2}
	expect := []int{1, 2, 3, 4, 7, 8, 9}
	flag := false
	actual := sort.MergeSort(in)
	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			flag = true
		}
	}
	if !flag {
		t.Log("归并排序成功")
	} else {
		t.Error("归并排序失败")
	}
}
