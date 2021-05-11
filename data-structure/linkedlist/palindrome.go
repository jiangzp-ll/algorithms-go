package linkedlist

import "github.com/zepeng-jiang/go-basic-demo/data-structure/linkedlist/single"

// 单链表回文实现

//  开放一个栈存放链表前半段
// IsPalindrome 判断是否是回文字符串方法一
func IsPalindrome(l *single.LinkedList) bool {
	lLen := l.GetLength()
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	s := make([]string, 0, lLen>>1)
	cur := l.GetHead()
	for i := 1; i <= lLen; i++ {
		cur = cur.GetNext()
		//如果链表有奇数个节点，中间的直接忽略
		if lLen%2 != 0 && i == (lLen>>1+1) {
			continue
		}
		if i <= lLen>>1 {
			s = append(s, cur.GetValue().(string))
		} else {
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}
	return true
}
