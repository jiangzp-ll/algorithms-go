package hash

import "testing"

func TestHash(t *testing.T) {
	length, key := 10, "hello"
	index := Hash(length, key)
	if index < 10 && index >= 0 {
		t.Logf("hash: %d \n", index)
	} else {
		t.Error("function Hash has bug")
	}
}

func TestHash_With_Key_Is_Empty(t *testing.T) {
	length, key := 10, ""
	index := Hash(length, key)
	if index == 0 {
		t.Logf("hash: %d \n", index)
	} else {
		t.Error("function Hash has bug")
	}
}

func TestHash_With_Multiple_Generation(t *testing.T) {
	length, key := 10, "0"
	first := Hash(length, key)
	second := Hash(length, key)
	third := Hash(length, key)
	if first == second && first == third {
		t.Logf("hash: %d \n", first)
	} else {
		t.Error("function Hash has bug")
	}
}
