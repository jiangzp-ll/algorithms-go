### Stack API
* <a href="#stack"><font size=4 color=#00f>type Stack</font></a>：栈
    * `func NewStack() *Stack` ：初始化栈
    * `func Flush()`：清空栈
    * `func IsEmpty() bool`：判断栈是否为空
    * `func Pop() interface{}`：元素出栈
    * `func Push(v interface{})`：元素入栈
    * `func Top() interface{}`：返回栈顶元素
 #### <a id="stack">type Stack</a>
```
type Stack interface {
    // Flush ,clear the stack
    Flush()
    // IsEmpty ,determine whether the stack is empty
    IsEmpty() bool
    // Pop , pop the element from the top of the stack
    Pop() interface{}
    // Push ,push the element to top of the stack
    Push(v interface{})
    // Top , get top of the stack
    Top() interface{}
}
```