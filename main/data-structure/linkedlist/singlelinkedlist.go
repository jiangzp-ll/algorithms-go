package linkedlist

import "fmt"

// 单链表实现

// ListNode 结点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// LinkedList 链表
type LinkedList struct {
	head   *ListNode
	length int
}

// NewListNode 初始化结点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// GetNext 获取下一个结点
func (l *ListNode) GetNext() *ListNode {
	return l.next
}

// GetValue 获取结点的值
func (l *ListNode) GetValue() interface{} {
	return l.value
}

// NewLinkedList 初始化链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// GetHead 获取链表头
func (l *LinkedList) GetHead() *ListNode {
	return l.head
}

// GetLength 获取链表长度
func (l *LinkedList) GetLength() int {
	return l.length
}

// InsertAfter 在某个结点后插入结点
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

// InsertBefore 在某个结点前插入结点
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == l.head {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++
	return true
}

// InsertToHead 在链表头部插入结点
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// InsertToTail 在链表尾部插入结点
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	return l.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找结点
func (l *LinkedList) FindByIndex(index int) *ListNode {
	if index >= l.length {
		return nil
	}
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除结点
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

// Print 打印链表
func (l *LinkedList) Print() {
	cur := l.head.next
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

// Reverse 单链表反转
func (l *LinkedList) Reverse() {
	if nil == l.head || nil == l.head.next {
		return
	}
	var pre *ListNode
	cur := l.head.next
	for nil != cur {
		next := cur.next
		cur.next = pre
		pre, cur = cur, next
	}
	l.head.next = pre
}

// HasCycle 判断链表是否有环
func (l *LinkedList) HasCycle() bool {
	if nil != l.head {
		slow, fast := l.head, l.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// DeleteReciprocal 删除倒数第几个节点
func (l *LinkedList) DeleteReciprocal(index int) *ListNode {
	if index < 0 || index > l.length || nil == l.head || nil == l.head.next {
		return nil
	}
	cur := l.head
	for i := 0; i <= l.length-index-1; i++ {
		cur = cur.next
	}
	dNode := cur.next
	newNext := cur.next.next
	cur.next = newNext
	return dNode
}

// GetMiddle 获取链表中间节点
func (l *LinkedList) GetMiddle() *ListNode {
	if nil == l.head || nil == l.head.next {
		return nil
	}
	if nil == l.head.next.next {
		return l.head.next
	}
	slow, fast := l.head, l.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}