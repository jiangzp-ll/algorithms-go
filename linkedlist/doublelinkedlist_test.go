package linkedlist

import (
	"testing"
)

func TestDoubleLinkedList_InsertToHead(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	for i := 100; i < 110; i++ {
		doubleLinkedList.InsertToHead(i)
	}
	if doubleLinkedList.GetHead().value == 109 && doubleLinkedList.GetTail().value == 100 {
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
	if doubleLinkedList.GetHead().value == 100 && doubleLinkedList.GetTail().value == 109 {
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
	if node.value == 103 {
		t.Log("通过索引查找结点成功")
	} else {
		t.Error("通过索引查找结点失败")
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
	if result && doubleLinkedList.FindByIndex(4).value == 0{
		t.Log("插入结点成功！")
	} else {
		t.Error("插入结点失败！")
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
	if result && doubleLinkedList.FindByIndex(3).value == 0{
		t.Log("插入结点成功！")
	} else {
		t.Error("插入结点失败！")
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
	if newNode.value == 104 {
		t.Log("删除结点成功")
	} else {
		t.Error("删除结点失败")
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
	newHead := doubleLinkedList.head
	if newHead.value == 101 && doubleLinkedList.length == 4{
		t.Log("删除链表头结点成功")
	} else {
		t.Error("删除链表头结点失败")
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
	newHead := doubleLinkedList.tail
	if newHead.value == 103 && doubleLinkedList.length == 4 {
		t.Log("删除链表头结点成功")
	} else {
		t.Error("删除链表头结点失败")
	}
}