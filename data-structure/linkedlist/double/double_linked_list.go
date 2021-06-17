package double

import (
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"github.com/zepeng-jiang/go-basic-demo/pkg/generic"
)

// Node ,LinkedList element
type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

// LinkedList ,double LinkedList
type LinkedList struct {
	// head ,head Node
	head *Node
	// tail ,tail Node
	tail *Node
	// len ,number of elements in LinkedList
	len int
	// typeOf ,LinkedList type
	// Because Go not have Generic
	typeOf string
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
func NewLinkedList(typeOf string) (*LinkedList, error) {
	if typeOf == "" {
		return nil, errors2.InvalidTypeError
	}
	return &LinkedList{
		head:   NewNode(nil),
		tail:   NewNode(nil),
		len:    0,
		typeOf: typeOf,
	}, nil
}

// Add , add a value to Double LinkedList tail
func (l *LinkedList) Add(val interface{}) error {
	if err := l.Check(val); err != nil {
		return err
	}
	node := NewNode(val)
	if l.IsEmpty() {
		l.head = node
		l.tail = node
		l.len++
		return nil
	}
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	cur.next, node.prev = node, cur
	l.tail = node
	l.len++
	return nil
}

// AddToHead ,add a value to Double LinkedList head
func (l *LinkedList) AddToHead(val interface{}) error {
	if err := l.Check(val); err != nil {
		return err
	}
	node := NewNode(val)
	if l.IsEmpty() {
		l.head = node
		l.tail = node
		l.len++
		return nil
	}
	oldHead := l.head
	node.next = oldHead
	l.head = node
	l.len++
	return nil
}

// AllIndexesOf ,Returns all indexes of the specified element in the LinkedList
// Using two pointers to speed up traversal of Double LinkedList
func (l *LinkedList) AllIndexesOf(val interface{}) ([]int, error) {
	indexes := make([]int, 0)
	if err := l.Check(val); err != nil {
		return indexes, err
	}
	if l.IsEmpty() {
		return indexes, errors2.LinkedListIsEmptyError
	}
	if l.tail == l.head && l.len == 1 {
		if val == l.head.value {
			indexes = append(indexes, 1)
			return indexes, nil
		} else {
			return nil, errors2.ValueNotExistError
		}
	}
	h, t := l.head, l.tail
	for {
		mid := l.len >> 1
		for i := 1; i <= mid; i++ {
			if h.value == val {
				indexes = append(indexes, i)
			}
			h = h.next
		}
		for i := l.len; i > mid; i-- {
			if t.value == val {
				indexes = append(indexes, i)
			}
			t = t.prev
		}
		if t == h || t.next == h {
			break
		}
	}
	return indexes, nil
}

// Check ,check input value type
func (l *LinkedList) Check(val interface{}) error {
	if nil == val {
		return errors2.InputValueCannotBeNilError
	}
	if err := generic.CheckType(l.typeOf, val); err != nil {
		return err
	}
	return nil
}

// checkElementIndex ,check whether index is in LinkedList
func (l *LinkedList) checkElementIndex(index int) bool {
	return index > 0 && index <= l.len
}

// checkNode , check whether the node and LinkedList is valid
func (l *LinkedList) checkNodeAndLinkedList(n *Node) error {
	if nil == n {
		return errors2.InputNodeIsEmptyError
	}
	if l.IsEmpty() {
		return errors2.LinkedListIsEmptyError
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
		return nil, errors2.InvalidIndexError
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

// IndexOf ,get the index where the element first appears in the Double LinkedList
// Why return index when there is an error?
//All returned indexes are invalid. Avoid not checking error, but insist on using the returned value
func (l *LinkedList) IndexOf(val interface{}) (int, error) {
	if err := l.Check(val); err != nil {
		return -1, err
	}
	if l.IsEmpty() {
		return -1, errors2.LinkedListIsEmptyError
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			return i, nil
		}
		cur = cur.next
	}
	return 0, errors2.ValueNotExistError
}

// InsertAfter ,insert a node after the specified node
func (l *LinkedList) InsertAfter(n *Node, val interface{}) error {
	if err := l.Check(val); err != nil {
		return err
	}
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
	return errors2.NotExistError
}

// InsertBefore ,insert a node before the specified node
func (l *LinkedList) InsertBefore(n *Node, val interface{}) error {
	if err := l.Check(val); err != nil {
		return err
	}
	if err := l.checkNodeAndLinkedList(n); err != nil {
		return err
	}
	if l.head == n {
		_ = l.AddToHead(val)
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
	return errors2.NotExistError
}

// IsEmpty ,determine whether the Double LinkedList is empty
func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

// LastIndexOf ,returns the index of the last occurrence of the specified element in the Double LinkedList
func (l *LinkedList) LastIndexOf(val interface{}) (int, error) {
	if err := l.Check(val); err != nil {
		return -1, err
	}
	if l.IsEmpty() {
		return -1, errors2.LinkedListIsEmptyError
	}
	cur := l.tail
	for i := l.len; i > 0; i-- {
		if cur.value == val {
			return i, nil
		}
		cur = cur.prev
	}
	return 0, errors2.ValueNotExistError
}

// Len ,get the number of elements in the Double LinkedList
func (l *LinkedList) Len() int {
	return l.len
}

// Remove ,remove the specified node of Double LinkedList
func (l *LinkedList) Remove(n *Node) error {
	if err := l.checkNodeAndLinkedList(n); err != nil {
		return err
	}
	if l.head == n {
		l.Clear()
		return nil
	}
	cur := l.head.next
	for nil != cur {
		if cur == n {
			break
		}
		cur = cur.next
	}
	if nil == cur {
		return errors2.NotExistError
	}
	pre, next := cur.prev, cur.next
	pre.next, next.prev = next, pre
	l.len--
	return nil
}

// RemoveOf ,remove the element of the specified index in the LinkedList
func (l *LinkedList) RemoveOf(index int) (*Node, error) {
	node, err := l.Get(index)
	if err != nil {
		return nil, err
	}
	_ = l.Remove(node)
	return node, nil
}

// Reverse ,reverse the Double LinkedList
func (l *LinkedList) Reverse() {
	cur := l.tail
	for nil != cur.prev {
		pre, next := cur.prev, cur.next
		cur.prev, cur.next = next, pre
		cur = pre
	}
	l.head, l.tail = l.tail, l.head
}

// Set ,set the value of the specified index in the LinkedList
func (l *LinkedList) Set(index int, val interface{}) (interface{}, error) {
	if err := l.Check(val); err != nil {
		return nil, err
	}
	if !l.checkElementIndex(index) {
		return nil, errors2.InvalidIndexError
	}
	n, _ := l.Get(index)
	oldVal := n.value
	n.value = val
	return oldVal, nil
}

// Tail ,get Double LinkedList tail
func (l *LinkedList) Tail() *Node {
	return l.tail
}
