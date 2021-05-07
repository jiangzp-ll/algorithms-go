package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect := 9
	index := BinarySearch(arr, expect)
	if index >= 0 && arr[index] == expect {
		t.Log("在数组中找到该元素")
	} else if index == -1 {
		t.Error("数组中不存在该元素")
	} else {
		t.Error("二分查找失败")
	}
}

func TestBinarySearchFindFirst(t *testing.T) {
	arr := []int{1, 3, 4, 4, 4, 4, 6, 8, 9}
	expect := 4
	index := BinarySearchFindFirst(arr, expect)
	if index == 2 && arr[index] == expect {
		t.Log("找到给定值，并且是第一个")
	} else if index != 2 && index > 0 && arr[index] == expect {
		t.Error("找到给定值，但是不是第一个")
	} else if index == -1 {
		t.Error("数组中不存在该元素")
	} else {
		t.Error("二分查找失败")
	}
}

func TestBinarySearchFindLast(t *testing.T) {
	arr := []int{1, 3, 4, 4, 4, 4, 6, 8, 9}
	expect := 4
	index := BinarySearchFindLast(arr, expect)
	if index == 5 && arr[index] == expect {
		t.Log("找到给定值，并且是最后一个")
	} else if index != 5 && index > 0 && arr[index] == expect {
		t.Error("找到给定值，但是不是最后一个")
	} else if index == -1 {
		t.Error("数组中不存在该元素")
	} else {
		t.Error("二分查找失败")
	}
}

func TestBinarySearchFirstGT(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect := 5
	index := BinarySearchFirstGT(arr, expect)
	if index == 3 && arr[index] > expect {
		t.Log("在数组中找到第一个大于等于给定值的元素")
	} else if index != 3 && index > 0 && arr[index] > expect {
		t.Error("找到的不是第一个大于等于给定值的元素")
	} else if index == -1 {
		t.Error("数组中不存在该元素")
	} else {
		t.Error("二分查找失败")
	}
}

func TestBinarySearchLastLT(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect := 5
	index := BinarySearchLastLT(arr, expect)
	if index == 2 && arr[index] < expect {
		t.Log("在数组中找到最后一个大于等于给定值的元素")
	} else if index != 2 && index > 0 && arr[index] < expect {
		t.Error("找到的不是最后一个大于等于给定值的元素")
	} else if index == -1 {
		t.Error("数组中不存在该元素")
	} else {
		t.Error("二分查找失败")
	}
}
