package sort

import (
	"github.com/zepeng-jiang/go-basic-demo/main/algorithm/sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []int{1, 6, 3, 5, 8, 6, 4}
	ret := sort.BucketSort(arr)
	if ret[0] == 1 && ret[len(arr)-1] == 8 {
		t.Log("桶排序成功")
	} else {
		t.Error("桶排序失败")
	}
}

func TestBucketSortSimple(t *testing.T) {
	arr := []int{1, 6, 3, 5, 8, 6, 4}
	ret := sort.BucketSortSimple(arr)
	if ret[0] == 1 && ret[len(arr)-1] == 8 {
		t.Log("桶排序成功")
	} else {
		t.Error("桶排序失败")
	}
}
