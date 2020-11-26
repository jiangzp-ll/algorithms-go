package recursion

import (
	"github.com/zepeng-jiang/go-basic-demo/main/algorithm/recursion"
	"testing"
)

func TestRangeType_RangeAll(t *testing.T) {
	input := []interface{}{"A", "B", "C"}
	ra := recursion.NewRangeArray(input)
	ra.RangeAll(0)
}

