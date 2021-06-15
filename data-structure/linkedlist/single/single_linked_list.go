package single

import (
	"github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"reflect"
)

// Node ,LinkedList element
type Node struct {
	next  *Node
	value interface{}
}

// LinkedList ,single LinkedList
type LinkedList struct {
	// head ,head Node
	head *Node
	// len ,,number of elements in LinkedList
	len int
	// typeOf , LinkedList type
	// Because Go not have Generic
	typeOf string
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
func NewLinkedList(typeOf string) (*LinkedList, error) {
	if typeOf == "" {
		return nil, errors.InvalidTypeError
	}
	return &LinkedList{
		NewNode(nil),
		0,
		typeOf,
	}, nil
}

// Add , add a value to Single LinkedList tail
func (l *LinkedList) Add(val interface{}) error {
	if err := l.checkType(val); err != nil {
		return err
	}
	node := NewNode(val)
	if l.IsEmpty() {
		l.head = node
		l.len++
		return nil
	}
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	cur.next = node
	l.len++
	return nil
}

// AddToHead ,add a value to Single LinkedList head
func (l *LinkedList) AddToHead(val interface{}) error {
	if err := l.checkType(val); err != nil {
		return err
	}
	node := NewNode(val)
	if l.IsEmpty() {
		l.head = node
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
func (l *LinkedList) AllIndexesOf(val interface{}) ([]int, error) {
	indexes := make([]int, 0)
	if err := l.checkType(val); err != nil {
		return indexes, err
	}
	if l.IsEmpty() {
		return indexes, errors.LinkedListIsEmptyError
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

// checkNode , check whether the node and LinkedList is valid
func (l *LinkedList) checkNodeAndLinkedList(n *Node) error {
	if nil == n {
		return errors.InputNodeIsEmptyError
	}
	if l.IsEmpty() {
		return errors.LinkedListIsEmptyError
	}
	return nil
}

// checkType ,check input type
func (l *LinkedList) checkType(val interface{}) error {
	t := reflect.TypeOf(val).Kind().String()
	if l.typeOf != t {
		return errors.InvalidTypeError
	}
	return nil
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
		return nil, errors.InvalidIndexError
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
// Why return index when there is an error?
//All returned indexes are invalid. Avoid not checking error, but insist on using the returned value
func (l *LinkedList) IndexOf(val interface{}) (int, error) {
	if err := l.checkType(val); err != nil {
		return -1, err
	}
	if l.IsEmpty() {
		return -1, errors.LinkedListIsEmptyError
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			return i, nil
		}
		cur = cur.next
	}
	return 0, errors.ValueNotExistError
}

// InsertAfter ,insert a node after the specified node
func (l *LinkedList) InsertAfter(n *Node, val interface{}) error {
	if err := l.checkType(val); err != nil {
		return err
	}
	if err := l.checkNodeAndLinkedList(n); err != nil {
		return err
	}
	NewNode := NewNode(val)
	cur := l.head
	for nil != cur {
		next := cur.next
		if cur.value == n.value {
			cur.next = NewNode
			NewNode.next = next
			l.len++
			return nil
		}
		cur = next
	}
	return errors.NotExistError
}

// InsertBefore ,insert a node before the specified node
func (l *LinkedList) InsertBefore(n *Node, val interface{}) error {
	if err := l.checkType(val); err != nil {
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
	return errors.NotExistError
}

// IsEmpty ,determine whether the Double LinkedList is empty
func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

// Last , get LinkedList last node
func (l *LinkedList) Last() *Node {
	if l.IsEmpty() {
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
		return 0, errors.ValueNotExistError
	}
	return indexes[len(indexes)-1], nil
}

// Len ,get the number of elements in the Single LinkedList
func (l *LinkedList) Len() int {
	return l.len
}

// Remove ,remove the specified node of Single LinkedList
func (l *LinkedList) Remove(n *Node) error {
	if err := l.checkNodeAndLinkedList(n); err != nil {
		return err
	}
	if l.head == n {
		l.Clear()
		return nil
	}
	cur, pre := l.head.next, l.head
	for nil != cur {
		if cur == n {
			break
		}
		pre, cur = cur, cur.next
	}
	if nil == cur {
		return errors.NotExistError
	}
	pre.next = n.next
	n = nil
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

// Reverse ,reverse the Single LinkedList
func (l *LinkedList) Reverse() {
	var pre *Node
	cur := l.head
	for nil != cur {
		next := cur.next
		cur.next = pre
		pre, cur = cur, next
	}
	l.head = pre
}

// Set ,set the value of the specified index in the LinkedList
func (l *LinkedList) Set(index int, val interface{}) (interface{}, error) {
	if err := l.checkType(val); err != nil {
		return nil, err
	}
	if !l.checkElementIndex(index) {
		return nil, errors.InvalidIndexError
	}
	n, _ := l.Get(index)
	oldVal := n.value
	n.value = val
	return oldVal, nil
}
