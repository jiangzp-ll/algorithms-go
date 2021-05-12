package single

import (
	"errors"
	"fmt"
)

var InsertValueMustBeNotEmptyError = errors.New("insert value must be not empty")

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
func NewNode(v interface{}) (*Node, error) {
	if nil == v {
		return nil, InsertValueMustBeNotEmptyError
	}
	return &Node{nil, v}, nil
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
	return &LinkedList{nil, 0}
}

// Add , add a value to Single LinkedList tail
func (l *LinkedList) Add(v interface{}) error {
	//cur := l.head
	//for nil != cur.next {
	//	cur = cur.next
	//}
	//return l.InsertAfter(cur, v)

	node, err := NewNode(v)
	if err != nil {
		return err
	}
	if nil == l.head {
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
func (l *LinkedList) AddToHead(v interface{}) error {
	node, err := NewNode(v)
	if err != nil {
		return err
	}
	if nil == l.head {
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
	if nil == val {
		return indexes, errors.New("value must be not nil")
	}
	if nil == l.head {
		return indexes, errors.New("LinkedList is empty")
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

// Clear , clear the Single LinkedList
func (l *LinkedList) Clear() {
	l.head = nil
	l.len = 0
}

// Contain ,determine whether the value contain in the Single LinkedList
func (l *LinkedList) Contain(val interface{}) (bool, error) {
	index, err := l.IndexOf(val)
	if err != nil {
		return false, err
	}
	return index >= 0, nil
}

// Get ,Get the index element in the Single LinkedList
func (l *LinkedList) Get(index int) (*Node, error) {
	if index < 0 || index > l.len {
		return nil, errors.New("input the index does not exist")
	}
	cur := l.head
	for i := 1; i <= index; i++ {
		cur = cur.next
	}
	return cur, nil
}

// HasCycle ,judge whether the Single LinkedList has cycle
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

// Head ,get Single LinkedList Head
func (l *LinkedList) Head() *Node {
	return l.head
}

// IndexOf ,get the position of the value in the Single LinkedList
func (l *LinkedList) IndexOf(val interface{}) (int, error) {
	if nil == val {
		return -1, errors.New("value must be not nil")
	}
	if nil == l.head {
		return -2, errors.New("LinkedList is empty")
	}
	cur := l.head
	for i := 1; i <= l.len; i++ {
		if cur.value == val {
			return i, nil
		}
		cur = cur.next
	}
	return 0, errors.New("the value does not exist")
}

// InsertAfter ,insert a node after the specified node
func (l *LinkedList) InsertAfter(n *Node, val interface{}) error {
	if nil == n {
		return errors.New("node is nil")
	}
	newNode, err := NewNode(val)
	if err != nil {
		return err
	}
	if nil == l.head {
		return l.AddToHead(val)
	}
	cur := l.head
	for nil != cur {
		next := cur.next
		if cur == n {
			n.next = newNode
			newNode.next = next
			return nil
		}
	}
	return errors.New("input node not in this LinkedList")
}

// InsertBefore ,insert a node before the specified node
func (l *LinkedList) InsertBefore(n *Node, val interface{}) error {
	if nil == n {
		return errors.New("node is nil")
	}
	if nil == l.head {
		return errors.New("LinkedList is empty")
	}
	if l.head == n {
		return l.AddToHead(val)
	}
	node, err := NewNode(val)
	if err != nil {
		return err
	}
	cur, pre := l.head, &Node{}
	for nil != cur.next {
		next := cur.next
		pre = cur
		cur = next
		if cur.value == val {
			pre.next = node
			node.next = cur
			return nil
		}
	}
	return errors.New("input node not in this LinkedList")
}

// Last , get LinkedList last node
func (l *LinkedList) Last() *Node {
	if nil == l.head {
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
		return 0, errors.New("the value does not exist")
	}
	return indexes[len(indexes)-1], nil
}

// Len ,get the number of elements in the Single LinkedList
func (l *LinkedList) Len() int {
	return l.len
}

// Middle ，get middle node of Single LinkedList
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
// TODO
func (l *LinkedList) Remove(n *Node) error {
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

// FindByIndex 通过索引查找结点
func (l *LinkedList) FindByIndex(index int) *Node {
	if index >= l.len {
		return nil
	}
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除结点
func (l *LinkedList) DeleteNode(p *Node) bool {
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
	l.len--
	return true
}

// Print 打印链表
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Value())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// DeleteReciprocal 删除倒数第几个节点
func (l *LinkedList) DeleteReciprocal(index int) *Node {
	if index < 0 || index > l.len || nil == l.head || nil == l.head.next {
		return nil
	}
	cur := l.head
	for i := 0; i <= l.len-index-1; i++ {
		cur = cur.next
	}
	dNode := cur.next
	newNext := cur.next.next
	cur.next = newNext
	return dNode
}
