package linkedlist

import "fmt"

// 单链表实现

// ListNode 结点
type ListNode struct {
	Next  *ListNode
	Value interface{}
}

// LinkedList 链表
type LinkedList struct {
	Head   *ListNode
	Length int
}

// NewListNode 初始化结点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// GetNext 获取下一个结点
func (l *ListNode) GetNext() *ListNode {
	return l.Next
}

// GetValue 获取结点的值
func (l *ListNode) GetValue() interface{} {
	return l.Value
}

// NewLinkedList 初始化链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// GetHead 获取链表头
func (l *LinkedList) GetHead() *ListNode {
	return l.Head
}

// GetLength 获取链表长度
func (l *LinkedList) GetLength() int {
	return l.Length
}

// InsertAfter 在某个结点后插入结点
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.Next
	p.Next = newNode
	newNode.Next = oldNext
	l.Length++
	return true
}

// InsertBefore 在某个结点前插入结点
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == l.Head {
		return false
	}
	cur := l.Head.Next
	pre := l.Head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur
	l.Length++
	return true
}

// InsertToHead 在链表头部插入结点
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.Head, v)
}

// InsertToTail 在链表尾部插入结点
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.Head
	for nil != cur.Next {
		cur = cur.Next
	}
	return l.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找结点
func (l *LinkedList) FindByIndex(index int) *ListNode {
	if index >= l.Length {
		return nil
	}
	cur := l.Head.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

// DeleteNode 删除结点
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := l.Head.Next
	pre := l.Head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if nil == cur {
		return false
	}
	pre.Next = p.Next
	p = nil
	l.Length--
	return true
}

// Print 打印链表
func (l *LinkedList) Print() {
	cur := l.Head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// Reverse 单链表反转
func (l *LinkedList) Reverse() {
	if nil == l.Head || nil == l.Head.Next {
		return
	}
	var pre *ListNode
	cur := l.Head.Next
	for nil != cur {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	l.Head.Next = pre
}

// HasCycle 判断链表是否有环
func (l *LinkedList) HasCycle() bool {
	if nil != l.Head {
		slow, fast := l.Head, l.Head
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// DeleteReciprocal 删除倒数第几个节点
func (l *LinkedList) DeleteReciprocal(index int) *ListNode {
	if index < 0 || index > l.Length || nil == l.Head || nil == l.Head.Next {
		return nil
	}
	cur := l.Head
	for i := 0; i <= l.Length-index-1; i++ {
		cur = cur.Next
	}
	dNode := cur.Next
	newNext := cur.Next.Next
	cur.Next = newNext
	return dNode
}

// GetMiddle 获取链表中间节点
func (l *LinkedList) GetMiddle() *ListNode {
	if nil == l.Head || nil == l.Head.Next {
		return nil
	}
	if nil == l.Head.Next.Next {
		return l.Head.Next
	}
	slow, fast := l.Head, l.Head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}