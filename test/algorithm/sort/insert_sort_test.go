package sort

import (
	"github.com/zepeng-jiang/go-basic-demo/main/algorithm/sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	in := []int{9, 4, 8, 7, 3, 1, 2}
	expect := []int{1, 2, 3, 4, 7, 8, 9}
	flag := false
	actual := sort.InsertSort(in)
	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			flag = true
		}
	}
	if !flag {
		t.Log("插入排序成功")
	} else {
		t.Error("插入排序失败")
	}
}