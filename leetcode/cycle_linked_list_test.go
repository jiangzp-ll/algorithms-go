package leetcode

import (
	"github.com/zepeng-jiang/go-basic-demo/data-structure/linkedlist"
	"testing"
)

// hasCycle 判断链表是否有环
// 快慢指针法
func hasCycle(head *linkedlist.ListNode) bool {
	if head.GetNext() == nil || head.GetNext().GetNext() == nil {
		return false
	}
	fast, slow := head.GetNext(), head
	for nil != head  {
		fast = fast.GetNext().GetNext()
		slow = slow.GetNext()
		if fast == slow {
			return true
		}
	}
	return false
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