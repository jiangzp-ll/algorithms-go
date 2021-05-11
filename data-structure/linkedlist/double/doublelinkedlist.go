package double

import (
	"errors"
	"fmt"
)

// 双向链表

// Node 结点
type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

// DoubleLinkedList 双向链表
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// NewNode 初始化结点
func NewNode(v interface{}) *Node {
	return &Node{nil, nil, v}
}

// GetValue 获取结点的值
func (d *Node) GetValue() interface{} {
	return d.value
}

// GetPrev 获取前一个结点
func (d *Node) GetPrev() *Node {
	return d.prev
}

// GetNext 获取下一个结点
func (d *Node) GetNext() *Node {
	return d.next
}

// NewDoubleLinkedList 初始化双向链表
func NewDoubleLinkedList() *LinkedList {
	return &LinkedList{
		head:   &Node{},
		tail:   &Node{},
		length: 0,
	}
}

// GetLength 获取链表长度
func (d *LinkedList) GetLength() int {
	return d.length
}

// GetHead 获取头结点
func (d *LinkedList) GetHead() *Node {
	return d.head
}

// GetTail 获取头结点
func (d *LinkedList) GetTail() *Node {
	return d.tail
}

// InsertAfter 在任意结点后插入数据
func (d *LinkedList) InsertAfter(p *Node, v interface{}) bool {
	if nil == p {
		return false
	}
	newDNode := NewNode(v)
	oldDNode := p.next
	p.next = newDNode
	newDNode.prev = p
	newDNode.next = oldDNode
	oldDNode.prev = newDNode
	d.length++
	return true
}

// InsertBefore 在任意结点前插入数据
func (d *LinkedList) InsertBefore(p *Node, v interface{}) bool {
	if nil == p {
		return false
	}
	newDNode := NewNode(v)
	oldDNode := p.prev
	oldDNode.next = newDNode
	newDNode.prev = oldDNode
	newDNode.next = p
	p.prev = newDNode
	d.length++
	return true
}

// InsertToHead 在双向链表头插入数据
func (d *LinkedList) InsertToHead(v interface{}) bool {
	node := NewNode(v)
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
func (d *LinkedList) InsertToTail(v interface{}) bool {
	node := NewNode(v)
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
func (d *LinkedList) FindByIndex(index int) *Node {
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
func (d *LinkedList) DeleteNode(p *Node) error {
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
func (d *LinkedList) DeleteTail() error {
	if d.length == 0 {
		return errors.New("空链表，无法进行删除尾结点操作！")
	}
	if d.length == 1 {
		d.head = &Node{}
		d.tail = &Node{}
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
func (d *LinkedList) DeleteHead() error {
	if d.length == 0 {
		return errors.New("空链表，无法进行删除头结点操作！")
	}
	if d.length == 1 {
		d.head = &Node{}
		d.tail = &Node{}
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
func (d *LinkedList) Print() {
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
