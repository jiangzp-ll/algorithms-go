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
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_SelectionSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := SelectionSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
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
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_SelectionSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := SelectionSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
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
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_InsertionSortOfTypeInt_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []int{1}
	input := []int{1}
	ret := InsertionSortOfTypeInt(input)
	if ret[0] == target[0] {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
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
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}

func Test_InsertionSortOfTypeFloat64_With_Input_Array_Has_One_Element(t *testing.T) {
	target := []float64{1.1}
	input := []float64{1.1}
	ret := InsertionSortOfTypeFloat64(input)
	if ret[0] == target[0] {
		t.Log("bubble sort is success")
	} else {
		t.Error("bubble sort is failed")
	}
}
