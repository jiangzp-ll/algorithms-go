package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	in := []int{4, 9, 8, 7, 3, 1, 2}
	expect := []int{1, 2, 3, 4, 7, 8, 9}
	flag := false
	actual := QuickSort(in, 0, len(in)-1)
	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			flag = true
		}
	}
	if !flag {
		t.Log("快速排序成功")
	} else {
		t.Error("快速排序失败")
	}
}

func TestQuickSortTwo(t *testing.T) {
	in := []int{4, 9, 8, 7, 3, 1, 2}
	expect := []int{1, 2, 3, 4, 7, 8, 9}
	flag := false
	actual := QuickSortTwo(in, 0, len(in)-1)
	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			flag = true
		}
	}
	if !flag {
		t.Log("快速排序成功")
	} else {
		t.Error("快速排序失败")
	}
}
