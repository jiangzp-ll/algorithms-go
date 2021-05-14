package double

import (
	"github.com/zepeng-jiang/go-basic-demo/data-structure/linkedlist/single"
)

// 双向链表

// Node ,LinkedList element
type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

// LinkedList ,double LinkedList
type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

// NewNode ,init Node
func NewNode(val interface{}) *Node {
	return &Node{nil, nil, val}
}

// Prev ,get prev Node
func (n *Node) Prev() *Node {
	return n.prev
}

// Next ,get next Node
func (n *Node) Next() *Node {
	return n.next
}

// Value ,get the value stored in Node
func (n *Node) Value() interface{} {
	return n.value
}

// NewLinkedList ,init a empty Double LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: NewNode(nil),
		tail: NewNode(nil),
		len:  0,
	}
}

// Add , add a value to Double LinkedList tail
func (l *LinkedList) Add(v interface{}) {
	node := NewNode(v)
	if nil == l.head.value {
		l.head = node
		l.tail = node
		l.len++
		return
	}
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	cur.next = node
	l.tail = node
	l.len++
	return
}

// AddToHead ,add a value to Double LinkedList head
func (l *LinkedList) AddToHead(v interface{}) {
	node := NewNode(v)
	if nil == l.head.value {
		l.head = node
		l.len++
		return
	}
	oldHead := l.head
	node.next = oldHead
	l.head = node
	l.len++
	return
}

// AllIndexesOf ,Returns all indexes of the specified element in the LinkedList
// Using two pointers to speed up traversal of Double LinkedList
func (l *LinkedList) AllIndexesOf(val interface{}) ([]int, error) {
	indexes := make([]int, 0)
	if nil == l.head.value {
		return indexes, single.LinkedListIsEmptyError
	}
	h, t := l.head, l.tail
	for {
		mid := l.len >> 1
		for i := l.len; i >= mid; i-- {
			if t.value == val {
				indexes = append(indexes, i)
			}
			t = t.prev
		}
		for i := 1; i <= mid; i++ {
			if h.value == val {
				indexes = append(indexes, i)
			}
			h = h.next
		}
		if t == h {
			break
		}
	}
	return indexes, nil
}

// checkElementIndex ,check whether index is in LinkedList
func (l *LinkedList) checkElementIndex(index int) bool {
	return index > 0 && index <= l.len
}

// checkNode , check whether the node and LinkedList is valid
func (l *LinkedList) checkNodeAndLinkedList(n *Node) error {
	if nil == n {
		return single.InputNodeIsEmptyError
	}
	if 0 == l.len {
		return single.LinkedListIsEmptyError
	}
	return nil
}

// Clear , clear the Double LinkedList
func (l *LinkedList) Clear() {
	l.head = NewNode(nil)
	l.tail = l.head
	l.len = 0
}

// Contain ,determine whether the value contain in the Double LinkedList
func (l *LinkedList) Contain(val interface{}) bool {
	_, err := l.IndexOf(val)
	if err != nil {
		return false
	}
	return true
}

// Get ,Get the index element in the Double LinkedList
func (l *LinkedList) Get(index int) (*Node, error) {
	if !l.checkElementIndex(index) {
		return nil, single.InvalidIndexError
	}
	cur := l.head
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	return cur, nil
}

// HasCycle ,judge whether the Double LinkedList has cycle
func (l *LinkedList) HasCycle() bool {
	if l.head == nil || l.head.next == nil {
		return false
	}
	slow, fast := l.head, l.head.next
	for fast != slow {
		if fast == nil || fast.next == nil {
			return false
		}
		slow, fast = slow.next, fast.next.next
	}
	return true
}

// Head ,get Double LinkedList Head
func (l *LinkedList) Head() *Node {
	return l.head
}

// IndexOf ,get the position of the value in the Single LinkedList
// Why return index when there is an error?
//All returned indexes are invalid. Avoid not checking error, but insist on using the returned value
func (l *LinkedList) IndexOf(val interface{}) (int, error) {
	if nil == l.head.value {
		return -1, single.LinkedListIsEmptyError
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			return i, nil
		}
		cur = cur.next
	}
	return 0, single.ValueNotExistError
}

// InsertAfter ,insert a node after the specified node
func (l *LinkedList) InsertAfter(n *Node, val interface{}) error {
	if err := l.checkNodeAndLinkedList(n); err != nil {
		return err
	}
	newNode := NewNode(val)
	cur := l.head
	for nil != cur {
		next := cur.next
		if cur.value == n.value {
			cur.next, newNode.next = newNode, next
			newNode.prev, next.prev = cur, newNode
			l.len++
			return nil
		}
		cur = next
	}
	return single.NodeNotExistError
}

// InsertBefore ,insert a node before the specified node
func (l *LinkedList) InsertBefore(n *Node, val interface{}) error {
	if err := l.checkNodeAndLinkedList(n); err != nil {
		return err
	}
	if l.head == n {
		l.AddToHead(val)
		return nil
	}
	node := NewNode(val)
	cur := l.head
	for nil != cur.next {
		next := cur.next
		if next.value == n.value {
			cur.next, node.next = node, next
			node.prev, next.prev = cur, node
			l.len++
			return nil
		}
		cur = next
	}
	return single.NodeNotExistError
}

// LastIndexOf ,returns the index of the last occurrence of the specified element in the LinkedList
func (l *LinkedList) LastIndexOf(val interface{}) (int, error) {
	if l.len == 0 {
		return -1, single.LinkedListIsEmptyError
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			return i, nil
		}
		cur = cur.next
	}
	return 0, single.ValueNotExistError
}
