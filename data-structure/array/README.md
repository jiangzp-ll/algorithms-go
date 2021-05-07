### Array API

* `func NewArray(cap int) *Array` ：初始化数组
* `func (arr *Array) Clear()`：清空数组
* `func (arr *Array) Contains(val interface{}) bool`：判断数组中是否包含该元素
* `func (arr *Array) Data() []interface{}`：获取数组中所有元素
* `func (arr *Array) Delete(index int) (interface{}, error)`：删除指定下标的元素
* `func (arr *Array) Expansion()`：数组扩容
* `func (arr *Array) Get(index int) (interface{}, error)`：获取指定下标的元素
* `func (arr *Array) IndexOf(val interface{}) int`：返回元素在数组中的位置 
* `func (arr *Array) Insert(index int, val interface{}) error`：插入元素
* `func (arr *Array) IsEmpty() bool`：判断数组是否为空
* `func (arr *Array) IsIndexOutOfRange(index int) bool`：判断数组下标是否越界
* `func (arr *Array) Len() int` ：获取数组长度
* `func (arr *Array) MergeArray(other *Array) *Array`：合并两个数组
* `func (arr *Array) Replace(index int, val interface{}) error`：替换元素
* `func (arr *Array) Set(index int, val interface{}) interface{}`：设置指定下标的元素
* `func (arr *Array) Sort()`：数组排序
* `func (arr *Array) SubArray(start, end int) *Array`：获取子数组
* `func (arr *Array) ToString() string`：转成用“, ”分隔的字符串