package linkedlist

import (
	"testing"
)

func TestIsPalindromeTrue(t *testing.T) {
	str := "ssoss"
	list := NewLinkedList()
	for _, s := range str {
		list.InsertToTail(string(s))
	}
	result := IsPalindrome(list)
	if result {
		t.Log("满足预期，是回文字符串")
	} else {
		t.Error("不满足预期，不是回文字符串")
	}
}

func TestIsPalindromeFalse(t *testing.T) {
	str := "hello"
	list := NewLinkedList()
	for _, s := range str {
		list.InsertToTail(string(s))
	}
	result := IsPalindrome(list)
	if !result {
		t.Log("满足预期，不是回文字符串")
	} else {
		t.Error("不满足预期，不是回文字符串")
	}
}
