### Single LinkedList API
* `func (l *LinkedList) Add(val interface{}) bool`：添加元素到链表
* `func (l *LinkedList) Clear()`：清空链表
* `func (l *LinkedList) Contain(val interface{}) bool`：判断链表中是否包含该元素
* `func (l *LinkedList) Get(index int) (interface{}, error)`：获取链表尾部
* `func (l *LinkedList) GetFirst() *Node`：获取链表头部
* `func (l *LinkedList) GetLast() *Node`：获取链表尾部
* `func (l *LinkedList) GetMiddle() *Node`：获取链表中间节点
* `func (l *LinkedList) HasCycle() bool`：判断链表是否有环
* `func (l *LinkedList) IndexOf(val interface{}) (int, error)`：返回元素在数组中的位置 
* `func (l *LinkedList) InsertAfter(p *Node, v interface{})`：在指定下标后插入元素
* `func (l *LinkedList) InsertBefore(p *Node, v interface{})`：在指定下标前插入元素
* `func (l *LinkedList) InsertToHead(v interface{})`：插入元素链表头部
* `func (l *LinkedList) InsertToTail(val interface{})`：插入元素到链表尾部
* `func (l *LinkedList) isIndexOutOfRange(index int) bool`：判断下标是否在链表内
* `func (l *LinkedList) LastIndexOf(val interface{}) int`：返回指定元素在链表中最后一次出现的索引
* `func (l *LinkedList) Len() int`：获取链表的元素个数
* `func (l *LinkedList) Remove(index int) (interface{}, error)`：获取链表头部
* `func (l *LinkedList) Reverse()`：单链表反转
* `func (l *LinkedList) Set(index int, val interface{}) (interface{}, error)`：设置指定下标元素的值
