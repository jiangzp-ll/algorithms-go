package linkedlist

import (
	"errors"
	"fmt"
)

// 双向链表

// DNode 结点
type DNode struct {
	prev  *DNode
	next  *DNode
	value interface{}
}

// DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	head   *DNode
	tail   *DNode
	length int
}

// NewDNode 初始化结点
func NewDNode(v interface{}) *DNode {
	return &DNode{nil, nil, v}
}

// GetValue 获取结点的值
func (d *DNode) GetValue() interface{} {
	return d.value
}

// GetPrev 获取前一个结点
func (d *DNode) GetPrev() *DNode {
	return d.prev
}

// GetNext 获取下一个结点
func (d *DNode) GetNext() *DNode {
	return d.next
}

// NewDoubleLinkedList 初始化双向链表
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head:   &DNode{},
		tail:   &DNode{},
		length: 0,
	}
}

// GetLength 获取链表长度
func (d *DoubleLinkedList) GetLength() int {
	return d.length
}

// GetHead 获取头结点
func (d *DoubleLinkedList) GetHead() *DNode {
	return d.head
}

// GetTail 获取头结点
func (d *DoubleLinkedList) GetTail() *DNode {
	return d.tail
}

// InsertAfter 在任意结点后插入数据
func (d *DoubleLinkedList) InsertAfter(p *DNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newDNode := NewDNode(v)
	oldDNode := p.next
	p.next = newDNode
	newDNode.prev = p
	newDNode.next = oldDNode
	oldDNode.prev = newDNode
	d.length++
	return true
}

// InsertBefore 在任意结点前插入数据
func (d *DoubleLinkedList) InsertBefore(p *DNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newDNode := NewDNode(v)
	oldDNode := p.prev
	oldDNode.next = newDNode
	newDNode.prev = oldDNode
	newDNode.next = p
	p.prev = newDNode
	d.length++
	return true
}

// InsertToHead 在双向链表头插入数据
func (d *DoubleLinkedList) InsertToHead(v interface{}) bool {
	node := NewDNode(v)
	oldHead := d.head
	if nil == oldHead.value { //如果是空链表
		d.head = node
		d.tail = node
		d.length++
	} else {
		node.next = oldHead
		oldHead.prev = node
		d.head = node
		// 更新尾结点
		cur := d.head.next
		for nil != cur {
			if nil == cur.next {
				break
			}
			cur = cur.next
		}
		d.tail = cur
		d.length++
	}
	return d.head.value == v
}

// InsetToTail 在双向列表尾部插入数据
func (d *DoubleLinkedList) InsertToTail(v interface{}) bool {
	node := NewDNode(v)
	oldHead := d.head
	if nil == oldHead.value { //如果是空链表
		d.head = node
		d.tail = node
		d.length++
	} else {
		cur := d.head
		for nil != cur.next {
			cur = cur.next
		}
		cur.next = node
		node.prev = cur
		d.tail = node
		d.length++
	}
	return d.tail.value == v
}

// FindByIndex 通过索引查找结点
func (d *DoubleLinkedList) FindByIndex(index int) *DNode {
	if index > d.length || index < 0 {
		return nil
	}
	cur := d.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteDNode 删除结点
func (d *DoubleLinkedList) DeleteNode(p *DNode) error {
	if 0 == d.length {
		return errors.New("链表为空，无法进行删除操作！")
	}
	if nil == p {
		return errors.New("删除的结点不能为空！")
	}
	if d.tail == p {
		return d.DeleteTail()
	}
	if d.head == p {
		return d.DeleteHead()
	}
	cur := d.head.next
	for nil != cur {
		if cur == p {
			break
		}
		cur = cur.next
	}
	pre := cur.prev
	next := cur.next
	pre.next = next
	next.prev = pre
	d.length--
	return nil
}

// DeleteTail 删除链表尾结点
func (d *DoubleLinkedList) DeleteTail() error {
	if d.length == 0 {
		return errors.New("空链表，无法进行删除尾结点操作！")
	}
	if d.length == 1 {
		d.head = &DNode{}
		d.tail = &DNode{}
		d.length = 0
		return nil
	}
	newTail := d.tail.prev
	newTail.next = nil
	d.tail = newTail
	d.length--
	return nil
}

// DeleteHead 删除链表头结点
func (d *DoubleLinkedList) DeleteHead() error {
	if d.length == 0 {
		return errors.New("空链表，无法进行删除头结点操作！")
	}
	if d.length == 1 {
		d.head = &DNode{}
		d.tail = &DNode{}
		d.length = 0
		return nil
	}
	newHead := d.head.next
	newHead.prev = nil
	d.head = newHead
	d.length--
	return nil
}

// Print 打印链表
func (d *DoubleLinkedList) Print() {
	cur := d.head
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
