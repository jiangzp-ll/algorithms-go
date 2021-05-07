### Array API
* `func NewArray(cap int) *Array` ：初始化数组
*  `func (arr *Array) Len() int`：获取数组长度
* `func (arr *Array) Data() []int`：获取数组中所有元素
* `func (arr *Array) IsIndexOutOfRange(index int) bool`：判断数组下标是否越界
* `func (arr *Array) Find(index int) (int, error)`：通过索引查找数组，索引范围[0,len-1]
* `func (arr *Array) Insert(index, val int) error`：插入元素
* `func (arr *Array) Delete(index int) (int, error)`：删除指定下标的元素
* `func (arr *Array) Expansion()`：数组扩容
* `func (arr *Array) MergeArray(other *Array) *Array`：合并两个数组
* `func (arr *Array) Split(start, end int) *Array`：按照指定下标切割数组