package sort

import (
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	in := []int{9, 4, 8, 7, 3, 1, 2}
	expect := []int{1, 2, 3, 4, 7, 8, 9}
	flag := true
	actual := BubbleSort(in)
	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			flag = false
		}
	}
	if flag {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_BubbleSort_With_Input_Only_One_Element(t *testing.T) {
	in := []int{9}
	expect := []int{9}
	actual := BubbleSort(in)
	if actual[0] == expect[0] && len(actual) == 1 {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}
