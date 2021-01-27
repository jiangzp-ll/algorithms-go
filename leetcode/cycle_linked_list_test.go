package leetcode

import (
	"github.com/zepeng-jiang/go-basic-demo/data-structure/linkedlist"
	"testing"
)

// hasCycle 判断链表是否有环
// 快慢指针法
func hasCycle(head *linkedlist.ListNode) bool {
	if nil == head || nil == head.Next {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func TestHasCycle(t *testing.T) {
	list := linkedlist.NewLinkedList()
	list.InsertToTail(3)
	list.InsertToTail(4)
	list.InsertToTail(1)
	list.InsertToTail(2)
	list.Head.Next.Next.Next = list.Head.Next

	hasCycle := hasCycle(list.GetHead())
	if hasCycle {
		t.Log("有环")
	} else {
		t.Error("没环")
	}
}
