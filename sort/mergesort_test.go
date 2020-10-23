package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	in := []int{9, 4, 8, 7, 3, 1, 2}
	actual := MergeSort(in)
	t.Log(actual)
}
