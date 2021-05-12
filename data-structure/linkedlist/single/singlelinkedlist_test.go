package single

//
//import (
//	"testing"
//)
//
//func TestLinkedList_InsertToHead(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 10; i++ {
//		linkedList.InsertToHead(i)
//	}
//	node := linkedList.FindByIndex(0)
//	t.Log(node.Value())
//	if node.Value() == 9 {
//		t.Log("成功从链表头部插入！")
//	} else {
//		t.Error("从链表头部插入失败！！！！")
//	}
//}
//
//func TestLinkedList_InsertToTail(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 10; i++ {
//		linkedList.Add(i)
//	}
//	node := linkedList.FindByIndex(linkedList.Len() - 1)
//	t.Log(node.Value())
//	if node.Value() == 9 {
//		t.Log("成功从链表尾部插入！")
//	} else {
//		t.Error("从链表尾部插入失败！！！！")
//	}
//}
//
//func TestLinkedList_InsertAfter(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.FindByIndex(2)
//	flag := linkedList.InsertAfter(node, 100)
//	if flag {
//		t.Logf("after: ")
//		linkedList.Print()
//		t.Log("某个节点后插入成功")
//	} else {
//		t.Error("某个节点后插入失败")
//	}
//}
//
//func TestLinkedList_InsertBefore(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.FindByIndex(2)
//	flag := linkedList.InsertBefore(node, 100)
//	if flag {
//		t.Logf("after: ")
//		linkedList.Print()
//		t.Log("某个节点前插入成功")
//	} else {
//		t.Error("某个节点前插入失败")
//	}
//}
//
//func TestLinkedList_DeleteNode(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.FindByIndex(2)
//	flag := linkedList.DeleteNode(node)
//	if flag {
//		t.Logf("after: ")
//		linkedList.Print()
//		t.Log("节点删除成功")
//	} else {
//		t.Error("节点删除失败")
//	}
//}
//
//func TestLinkedList_Reverse(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	linkedList.Reverse()
//	if linkedList.FindByIndex(4).Value() == 0 {
//		t.Log("单链表反转成功")
//		t.Logf("after: ")
//		linkedList.Print()
//	} else {
//		t.Error("单链表反转失败")
//	}
//}
//
////func TestLinkedList_HasCycle(t *testing.T) {
////	arr := []string{"a", "b", "c", "d", "b", "c", "d"}
////	linkedList := linkedlist.NewLinkedList()
////	for _, a := range arr {
////		linkedList.Add(a)
////	}
////	linkedList.Head().Next().Next().Next().Next() = linkedList.Head().Next()
////	hasCycle := linkedList.HasCycle()
////	if hasCycle {
////		t.Log("链表有环")
////	} else {
////		t.Error("链表没有环")
////	}
////}
//
//func TestLinkedList_DeleteReciprocal(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.DeleteReciprocal(3)
//	if nil != node {
//		t.Log("删除倒数第三个元素成功,该元素是：", node.Value())
//		t.Logf("after: ")
//		linkedList.Print()
//	} else {
//		t.Error("删除倒数第二个元素成功")
//	}
//}
//
//func TestLinkedList_GetMiddle(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	middle := linkedList.Middle()
//	if nil != middle && middle.Value() == 2 {
//		t.Log("获取链表中间节点成功")
//	} else {
//		t.Error("获取链表中间节点失败")
//	}
//}
