### Queue API
* <a href="#queue"><font size=4 color=#00f>type Queue</font></a>：队列
    * `func NewQueue() *Queue` ：初始化队列
    * `func Add(val interface{}) error`：元素入队
    * `func Clear()`：清空队列
    * `func Contain(val interface{}) bool`：判断元素是否在队列中
    * `func IsEmpty() bool`：判断队列是否为空
    * `func Len() int`：获取队列中元素的个数
    * `funx Peek() (interface{}, error)`：获取队列头部元素但不删除
    * `func Remove() (interface{}, error)`：获取队列头部元素并删除
 #### <a id="queue">type Queue</a>
```
type Queue interface {
    // Add ,add the element to the end of the queue
    Add(val interface{}) error
    // Clear ,clear the queue
    Clear()
    // Contain ,determine whether the value contain in the queue
    Contain(val interface{}) bool
    // IsEmpty ,determine whether the queue is empty
    IsEmpty() bool
    // Len ,get the number of elements in the queue
    Len() int
    // Peek ,get and not remove the element from the header of the queue
    Peek() (interface{}, error)
    // Remove ,get and remove the element from the header of the queue
    Remove() (interface{}, error)
}
```