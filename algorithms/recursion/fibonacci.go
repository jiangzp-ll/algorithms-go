package recursion

import (
	"fmt"
	"github.com/zepeng-jiang/go-basic-demo/algorithms/sort"
)

// 递归实现斐波那契数列

// Fib .
type Fib struct {
	// 使用map存储结果
	value map[int]int
}

// NewFib 初始化
func NewFib(n int) *Fib {
	return &Fib{value: make(map[int]int, n)}
}

// Fibonacci 斐波那契数列
func (f *Fib) Fibonacci(n int) int {
	if f.value[n] != 0 {
		return f.value[n]
	}
	if n <= 1 {
		f.value[1] = 1
		return 1
	} else if n == 2 {
		f.value[2] = 1
		return 1
	} else {
		val := f.Fibonacci(n-1) + f.Fibonacci(n-2)
		f.value[n] = val
		return val
	}
}

// Print 打印结果
func (f *Fib) Print() string {
	var keys = []int{}
	for k, _ := range f.value {
		keys = append(keys, k)
	}
	newKeys := sort.QuickSort(keys, 0, len(keys)-1)
	result := "["
	for i := 0; i < len(newKeys)-1; i++ {
		result += fmt.Sprintf("%d,", f.value[newKeys[i]])
	}
	result += fmt.Sprintf("%d]", f.value[newKeys[len(newKeys)-1]])
	return result
}
