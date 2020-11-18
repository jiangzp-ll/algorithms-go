package main

// 求 n!

// Fact .
type Fact struct {
	value map[int]int
}

// NewFactorial 初始化
func NewFactorial(n int) *Fact {
	return &Fact{value: make(map[int]int, n)}
}

// Factorial 阶乘
func (f *Fact) Factorial(n int) int {
	if f.value[n] != 0 {
		return f.value[n]
	}

	if n <= 1 {
		f.value[n] = 1
		return 1
	} else {
		ret := n * f.Factorial(n-1)
		f.value[n] = ret
		return ret
	}
}

// Print 打印结果
func (f *Fact) Print(n int) int {
	return f.value[n]
}