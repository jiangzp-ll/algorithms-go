package binary_search

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

func TestBinarySearchOfTypeInt(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect := 9
	index, err := BinarySearchOfTypeInt(arr, expect)
	if err != nil {
		t.Errorf("binary search has error! error: %v ", err.Error())
		return
	}
	if index >= 0 && arr[index] == expect {
		t.Logf("found %d int the slice", expect)
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeInt_With_Expect_Value_Is_Mid(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect := 3
	index, err := BinarySearchOfTypeInt(arr, expect)
	if err != nil {
		t.Errorf("binary search has error! error: %v ", err.Error())
		return
	}
	if index >= 0 && arr[index] == expect {
		t.Logf("found %d int the slice", expect)
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeInt_With_Slice_Only_One(t *testing.T) {
	arr := []int{9}
	expect := 9
	index, err := BinarySearchOfTypeInt(arr, expect)
	if err != nil {
		t.Errorf("binary search has error! error: %v ", err.Error())
		return
	}
	if index >= 0 && arr[index] == expect {
		t.Logf("found %d int the slice", expect)
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeInt_With_Slice_Is_Empty(t *testing.T) {
	arr := make([]int, 0)
	expect := 9
	_, err := BinarySearchOfTypeInt(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.InputSliceIsEmptyError) {
			t.Log("input slice not has element")
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeInt_With_Slice_Only_One_And_Not_Exited_The_Value(t *testing.T) {
	arr := []int{1}
	expect := 9
	_, err := BinarySearchOfTypeInt(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Logf("%d does not exist in the slice", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeInt_With_Slice_Not_Exited_The_Value(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8}
	expect := 9
	_, err := BinarySearchOfTypeInt(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Logf("%d does not exist in the slice", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeFloat64(t *testing.T) {
	arr := []float64{1.0, 3.0, 4.0, 6.0, 8.0, 9.0}
	expect := 3.0
	index, err := BinarySearchOfTypeFloat64(arr, expect)
	if err != nil {
		t.Errorf("binary search has error! error: %v ", err.Error())
		return
	}
	if index >= 0 && arr[index] == expect {
		t.Logf("found %f int the slice", expect)
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeFloat64_With_Slice_Only_One(t *testing.T) {
	arr := []float64{9.0}
	expect := 9.0
	index, err := BinarySearchOfTypeFloat64(arr, expect)
	if err != nil {
		t.Errorf("binary search has error! error: %v ", err.Error())
		return
	}
	if index >= 0 && arr[index] == expect {
		t.Logf("found %f int the slice", expect)
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeFloat64_With_Slice_Is_Empty(t *testing.T) {
	arr := make([]float64, 0)
	expect := 9.0
	_, err := BinarySearchOfTypeFloat64(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.InputSliceIsEmptyError) {
			t.Log("input slice not has element")
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeFloat64_With_Slice_Only_One_And_Not_Exited_The_Value(t *testing.T) {
	arr := []float64{1.0}
	expect := 9.0
	_, err := BinarySearchOfTypeFloat64(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Logf("%f does not exist in the slice", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchOfTypeFloat64_With_Slice_Not_Exited_The_Value(t *testing.T) {
	arr := []float64{1.0, 3.0, 4.0, 6.0, 8.0}
	expect := 9.0
	_, err := BinarySearchOfTypeFloat64(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Logf("%f does not exist in the slice", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearch has bug")
	}
}

func TestBinarySearchFindFirst(t *testing.T) {
	arr := []int{1, 3, 4, 4, 4, 4, 6, 8, 9}
	expect, index := 4, 2
	ret, err := BinarySearchFindFirst(arr, expect)
	if err != nil {
		t.Errorf("find first equal the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the first index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchFindFirst has bug")
	}
}

func TestBinarySearchFindFirst_With_Slice_Not_Exited_The_Value(t *testing.T) {
	arr := []int{1, 3, 4, 4, 4, 4, 6, 8, 9}
	expect := 2
	if _, err := BinarySearchFindFirst(arr, expect); err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Logf("%d does not exist in the slice", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchFindFirst has bug")
	}
}

func TestBinarySearchFindFirst_With_Slice_Is_Empty(t *testing.T) {
	arr := make([]int, 0)
	expect := 9
	if _, err := BinarySearchFindFirst(arr, expect); err != nil {
		if errors.Is(err, errors2.InputSliceIsEmptyError) {
			t.Log("input slice not has element")
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchFindFirst has bug")
	}
}

func TestBinarySearchFindLast(t *testing.T) {
	arr := []int{1, 3, 4, 4, 4, 4, 6, 8, 9}
	expect, index := 4, 5
	ret, err := BinarySearchFindLast(arr, expect)
	if err != nil {
		t.Errorf("find last equal the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the last index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchFindLast has bug")
	}
}

func TestBinarySearchFindLast_With_Slice_Not_Exited_The_Value(t *testing.T) {
	arr := []int{1, 3, 4, 4, 4, 4, 6, 8, 9}
	expect := 2
	if _, err := BinarySearchFindLast(arr, expect); err != nil {
		if errors.Is(err, errors2.NotExistError) {
			t.Logf("%d does not exist in the slice", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchFindLast has bug")
	}
}

func TestBinarySearchFindLast_With_Slice_Is_Empty(t *testing.T) {
	arr := make([]int, 0)
	expect := 9
	if _, err := BinarySearchFindLast(arr, expect); err != nil {
		if errors.Is(err, errors2.InputSliceIsEmptyError) {
			t.Log("input slice not has element")
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchFindLast has bug")
	}
}

func TestBinarySearchFirstGT(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect, index := 5, 3
	ret, err := BinarySearchFirstGT(arr, expect)
	if err != nil {
		t.Errorf("find first greater than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the first greater than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchFirstGT has bug")
	}
}

func TestBinarySearchFirstGT_With_Slice_Contains_The_Value(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect, index := 6, 5
	ret, err := BinarySearchFirstGT(arr, expect)
	if err != nil {
		t.Errorf("find first greater than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the first greater than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchFirstGT has bug")
	}
}

func TestBinarySearchFirstGT_With_Input_Value_Less_Than_Slice_First_Element(t *testing.T) {
	arr := []int{9, 10, 11, 12}
	expect, index := 8, 0
	ret, err := BinarySearchFirstGT(arr, expect)
	if err != nil {
		t.Errorf("find first greater than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the first greater than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchFirstGT has bug")
	}
}

func TestBinarySearchFirstGT_With_Slice_Is_Empty(t *testing.T) {
	arr := make([]int, 0)
	expect := 9
	_, err := BinarySearchFirstGT(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.InputSliceIsEmptyError) {
			t.Log("input slice not has element")
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchFirstGT has bug")
	}
}

func TestBinarySearchFirstGT_With_Input_Value_Greater_Than_Slice_Last_Element(t *testing.T) {
	arr := []int{9, 10, 11, 12}
	expect := 13
	_, err := BinarySearchFirstGT(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.InvalidInputValueError) {
			t.Logf("%d greater than the slice the biggest element", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchFirstGT has bug")
	}
}

func TestBinarySearchLastLT(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect, index := 8, 3
	ret, err := BinarySearchLastLT(arr, expect)
	if err != nil {
		t.Errorf("find last less than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the last less than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchLastLT has bug")
	}
}

func TestBinarySearchLastLT_With_Input_Value_Less_Than_Mid(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect, index := 2, 0
	ret, err := BinarySearchLastLT(arr, expect)
	if err != nil {
		t.Errorf("find last less than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the last less than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchLastLT has bug", ret)
	}
}

func TestBinarySearchLastLT_With_Slice_Contains_The_Value(t *testing.T) {
	arr := []int{1, 3, 4, 6, 8, 9}
	expect, index := 6, 2
	ret, err := BinarySearchLastLT(arr, expect)
	if err != nil {
		t.Errorf("find last less than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the last less than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchLastLT has bug")
	}
}

func TestBinarySearchLastLT_With_Input_Value_Greater_Than_Slice_Last_Element(t *testing.T) {
	arr := []int{9, 10, 11, 12}
	expect, index := 13, len(arr)-1
	ret, err := BinarySearchLastLT(arr, expect)
	if err != nil {
		t.Errorf("find last less than the value has error! error: %v ", err)
		return
	}
	if ret == index {
		t.Logf("the last less than the value index of %d in slice is %d", expect, ret)
	} else {
		t.Error("function BinarySearchLastLT has bug")
	}
}

func TestBinarySearchLastLT_With_Slice_Is_Empty(t *testing.T) {
	arr := make([]int, 0)
	expect := 9
	_, err := BinarySearchLastLT(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.InputSliceIsEmptyError) {
			t.Log("input slice not has element")
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchLastLT has bug")
	}
}

func TestBinarySearchLastLT_With_Input_Value_Less_Than_Slice_First_Element(t *testing.T) {
	arr := []int{9, 10, 11, 12}
	expect := 8
	_, err := BinarySearchLastLT(arr, expect)
	if err != nil {
		if errors.Is(err, errors2.InvalidInputValueError) {
			t.Logf("%d less than the slice the biggest element", expect)
			return
		} else {
			t.Errorf("unknown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function BinarySearchLastLT has bug")
	}
}
