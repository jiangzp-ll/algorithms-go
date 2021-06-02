package sorts

import "testing"

func Test_BubbleSortOfTypeInt(t *testing.T) {
	target := []int{1, 2, 3}
	input := []int{3, 1, 2}
	flag := true
	ret := BubbleSortOfTypeInt(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_BubbleSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := BubbleSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_BubbleSortOfTypeFloat64(t *testing.T) {
	target := []float64{1.1, 2.2, 3.3}
	input := []float64{3.3, 1.1, 2.2}
	flag := true
	ret := BubbleSortOfTypeFloat64(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_BubbleSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := BubbleSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_SelectionSortOfTypeInt(t *testing.T) {
	target := []int{1, 2, 3}
	input := []int{3, 1, 2}
	flag := true
	ret := SelectionSortOfTypeInt(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("selection sort is success")
	} else {
		t.Error("selection sort is failed")
	}
}

func Test_SelectionSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := SelectionSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("selection sort is success")
	} else {
		t.Error("selection sort is failed")
	}
}

func Test_SelectionSortOfTypeFloat64(t *testing.T) {
	target := []float64{1.1, 2.2, 3.3}
	input := []float64{3.3, 1.1, 2.2}
	flag := true
	ret := SelectionSortOfTypeFloat64(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("selection sort is success")
	} else {
		t.Error("selection sort is failed")
	}
}

func Test_SelectionSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := SelectionSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("selection sort is success")
	} else {
		t.Error("selection sort is failed")
	}
}

func Test_InsertionSortOfTypeInt(t *testing.T) {
	target := []int{1, 2, 3}
	input := []int{3, 1, 2}
	flag := true
	ret := InsertionSortOfTypeInt(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("insertion sort is success")
	} else {
		t.Error("insertion sort is failed")
	}
}

func Test_InsertionSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := InsertionSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("insertion sort is success")
	} else {
		t.Error("insertion sort is failed")
	}
}

func Test_InsertionSortOfTypeFloat64(t *testing.T) {
	target := []float64{1.1, 2.2, 3.3}
	input := []float64{3.3, 1.1, 2.2}
	flag := true
	ret := InsertionSortOfTypeFloat64(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("insertion sort is success")
	} else {
		t.Error("insertion sort is failed")
	}
}

func Test_InsertionSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := InsertionSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("insertion sort is success")
	} else {
		t.Error("insertion sort is failed")
	}
}

func Test_ShellSortOfTypeInt_With_Array_Number_Is_Odd(t *testing.T) {
	target := []int{1, 2, 3}
	input := []int{3, 1, 2}
	flag := true
	ret := ShellSortOfTypeInt(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("shell sort is success")
	} else {
		t.Error("shell sort is failed")
	}
}

func Test_ShellSortOfTypeInt_With_Array_Number_Is_Even(t *testing.T) {
	target := []int{1, 2, 3, 4}
	input := []int{4, 3, 1, 2}
	flag := true
	ret := ShellSortOfTypeInt(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("shell sort is success")
	} else {
		t.Error("shell sort is failed")
	}
}

func Test_ShellSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := ShellSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("shell sort is success")
	} else {
		t.Error("shell sort is failed")
	}
}

func Test_ShellSortOfTypeFloat64_With_Array_Number_Is_Odd(t *testing.T) {
	target := []float64{1.1, 2.2, 3.3}
	input := []float64{3.3, 1.1, 2.2}
	flag := true
	ret := ShellSortOfTypeFloat64(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("shell sort is success")
	} else {
		t.Error("shell sort is failed")
	}
}

func Test_ShellSortOfTypeFloat64_With_Array_Number_Is_Even(t *testing.T) {
	target := []float64{1.1, 2.2, 3.3, 4.4}
	input := []float64{4.4, 3.3, 1.1, 2.2}
	flag := true
	ret := ShellSortOfTypeFloat64(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("shell sort is success")
	} else {
		t.Error("shell sort is failed")
	}
}

func Test_ShellSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := ShellSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("shell sort is success")
	} else {
		t.Error("shell sort is failed")
	}
}

func Test_MergeSortOfTypeInt(t *testing.T) {
	target := []int{1, 2, 3}
	input := []int{3, 1, 2}
	flag := true
	ret := MergeSortOfTypeInt(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("merge sort is success")
	} else {
		t.Error("merge sort is failed")
	}
}

func Test_MergeSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := MergeSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("merge sort is success")
	} else {
		t.Error("merge sort is failed")
	}
}

func Test_MergeSortOfTypeFloat64(t *testing.T) {
	target := []float64{1.1, 2.2, 3.3}
	input := []float64{3.3, 1.1, 2.2}
	flag := true
	ret := MergeSortOfTypeFloat64(input)
	for i := 0; i < len(target); i++ {
		if target[i] != ret[i] {
			flag = false
		}
	}
	if flag {
		t.Log("merge sort is success")
	} else {
		t.Error("merge sort is failed")
	}
}

func Test_MergeSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := MergeSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("merge sort is success")
	} else {
		t.Error("merge sort is failed")
	}
}
