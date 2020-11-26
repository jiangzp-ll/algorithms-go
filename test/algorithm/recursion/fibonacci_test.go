package recursion

import (
	"github.com/zepeng-jiang/go-basic-demo/main/algorithm/recursion"
	"testing"
)

func TestFibonacci(t *testing.T) {
	fib := recursion.NewFib(8)
	n := fib.Fibonacci(7)
	if n == 13 {
		t.Log("递归成功")
	} else {
		t.Error("递归失败")
	}
}

func TestFib_Print(t *testing.T) {
	fib := recursion.NewFib(8)
	n := fib.Fibonacci(7)
	result := fib.Print()
	except := "[1,1,2,3,5,8,13]"
	if result == except && n == 13 {
		t.Log("打印成功")
	} else {
		t.Error("打印失败")
	}
}