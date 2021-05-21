### Array API
* <a href="#array"><font size=4 color=#00f>type Array</font></a>：数组，类似 Java 的 ArrayList  
    * `func NewArray(cap int) *Array` ：初始化数组
    * `func (arr *Array) Add(val interface{})`：添加元素到末尾
    * `func (arr *Array) Clear()`：清空数组
    * `func (arr *Array) Contain(val interface{}) bool`：判断数组中是否包含该元素
    * `func (arr *Array) Data() []interface{}`：获取数组中所有元素
    * `func (arr *Array) expansion()`：数组扩容
    * `func (arr *Array) Get(index int) (interface{}, error)`：获取指定下标的元素
    * `func (arr *Array) IndexOf(val interface{}) (int, error)`：返回元素在数组中的位置 
    * `func (arr *Array) Insert(index int, val interface{}) error`：插入元素
    * `func (arr *Array) IsEmpty() bool`：判断数组是否为空
    * `func (arr *Array) isIndexOutOfRange(index int) bool`：判断数组下标是否越界
    * `func (arr *Array) Len() int` ：获取数组元素个数
    * `func (arr *Array) MergeArray(other *Array) *Array`：合并两个数组
    * `func (arr *Array) Remove(index int) (interface{}, error)`：删除并返回指定下标的元素
    * `func (arr *Array) Replace(index int, val interface{}) (interface{}, error)`：替换元素并返回旧值
    * `func (arr *Array) Set(index int, val interface{}) error`：设置指定下标的元素值
    * `func (arr *Array) SubArray(start, end int) *Array`：获取子数组
    * `func (arr *Array) Tointerface{}() interface{}`：转成用“, ”分隔的字符串  
 #### <a id="array">type Array</a>
```
type Array struct {
    // data ,Array data
    data []interface{}
    // len ,number of elements in array
    len  int
}
```