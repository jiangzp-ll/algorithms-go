### Stack API
* <a href="#stack"><font size=4 color=#00f>type Stack</font></a>：栈
    * `func NewStack() *Stack` ：初始化栈
    * `func Flush()`：清空栈
    * `func IsEmpty() bool`：判断栈是否为空
    * `funx Peek() (interface{}, error)`：获取栈顶元素但不删除
    * `func Pop() interface{}`：获取栈顶元素并删除
    * `func Push(v interface{})`：将元素放到栈顶
    * `func Search(val interface{}) (int, error)`：返回该值在栈中的第几个位置
 #### <a id="stack">type Stack</a>
```
type Stack interface {
    // Flush ,clear the stack
    Flush()
    // IsEmpty ,determine whether the stack is empty
    IsEmpty() bool
    // Peek ,get and not remove the element from the top of the stack
    Peek() (interface{}, error)
    // Pop ,pop and remove the element from the top of the stack
    Pop() (interface{}, error)
    // Push ,push the element to top of the stack
    Push(v interface{})
    // Search ,return the index of the value in the stack
    Search(val interface{}) (int, error)
}
```