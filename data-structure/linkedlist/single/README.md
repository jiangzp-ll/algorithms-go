### Single LinkedList API
1. <a href="#node"><font size=4 color=#00f>type Node</font></a>：单向链表节点
    * `func func NewNode(v interface{}) *Node`：初始化节点
    * `func (n *Node) Next() *Node`：获取下一个节点
    * `func (n *Node) Value() interface{}`：获取节点值
2. <a href="#list"><font size=4 color=#00f>type LinkedList</font></a>：单单向链表
    * `func NewLinkedList() *LinkedList`：初始化单向链表
    * `func (l *LinkedList) Add(val interface{}) error`：添加元素到单向链表的尾部
    * `func (l *LinkedList) AddToHead(v interface{}) error`：插入元素单向链表头部
    * `func (l *LinkedList) AllIndexesOf(val interface{}) ([]int, error)`：返回指定元素在单向链表中出现的所有索引
    * `func (l *LinkedList) checkElementIndex(index int) bool`：检查下标是否在单向链表中
    * `func (l *LinkedList) Clear()`：清空单向链表
    * `func (l *LinkedList) Contain(val interface{}) (bool, error)`：判断单向链表中是否包含该元素
    * `func (l *LinkedList) Get(index int) (*Node, error)`：获取单向链表尾部
    * `func (l *LinkedList) HasCycle() bool`：判断单向链表是否有环
    * `func (l *LinkedList) Head() *Node`：获取单向链表头部
    * `func (l *LinkedList) IndexOf(val interface{}) (int, error)`：返回元素在单向链表中的位置 
    * `func (l *LinkedList) InsertAfter(p *Node, v interface{}) error`：在单向链表指定下标后插入元素
    * `func (l *LinkedList) InsertBefore(p *Node, v interface{}) error`：在单向链表指定下标前插入元素
    * `func (l *LinkedList) Last() *Node`：获取单向链表尾部
    * `func (l *LinkedList) LastIndexOf(val interface{}) (int, error)`：返回指定元素在单向链表中最后一次出现的索引
    * `func (l *LinkedList) Len() int`：获取单向链表的元素个数
    * `func (l *LinkedList) Middle() *Node`：获取单向链表中间节点
    * `func (l *LinkedList) Remove(n *Node) error`：删除单向链表的指定节点
    * `func (l *LinkedList) Reverse()`：单向链表反转
    * `func (l *LinkedList) Set(index int, val interface{}) (interface{}, error)`：设置单向链表指定下标元素的值  
    
#### <a id="node">type Node</a>
```
type Node struct {
    // next, next Node ptr
    next  *Node
    // value, Node store the value
    value interface{}
}
```

#### <a id="list">type LinkedList</a>
```
type LinkedList struct {
    // head , first Node in LinkedList
    head *Node
    // len , number of elements in LinkedList
    len  int
}
```