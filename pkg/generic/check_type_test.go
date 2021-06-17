package generic

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

func TestCheckType_With_String(t *testing.T) {
	typeOf := "string"
	var in string
	in = "a"
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of string is success")
	}
}

func TestCheckType_With_String_And_Input_Other_Type(t *testing.T) {
	typeOf := "string"
	var in int
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Int(t *testing.T) {
	typeOf := "int"
	var in int
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of int is success")
	}
}

func TestCheckType_With_Int_And_Input_Other_Type(t *testing.T) {
	typeOf := "int"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Int32(t *testing.T) {
	typeOf := "int32"
	var in int32
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of int32 is success")
	}
}

func TestCheckType_With_Int32_And_Input_Other_Type(t *testing.T) {
	typeOf := "int32"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Int64(t *testing.T) {
	typeOf := "int64"
	var in int64
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of int64 is success")
	}
}

func TestCheckType_With_Int64_And_Input_Other_Type(t *testing.T) {
	typeOf := "int64"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Uint(t *testing.T) {
	typeOf := "uint"
	var in uint
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of uint is success")
	}
}

func TestCheckType_With_Uint_And_Input_Other_Type(t *testing.T) {
	typeOf := "uint"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Uint32(t *testing.T) {
	typeOf := "uint32"
	var in uint32
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of uint32 is success")
	}
}

func TestCheckType_With_Uint32_And_Input_Other_Type(t *testing.T) {
	typeOf := "uint32"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Uint64(t *testing.T) {
	typeOf := "uint64"
	var in uint64
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of int64 is success")
	}
}

func TestCheckType_With_Uint64_And_Input_Other_Type(t *testing.T) {
	typeOf := "uint64"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Float32(t *testing.T) {
	typeOf := "float32"
	var in float32
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of float32 is success")
	}
}

func TestCheckType_With_Float32_And_Input_Other_Type(t *testing.T) {
	typeOf := "float32"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Float64(t *testing.T) {
	typeOf := "float64"
	var in float64
	in = 1
	if err := CheckType(typeOf, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		t.Log("check type of float64 is success")
	}
}

func TestCheckType_With_Float64_And_Input_Other_Type(t *testing.T) {
	typeOf := "float64"
	var in string
	in = "1"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Slice(t *testing.T) {
	typeOfSlice, typeOfElement := "slice", "int"
	in := []int{1, 2, 3}
	if err := CheckType(typeOfSlice, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		if err := CheckType(typeOfElement, in[0]); err != nil {
			t.Error("function CheckType has bug")
		}
		t.Log("check type of slice is success")
	}
}

func TestCheckType_With_Slice_And_Input_Other_Type(t *testing.T) {
	typeOf := "slice"
	in := "slice"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Slice_And_Input_Type_Different_Of_Element(t *testing.T) {
	typeOfSlice, typeOfElement := "slice", "string"
	in := []int{1, 2, 3}
	if err := CheckType(typeOfSlice, in); err != nil {
		t.Errorf("unknow error! error: %v ", err)
	} else {
		if err := CheckType(typeOfElement, in[0]); err != nil {
			if errors.Is(err, errors2.InvalidTypeError) {
				t.Log("slice element type not same")
				return
			} else {
				t.Errorf("unknow error! error: %v ", err)
				return
			}
		}
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Map(t *testing.T) {
	typeOfSlice, typeOfElement := "map", "string"
	key, val := "hello", "world"
	in := map[string]string{key: val}
	if err := CheckType(typeOfSlice, in); err != nil {
		t.Error("function CheckType has bug")
	} else {
		if err := CheckType(typeOfElement, key); err != nil {
			t.Error("function CheckType has bug")
			return
		}
		if err := CheckType(typeOfElement, val); err != nil {
			t.Error("function CheckType has bug")
			return
		}
		t.Log("check type of map is success")
	}
}

func TestCheckType_With_Map_And_Input_Other_Type(t *testing.T) {
	typeOf := "map"
	in := "map"
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Map_And_Input_Type_Different_Of_Element(t *testing.T) {
	typeOfMap, typeOfElement := "map", "string"
	key, val := "key", 1
	in := map[string]int{key: val}
	if err := CheckType(typeOfMap, in); err != nil {
		t.Errorf("unknow error! error: %v ", err)
	} else {
		if err := CheckType(typeOfElement, key); err != nil {
			t.Errorf("unknow error! error: %v ", err)
			return
		}
		if err := CheckType(typeOfElement, val); err != nil {
			if errors.Is(err, errors2.InvalidTypeError) {
				t.Log("value type not same")
				return
			} else {
				t.Errorf("unknow error! error: %v ", err)
				return
			}
		}
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Struct(t *testing.T) {
	typeOf := "generic.TestDemo1"
	in := TestDemo1{Name: "test"}
	if err := CheckType(typeOf, in); err != nil {
		t.Errorf("function CheckType has error! error: %#v ", err.Error())
	} else {
		t.Log("check type of struct is success")
	}
}

func TestCheckType_With_Struct_And_Input_Other_Type(t *testing.T) {
	typeOf := "generic.TestDemo1"
	in := TestDemo2{Name: "test"}
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("struct type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

func TestCheckType_With_Ptr(t *testing.T) {
	typeOf := "*generic.TestDemo1"
	in := &TestDemo1{Name: "test"}
	if err := CheckType(typeOf, in); err != nil {
		t.Errorf("function CheckType has error! error: %#v ", err.Error())
	} else {
		t.Log("check type of struct is success")
	}
}

func TestCheckType_With_Ptr_And_Input_Other_Type(t *testing.T) {
	typeOf := "*generic.TestDemo1"
	in := &TestDemo2{Name: "test"}
	if err := CheckType(typeOf, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Log("struct type not same")
		} else {
			t.Errorf("unknow error! error: %v ", err)
		}
	} else {
		t.Error("function CheckType has bug")
	}
}

// TestDemo1 .
type TestDemo1 struct {
	Name string
}

// TestDemo2 .
type TestDemo2 struct {
	Name string
}
