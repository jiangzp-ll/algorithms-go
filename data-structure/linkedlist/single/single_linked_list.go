package single

import (
	"errors"
)

var (
	LinkedListIsEmptyError = errors.New("LinkedList is empty")
	InputNodeIsEmptyError  = errors.New("input node is empty")
	NodeNotExistError      = errors.New("node not exist in this LinkedList")
	InvalidIndexError      = errors.New("invalid index")
	ValueNotExistError     = errors.New("value exist in the LinkedList")
)

// Node ,LinkedList element
type Node struct {
	next  *Node
	value interface{}
}

// LinkedList ,single LinkedList
type LinkedList struct {
	head *Node
	len  int
}

// NewNode ，init Node
func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

// Next ,get next Node
func (n *Node) Next() *Node {
	return n.next
}

// Value ,get the value stored in Node
func (n *Node) Value() interface{} {
	return n.value
}

// NewLinkedList ,init a empty Single LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(nil), 0}
}

// Add , add a value to Single LinkedList tail
func (l *LinkedList) Add(v interface{}) {
	node := NewNode(v)
	if nil == l.head.value {
		l.head = node
		l.len++
		return
	}
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	cur.next = node
	l.len++
	return
}

// AddToHead ,add a value to Single LinkedList head
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
func (l *LinkedList) AllIndexesOf(val interface{}) ([]int, error) {
	indexes := make([]int, 0)
	if nil == l.head.value {
		return indexes, LinkedListIsEmptyError
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			indexes = append(indexes, i)
		}
		cur = cur.next
	}
	return indexes, nil
}

// checkElementIndex ,check whether index is in LinkedList
func (l *LinkedList) checkElementIndex(index int) bool {
	return index > 0 && index <= l.len
}

// Clear , clear the Single LinkedList
func (l *LinkedList) Clear() {
	l.head = NewNode(nil)
	l.len = 0
}

// Contain ,determine whether the value contain in the Single LinkedList
func (l *LinkedList) Contain(val interface{}) bool {
	_, err := l.IndexOf(val)
	if err != nil {
		return false
	}
	return true
}

// Get ,Get the index element in the Single LinkedList
func (l *LinkedList) Get(index int) (*Node, error) {
	if !l.checkElementIndex(index) {
		return nil, InvalidIndexError
	}
	cur := l.head
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	return cur, nil
}

// HasCycle ,judge whether the Single LinkedList has cycle
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

// Head ,get Single LinkedList Head
func (l *LinkedList) Head() *Node {
	return l.head
}

// IndexOf ,get the position of the value in the Single LinkedList
func (l *LinkedList) IndexOf(val interface{}) (int, error) {
	if nil == l.head.value {
		return -1, LinkedListIsEmptyError
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			return i, nil
		}
		cur = cur.next
	}
	return 0, ValueNotExistError
}

// InsertAfter ,insert a node after the specified node
func (l *LinkedList) InsertAfter(n *Node, val interface{}) error {
	if nil == n {
		return InputNodeIsEmptyError
	}
	if nil == l.head.value {
		return LinkedListIsEmptyError
	}
	newNode := NewNode(val)
	cur := l.head
	for nil != cur {
		next := cur.next
		if cur.value == n.value {
			cur.next = newNode
			newNode.next = next
			l.len++
			return nil
		}
		cur = next
	}
	return NodeNotExistError
}

// InsertBefore ,insert a node before the specified node
func (l *LinkedList) InsertBefore(n *Node, val interface{}) error {
	if nil == n {
		return InputNodeIsEmptyError
	}
	if nil == l.head.value {
		return LinkedListIsEmptyError
	}
	if l.head == n {
		l.AddToHead(val)
		return nil
	}
	node := NewNode(val)
	cur, pre := l.head, &Node{}
	for nil != cur.next {
		next := cur.next
		pre = cur
		if next.value == n.value {
			pre.next, node.next = node, next
			l.len++
			return nil
		}
		cur = next
	}
	return NodeNotExistError
}

// Last , get LinkedList last node
func (l *LinkedList) Last() *Node {
	if nil == l.head.value {
		return nil
	}
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	return cur
}

// LastIndexOf ,returns the index of the last occurrence of the specified element in the LinkedList
func (l *LinkedList) LastIndexOf(val interface{}) (int, error) {
	indexes, err := l.AllIndexesOf(val)
	if err != nil {
		return -1, err
	}
	if len(indexes) == 0 {
		return 0, ValueNotExistError
	}
	return indexes[len(indexes)-1], nil
}

// Len ,get the number of elements in the Single LinkedList
func (l *LinkedList) Len() int {
	return l.len
}

// Middle ，get middle node of Single LinkedList
// TODO 重构这个方法
func (l *LinkedList) Middle() *Node {
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

// Remove ,remove the specified node of Single LinkedList
func (l *LinkedList) Remove(n *Node) error {
	if nil == n {
		return InputNodeIsEmptyError
	}
	cur, pre := l.head.next, l.head
	for nil != cur {
		if cur == n {
			break
		}
		pre, cur = cur, cur.next
	}
	if nil == cur {
		return NodeNotExistError
	}
	pre.next = n.next
	n = nil
	l.len--
	return nil
}

// Reverse ,reverse the Single LinkedList
func (l *LinkedList) Reverse() {
	if nil == l.head || nil == l.head.next {
		return
	}
	var pre *Node
	cur := l.head.next
	for nil != cur {
		next := cur.next
		cur.next = pre
		pre, cur = cur, next
	}
	l.head.next = pre
}

// Set ,set the value of the specified index in the LinkedList
func (l *LinkedList) Set(index int, val interface{}) (interface{}, error) {
	if !l.checkElementIndex(index) {
		return nil, InvalidIndexError
	}
	n, err := l.Get(index)
	if err != nil {
		return nil, err
	}
	oldVal := n.value
	n.value = val
	return oldVal, nil
}

//// DeleteReciprocal 删除倒数第几个节点
//func (l *LinkedList) DeleteReciprocal(index int) *Node {
//	if index < 0 || index > l.len || nil == l.head || nil == l.head.next {
//		return nil
//	}
//	cur := l.head
//	for i := 0; i <= l.len-index-1; i++ {
//		cur = cur.next
//	}
//	dNode := cur.next
//	newNext := cur.next.next
//	cur.next = newNext
//	return dNode
//}
