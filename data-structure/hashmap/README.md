### HashMap API
* <a href="#hash"><font size=4 color=#00f>type HashMap</font></a>：HashMap
    * `func NewHashMap() *HashMap` ：初始化一个默认容量的 HashMap 
    * `func NewHashMapWithCap(c int) (*HashMap, error)` ：初始化一个容量为 c 的 HashMap 
    * `func Clear()`：清空 HashMap
    * `func ContainKey(key string) bool`：判断 key 是否在 HashMap 中
    * `func ContainValue(value interface{}) bool`：判断 value 是否在 HashMap 中
    * `func EntrySet() []*Entry`：包含 key 和 value的集合
    * `func Get(key string) (interface{}, error)`：根据 key 来获取在 HashMap 中的位置
    * `func hash(key string) int`：获取 key 的 hash code
    * `func hashCode(key string) int`：获取 key 的 hash code
    * `func IsEmpty() bool`：判断 HashMap 是否为空
    * `func KeySet() []string`：获取所有 key 的集合
    * `funx Put(key string, value interface{}) error`：将 key 和 value 放入 HashMap 中
    * `func Remove(key string) (interface{}, error)`：根据 key 来删除 HashMap 中的 value
    * `func Replace(key string, val interface{}) error`：根据 key 来替换 HashMap 中的 value
    * `func Values() []interface{}`：获取所有 value 的集合
 #### <a id="hash">type HashMap</a>
```
type HashMap interface {
    // Clear ,clear the HashMap
    Clear()
    // ContainKey ,determine whether the key is in the HashMap
    ContainKey(key string) bool
    // ContainValue ,determine whether the value is in the HashMap
    ContainValue(value interface{}) bool
    // EntrySet ,Set containing key and value
    EntrySet() []*Entry
    // Get ,get the value according to the key
    Get(key string) (interface{}, error)
    // IsEmpty ,determine whether the HashMap is empty
    IsEmpty() bool
    // KeySet ,get the Set of all keys
    KeySet() []string
    // Put ,put key and value into HashMap
    Put(key string, value interface{}) error
    // Remove ,remove the value in HashMap according to the key
    Remove(key string) (interface{}, error)
    // Replace ,replace the value in HashMap according to the key
    Replace(key string, val interface{}) error
    // Values ,get the Set of all values
    Values() []interface{}
}
```