### ArrayList API
* <a href="#array"><font size=4 color=#00f>type ArrayList</font></a>：数组，类似 Java 的 ArrayListList  
    * `func NewArrayList(typeOf string) (*ArrayList, error)` ：初始化数组，包含数组中元素类型
    * `func NewArrayList(typeOf string, cap int) (*ArrayList, error)` ：初始化数组，包含数组中元素类型以及容量
    * `func (arr *ArrayList) Add(val interface{}) error`：添加元素到末尾
    * `func (arr *ArrayList) checkType(val interface{}) error`：检查输入的 value 的类型是否和初始化一致
    * `func (arr *ArrayList) Clear()`：清空数组
    * `func (arr *ArrayList) Contain(val interface{}) bool`：判断数组中是否包含该元素
    * `func (arr *ArrayList) Data() []interface{}`：获取数组中所有元素
    * `func (arr *ArrayList) expansion()`：数组扩容
    * `func (arr *ArrayList) Get(index int) (interface{}, error)`：获取指定下标的元素
    * `func (arr *ArrayList) IndexOf(val interface{}) (int, error)`：返回元素在数组中的位置 
    * `func (arr *ArrayList) Insert(index int, val interface{}) error`：插入元素
    * `func (arr *ArrayList) IsEmpty() bool`：判断数组是否为空
    * `func (arr *ArrayList) isIndexOutOfRange(index int) bool`：判断数组下标是否越界
    * `func (arr *ArrayList) Len() int` ：获取数组元素个数
    * `func (arr *ArrayList) MergeArrayList(other *ArrayList) *ArrayList`：合并两个数组
    * `func (arr *ArrayList) Remove(index int) (interface{}, error)`：删除并返回指定下标的元素
    * `func (arr *ArrayList) Replace(index int, val interface{}) (interface{}, error)`：替换元素并返回旧值
    * `func (arr *ArrayList) Set(index int, val interface{}) error`：设置指定下标的元素值
    * `func (arr *ArrayList) SubArrayList(start, end int) *ArrayList`：获取子数组
 #### <a id="array">type ArrayList</a>
```
type ArrayList struct {
    // data ,store input vale
    data []interface{}
    // len ,number of elements in ArrayList
    len int
    // typeOf ,ArrayList type
    // Because Go not have Generic
    typeOf string
}
```