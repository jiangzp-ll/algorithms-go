package recursion

import (
	"github.com/zepeng-jiang/go-basic-demo/algorithm/recursion"
	"testing"
)

func TestFact_Factorial(t *testing.T) {
	f := recursion.NewFactorial(5)
	ret := f.Factorial(5)
	if ret == 120 && f.Print(5) == 120 {
		t.Log("阶乘求解成功")
	} else {
		t.Error("阶乘求解失败")
	}
}