package main

import "fmt"

// 实现一组数据的全排列

// RangeType .
type RangeType struct {
	value []interface{}
}

// NewRangeArray 全排列
func NewRangeArray(n int) *RangeType {
	return &RangeType{
		make([]interface{}, n),
	}
}

// RangeAll 全排列
func (r *RangeType) RangeAll(start int) {
	length := len(r.value)
	if start == length-1 {
		// 如果已经是最后位置，直接将数组数据合并输出
		fmt.Println(r.value)
	}

	for i := start; i < length; i++ {
		// i = start 时输出自己
		// 如果i和start的值相同就没有必要交换
		if i == start || r.value[i] != r.value[start] {
			//交换当前这个与后面的位置
			r.value[i], r.value[start] = r.value[start], r.value[i]
			//递归处理索引+1
			r.RangeAll(start + 1)
			//换回来，因为是递归，如果不换回来会影响后面的操作，并且出现重复
			r.value[i], r.value[start] = r.value[start], r.value[i]
		}
	}
}