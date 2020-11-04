package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_InsertToHead(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToHead(i)
	}
	node := linkedList.FindByIndex(0)
	fmt.Println(node.value)
	if node.value == 9 {
		t.Log("成功从链表头部插入！")
	} else {
		t.Error("从链表头部插入失败！！！！")
	}
}

func TestLinkedList_InsertToTail(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToTail(i)
	}
	node := linkedList.FindByIndex(linkedList.length - 1)
	fmt.Println(node.value)
	if node.value == 9 {
		t.Log("成功从链表尾部插入！")
	} else {
		t.Error("从链表尾部插入失败！！！！")
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 5; i++ {
		linkedList.InsertToTail(i)
	}
	fmt.Print("before: ")
	linkedList.Print()
	node := linkedList.FindByIndex(2)
	flag := linkedList.InsertAfter(node, 100)
	if flag {
		fmt.Print("after: ")
		linkedList.Print()
		t.Log("某个结点后插入成功")
	} else {
		t.Error("某个结点后插入失败")
	}
}

func TestLinkedList_InsertBefore(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 5; i++ {
		linkedList.InsertToTail(i)
	}
	fmt.Print("before: ")
	linkedList.Print()
	node := linkedList.FindByIndex(2)
	flag := linkedList.InsertBefore(node, 100)
	if flag {
		fmt.Print("after: ")
		linkedList.Print()
		t.Log("某个结点前插入成功")
	} else {
		t.Error("某个结点前插入失败")
	}
}

func TestLinkedList_DeleteNode(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 5; i++ {
		linkedList.InsertToTail(i)
	}
	fmt.Print("before: ")
	linkedList.Print()
	node := linkedList.FindByIndex(2)
	flag := linkedList.DeleteNode(node)
	if flag {
		fmt.Print("after: ")
		linkedList.Print()
		t.Log("结点删除成功")
	} else {
		t.Error("结点删除失败")
	}
}