package skiplist

import "testing"

func TestSkipList_Insert(t *testing.T) {
	list := NewSkipList()
	list.Insert("A", 65)
	list.Insert("B", 66)
	list.Insert("C", 67)
	list.Insert("a", 97)
	list.Insert("b", 98)
	list.Insert("c", 99)
	if list.length == 6 && list.level == 5{
		t.Log("插入成功")
	} else {
		t.Error("插入失败")
	}
}

func TestSkipList_Find(t *testing.T) {
	list := NewSkipList()
	list.Insert("A", 65)
	list.Insert("B", 66)
	list.Insert("C", 67)
	list.Insert("a", 97)
	list.Insert("b", 98)
	list.Insert("c", 99)
	node := list.Find("b", 98)
	if node.v.(string) == "b" && node.score == 98 {
		t.Log("查找成功")
	} else {
		t.Error("查找失败")
	}
}

func TestSkipList_Delete(t *testing.T) {
	list := NewSkipList()
	list.Insert("A", 65)
	list.Insert("B", 66)
	list.Insert("C", 67)
	list.Insert("a", 97)
	list.Insert("b", 98)
	list.Insert("c", 99)
	ret := list.Delete("b", 98)
	node := list.Find("b", 98)
	if nil == node && ret == 0 {
		t.Log("删除成功")
	} else {
		t.Error("删除失败")
	}
}