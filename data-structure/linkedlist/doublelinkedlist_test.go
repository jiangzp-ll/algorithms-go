package linkedlist

import (
	"testing"
)

func TestDoubleLinkedList_InsertToHead(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 110; i++ {
		doubleLinkedList.InsertToHead(i)
	}
	if doubleLinkedList.GetHead().GetValue() == 109 && doubleLinkedList.GetTail().GetValue() == 100 {
		t.Log("从链表头部插入成功")
	} else {
		t.Error("从链表头部插入失败")
	}
}

func TestDoubleLinkedList_InsertToTail(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 110; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	if doubleLinkedList.GetHead().GetValue() == 100 && doubleLinkedList.GetTail().GetValue() == 109 {
		t.Log("从链表尾部插入成功")
	} else {
		t.Error("从链表尾部插入失败")
	}
}

func TestDoubleLinkedList_FindByIndex(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 105; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	node := doubleLinkedList.FindByIndex(3)
	doubleLinkedList.Print()
	if node.GetValue() == 103 {
		t.Log("通过索引查找节点成功")
	} else {
		t.Error("通过索引查找节点失败")
	}
}

func TestDoubleLinkedList_InsertAfter(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 105; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	node := doubleLinkedList.FindByIndex(3)
	result := doubleLinkedList.InsertAfter(node, 0)
	doubleLinkedList.Print()
	if result && doubleLinkedList.FindByIndex(4).GetValue() == 0{
		t.Log("插入节点成功！")
	} else {
		t.Error("插入节点失败！")
	}
}

func TestDoubleLinkedList_InsertBefore(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 105; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	node := doubleLinkedList.FindByIndex(3)
	result := doubleLinkedList.InsertBefore(node, 0)
	doubleLinkedList.Print()
	if result && doubleLinkedList.FindByIndex(3).GetValue() == 0{
		t.Log("插入节点成功！")
	} else {
		t.Error("插入节点失败！")
	}
}

func TestDoubleLinkedList_DeleteNode(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 105; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	node := doubleLinkedList.FindByIndex(3)
	err := doubleLinkedList.DeleteNode(node)
	if err != nil {
		t.Error(err.Error())
	}
	newNode := doubleLinkedList.FindByIndex(3)
	if newNode.GetValue() == 104 {
		t.Log("删除节点成功")
	} else {
		t.Error("删除节点失败")
	}
}

func TestDoubleLinkedList_DeleteHead(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 105; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	err := doubleLinkedList.DeleteHead()
	if err != nil {
		t.Error(err.Error())
	}
	doubleLinkedList.Print()
	newHead := doubleLinkedList.GetHead()
	if newHead.GetValue() == 101 && doubleLinkedList.GetLength() == 4{
		t.Log("删除链表头节点成功")
	} else {
		t.Error("删除链表头节点失败")
	}
}

func TestDoubleLinkedList_DeleteTail(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 105; i++ {
		doubleLinkedList.InsertToTail(i)
	}
	err := doubleLinkedList.DeleteTail()
	if err != nil {
		t.Error(err.Error())
	}
	doubleLinkedList.Print()
	newHead := doubleLinkedList.GetTail()
	if newHead.GetValue() == 103 && doubleLinkedList.GetLength() == 4 {
		t.Log("删除链表头节点成功")
	} else {
		t.Error("删除链表头节点失败")
	}
}