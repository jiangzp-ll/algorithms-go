package main

import "testing"

func TestRangeType_RangeAll(t *testing.T) {
	ra := NewRangeArray(3)
	ra.value[0] = "A"
	ra.value[1] = "B"
	ra.value[2] = "C"
	ra.RangeAll(0)
}

