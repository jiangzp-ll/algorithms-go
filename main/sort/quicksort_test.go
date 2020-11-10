package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	in := []int{9, 4, 8, 7, 3, 1, 2}
	result := QuickSort(in, 0, len(in)-1)
	t.Log(result)
}
