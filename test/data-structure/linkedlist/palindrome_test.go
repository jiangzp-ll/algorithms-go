package linkedlist

import (
	"github.com/zepeng-jiang/go-basic-demo/main/data-structure/linkedlist"
	"testing"
)

func TestIsPalindromeTrue(t *testing.T) {
	str := "ssoss"
	list := linkedlist.NewLinkedList()
	for _, s := range str {
		list.InsertToTail(string(s))
	}
	result := linkedlist.IsPalindrome(list)
	if result {
		t.Log("满足预期，是回文字符串")
	} else {
		t.Error("不满足预期，不是回文字符串")
	}
}

func TestIsPalindromeFalse(t *testing.T) {
	str := "hello"
	list := linkedlist.NewLinkedList()
	for _, s := range str {
		list.InsertToTail(string(s))
	}
	result := linkedlist.IsPalindrome(list)
	if !result {
		t.Log("满足预期，不是回文字符串")
	} else {
		t.Error("不满足预期，不是回文字符串")
	}
}
