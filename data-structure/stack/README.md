### Stack API
* <a href="#stack"><font size=4 color=#00f>type Stack</font></a>：栈
    * `func NewStack() *Stack` ：初始化栈
    * `func Check(val interface{}) error`：检查输入元素类型是否和声明的一致
    * `func Flush()`：清空栈
    * `func IsEmpty() bool`：判断栈是否为空
    * `func Len() int`：获取栈中元素的个数
    * `funx Peek() (interface{}, error)`：获取栈顶元素但不删除
    * `func Pop() interface{}`：获取栈顶元素并删除
    * `func Push(val interface{})`：将元素放到栈顶
    * `func Search(val interface{}) (int, error)`：返回该值在栈中的第几个位置
 #### <a id="stack">type Stack</a>
```
type Stack interface {
    // Check ,check input value type
    Check(val interface{}) error
    // Flush ,clear the stack
    Flush()
    // IsEmpty ,determine whether the stack is empty
    IsEmpty() bool
    // Len ,get the number of elements in the stack
    Len() int
    // Peek ,get and not remove the element from the top of the stack
    Peek() (interface{}, error)
    // Pop ,pop and remove the element from the top of the stack
    Pop() (interface{}, error)
    // Push ,push the element to top of the stack
    Push(val interface{}) error
    // Search ,return the index of the value in the stack
    Search(val interface{}) (int, error)
}
```