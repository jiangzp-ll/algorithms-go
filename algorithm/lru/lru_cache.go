package lru

const (
	HOST_BIT = uint64(^uint(0)) == ^uint64(0)
	LENGTH   = 100
)

// Node 元素
type Node struct {
	prev *Node
	next *Node

	key   int // lru key
	value int // lru value

	hnext *Node // 拉链
}

// LRUCache lru缓存
type LRUCache struct {
	nodes []Node // hash list

	head *Node // lru head nodes
	tail *Node // lru tail nodes

	capacity int //
	used     int //
}

// NewLRUCache 初始化
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		nodes:    make([]Node, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

// Get 获取元素
func (l *LRUCache) Get(key int) int {
	if l.tail == nil {
		return -1
	}

	if tmp := l.searchNode(key); tmp != nil {
		l.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

// Put 添加元素
func (l *LRUCache) Put(key int, value int) {
	// 1. 首次插入数据
	// 2. 插入数据不在 LRU 中
	// 3. 插入数据在 LRU 中
	// 4. 插入数据不在 LRU 中, 并且 LRU 已满

	if tmp := l.searchNode(key); tmp != nil {
		tmp.value = value
		l.moveToTail(tmp)
		return
	}
	l.addNode(key, value)

	if l.used > l.capacity {
		l.deleteNode()
	}
}

// addNode 添加节点
func (l *LRUCache) addNode(key int, value int) {
	newNode := &Node{
		key:   key,
		value: value,
	}

	tmp := &l.nodes[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	l.used++

	if l.tail == nil {
		l.tail, l.head = newNode, newNode
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

// deleteNode 删除节点
func (l *LRUCache) deleteNode() {
	if l.head == nil {
		return
	}
	prev := &l.nodes[hash(l.head.key)]
	tmp := prev.hnext

	for tmp != nil && tmp.key != l.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	l.head = l.head.next
	l.head.prev = nil
	l.used--
}

// searchNode 查找结点
func (l *LRUCache) searchNode(key int) *Node {
	if l.tail == nil {
		return nil
	}

	// 查找
	tmp := l.nodes[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

// moveToTail 移动到尾部
func (l *LRUCache) moveToTail(node *Node) {
	if l.tail == node {
		return
	}
	if l.head == node {
		l.head = node.next
		l.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	node.next = nil
	l.tail.next = node
	node.prev = l.tail

	l.tail = node
}

// hash 散列算法
func hash(key int) int {
	if HOST_BIT {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}
