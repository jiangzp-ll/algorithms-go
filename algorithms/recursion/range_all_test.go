package recursion

import (
	"testing"
)

func TestRangeType_RangeAll(t *testing.T) {
	input := []interface{}{"A", "B", "C"}
	ra := NewRangeArray(input)
	ra.RangeAll(0)
}

