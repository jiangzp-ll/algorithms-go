package main

import (
	"testing"
)

func TestFact_Factorial(t *testing.T) {
	f := NewFactorial(5)
	ret := f.Factorial(5)
	if ret == 120 && f.Print(5) == 120 {
		t.Log("阶乘求解成功")
	} else {
		t.Error("阶乘求解失败")
	}
}