package single

import (
	"errors"
	"testing"
)

var list = NewLinkedList()

func TestNode_NewNode_Success(t *testing.T) {
	in := "a"
	node := NewNode(in)
	if node.Value() == in && node.Next() == nil {
		t.Logf("init node success")
	} else {
		t.Error("init node failed")
	}
}

func Test_LinkedList_Add(t *testing.T) {
	defer list.Clear()
	in := "a"
	list.Add(in)
	if list.Head().Value() == in && list.Head().Next() == nil && list.Len() == 1 {
		t.Log("add element to LinkedList with success")
	} else {
		t.Error("function Add has bug")
	}
}

func Test_LinkedList_AddToHead(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	list.AddToHead(in)
	if list.Head().Value() == in && list.Len() == len(elements)+1 {
		t.Log("add element to LinkedList head is success")
	} else {
		t.Error("function Add has bug")
	}
}

func Test_LinkedList_AllIndexesOf(t *testing.T) {
	defer list.Clear()
	flag := true
	target := []int{1, 3}
	in := "a"
	elements := []string{in, "b", in}
	for _, e := range elements {
		list.Add(e)
	}
	indexes, err := list.AllIndexesOf(in)
	if err != nil {
		t.Errorf("get element all index has error! error: %v \n", err)
		return
	}
	if len(indexes) != len(target) {
		t.Error("function AllIndexesOf has bug")
		return
	}
	for i := 0; i < len(indexes); i++ {
		if target[i] != indexes[i] {
			flag = false
		}
	}
	if flag && list.head.value == in {
		t.Log("get all index of the element")
	} else {
		t.Error("function AllIndexesOf has bug")
	}
}

func Test_LinkedList_AllIndexesOf_With_LinkedList_Is_Empty(t *testing.T) {
	in := "a"
	l := NewLinkedList()
	_, err := l.AllIndexesOf(in)
	if err != nil {
		if errors.Is(err, LinkedListIsEmptyError) {
			t.Log("LinkedList is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function AllIndexesOf has bug")
	}
}

func Test_LinkedList_Contain_With_Value_In_The_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	list.Add(in)
	isContain := list.Contain(in)
	if isContain {
		t.Log("the value in LinkedList")
	} else {
		t.Error("the value not in LinkedList")
	}
}

func Test_LinkedList_Contain_With_Value_Not_In_The_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	isContain := list.Contain(in)
	if !isContain {
		t.Log("the value not in LinkedList")
	} else {
		t.Error("function Contain has bug")
	}
}

func Test_LinkedList_Get_With_Index_In_Range(t *testing.T) {
	defer list.Clear()
	target := "b"
	index := 2
	elements := []string{"a", "b", "c"}
	for _, e := range elements {
		list.Add(e)
	}
	ret, err := list.Get(index)
	if err != nil {
		t.Errorf("get element has error! error: %v \n", err)
		return
	}
	if ret.Value() == target {
		t.Log("get element success")
	} else {
		t.Error("get element failed")
	}
}

func Test_LinkedList_Get_With_Index_Less_Than_Zero(t *testing.T) {
	defer list.Clear()
	index := -1
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	_, err := list.Get(index)
	if err != nil {
		if errors.Is(err, InvalidIndexError) {
			t.Log("index not in LinkedList range")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Get has bug")
	}
}

func Test_LinkedList_Get_With_Index_Out_Of_LinkedList_Range(t *testing.T) {
	defer list.Clear()
	elements := []int{1, 2, 3}
	index := len(elements) + 1
	for _, e := range elements {
		list.Add(e)
	}
	_, err := list.Get(index)
	if err != nil {
		if errors.Is(err, InvalidIndexError) {
			t.Log("index not in LinkedList range")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Get has bug")
	}
}

func Test_LinkedList_HasCycle(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	for _, e := range elements {
		list.Add(e)
	}
	list.head.next.next.next = list.head.next
	hasCycle := list.HasCycle()
	if hasCycle {
		t.Log("LinkedList has cycle")
	} else {
		t.Error("LinkedList not has cycle")
	}
}

func Test_LinkedList_IndexOf_With_The_Value_In_LinkedList(t *testing.T) {
	defer list.Clear()
	index := 2
	target := "b"
	elements := []string{"a", target, "c"}
	for _, e := range elements {
		list.Add(e)
	}
	ret, err := list.IndexOf(target)
	if err != nil {
		t.Errorf("get index of the value has error! error: %v \n", err)
		return
	}
	if index == ret {
		t.Log("get index of the value success")
	} else {
		t.Errorf("get index of the value failed")
	}
}

func Test_LinkedList_IndexOf_With_The_Value_Not_In_LinkedList(t *testing.T) {
	defer list.Clear()
	target := 1
	elements := []string{"a", "b", "c"}
	for _, e := range elements {
		list.Add(e)
	}
	_, err := list.IndexOf(target)
	if err != nil {
		if errors.Is(err, ValueNotExistError) {
			t.Log("value not in this LinkedList")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function IndexOf has bug")
	}
}

func Test_LinkedList_IndexOf_With_LinkedList_Is_Empty(t *testing.T) {
	l := NewLinkedList()
	_, err := l.IndexOf("a")
	if err != nil {
		if errors.Is(err, LinkedListIsEmptyError) {
			t.Log("LinkedList is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function IndexOf has bug")
	}
}

func Test_LinkedList_InsertAfter(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	index := 2
	node, _ := list.Get(index)
	if err := list.InsertAfter(node, in); err != nil {
		t.Errorf("insert value has error! error: %v \n", err)
		return
	}
	n, _ := list.Get(index + 1)
	if n.value == in && list.Len() == len(elements)+1 {
		t.Log("insert succeeded after the specified element")
	} else {
		t.Error("failed to insert after the specified element")
	}
}

func Test_LinkedList_InsertAfter_With_Not_Existed_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	node := NewNode("A")
	if err := list.InsertAfter(node, in); err != nil {
		if errors.Is(err, NodeNotExistError) {
			t.Log("node not existed")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func Test_LinkedList_InsertAfter_With_Empty_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	var node *Node
	if err := list.InsertAfter(node, in); err != nil {
		if errors.Is(err, InputNodeIsEmptyError) {
			t.Log("node is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func Test_LinkedList_InsertAfter_With_Empty_LinkedList(t *testing.T) {
	l := NewLinkedList()
	node := NewNode(1)
	in := "a"
	if err := l.InsertAfter(node, in); err != nil {
		if errors.Is(err, LinkedListIsEmptyError) {
			t.Log("LinkedList is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func Test_LinkedList_InsertBefore(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	index := 2
	node, _ := list.Get(index)
	if err := list.InsertBefore(node, in); err != nil {
		t.Errorf("insert value has error! error: %v \n", err)
		return
	}
	n, _ := list.Get(index)
	if n.value == in && list.Len() == len(elements)+1 {
		t.Log("insert succeeded after the specified element")
	} else {
		t.Error("failed to insert after the specified element")
	}
}

func Test_LinkedList_InsertBefore_With_Not_Existed_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	node := NewNode("A")
	if err := list.InsertBefore(node, in); err != nil {
		if errors.Is(err, NodeNotExistError) {
			t.Log("node not existed")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func Test_LinkedList_InsertBefore_With_Empty_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	var node *Node
	if err := list.InsertBefore(node, in); err != nil {
		if errors.Is(err, InputNodeIsEmptyError) {
			t.Log("node is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func Test_LinkedList_InsertBefore_With_Empty_LinkedList(t *testing.T) {
	l := NewLinkedList()
	node := NewNode(1)
	in := "a"
	if err := l.InsertBefore(node, in); err != nil {
		if errors.Is(err, LinkedListIsEmptyError) {
			t.Log("LinkedList is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func Test_LinkedList_Last(t *testing.T) {
	defer list.Clear()
	target := 3
	elements := []int{1, 2, target}
	for _, e := range elements {
		list.Add(e)
	}
	node := list.Last()
	if node.Value() == target && node.Next() == nil {
		t.Log("get last node success")
	} else {
		t.Error("get last node failed")
	}
}

func Test_LinkedList_Last_With_Empty_LinkedList(t *testing.T) {
	l := NewLinkedList()
	node := l.Last()
	if node == nil {
		t.Log("LinkedList is empty")
	} else {
		t.Error("function Last has bug")
	}
}

func Test_LinkedList_LastIndexOf(t *testing.T) {
	defer list.Clear()
	ti := 5
	in := 2
	elements := []int{1, in, 3, 4, in}
	for _, e := range elements {
		list.Add(e)
	}
	index, err := list.LastIndexOf(in)
	if err != nil {
		t.Errorf("get last index of value has error! error: %v \n", err)
		return
	}
	if ti == index {
		t.Log("get last index of the value is success")
	} else {
		t.Error("get last index of the value is failed")
	}
}

func Test_LinkedList_LastIndexOf_With_Value_Not_In_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	_, err := list.LastIndexOf(in)
	if err != nil {
		if errors.Is(err, ValueNotExistError) {
			t.Log("value not exist")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function LastIndexOf has bug")
	}
}

func Test_LinkedList_LastIndexOf_With_LinkedList_Is_Empty(t *testing.T) {
	l := NewLinkedList()
	in := "a"
	_, err := l.LastIndexOf(in)
	if err != nil {
		if errors.Is(err, LinkedListIsEmptyError) {
			t.Log("LinkedList is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function LastIndexOf has bug")
	}
}

//
//import (
//	"testing"
//)
//
//func TestLinkedList_InsertToHead(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 10; i++ {
//		linkedList.InsertToHead(i)
//	}
//	node := linkedList.FindByIndex(0)
//	t.Log(node.Value())
//	if node.Value() == 9 {
//		t.Log("成功从链表头部插入！")
//	} else {
//		t.Error("从链表头部插入失败！！！！")
//	}
//}
//
//func TestLinkedList_InsertToTail(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 10; i++ {
//		linkedList.Add(i)
//	}
//	node := linkedList.FindByIndex(linkedList.Len() - 1)
//	t.Log(node.Value())
//	if node.Value() == 9 {
//		t.Log("成功从链表尾部插入！")
//	} else {
//		t.Error("从链表尾部插入失败！！！！")
//	}
//}
//
//func TestLinkedList_InsertAfter(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.FindByIndex(2)
//	flag := linkedList.InsertAfter(node, 100)
//	if flag {
//		t.Logf("after: ")
//		linkedList.Print()
//		t.Log("某个节点后插入成功")
//	} else {
//		t.Error("某个节点后插入失败")
//	}
//}
//
//func TestLinkedList_InsertBefore(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.FindByIndex(2)
//	flag := linkedList.InsertBefore(node, 100)
//	if flag {
//		t.Logf("after: ")
//		linkedList.Print()
//		t.Log("某个节点前插入成功")
//	} else {
//		t.Error("某个节点前插入失败")
//	}
//}
//
//func TestLinkedList_DeleteNode(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.FindByIndex(2)
//	flag := linkedList.DeleteNode(node)
//	if flag {
//		t.Logf("after: ")
//		linkedList.Print()
//		t.Log("节点删除成功")
//	} else {
//		t.Error("节点删除失败")
//	}
//}
//
//func TestLinkedList_Reverse(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	linkedList.Reverse()
//	if linkedList.FindByIndex(4).Value() == 0 {
//		t.Log("单链表反转成功")
//		t.Logf("after: ")
//		linkedList.Print()
//	} else {
//		t.Error("单链表反转失败")
//	}
//}
//
////func TestLinkedList_HasCycle(t *testing.T) {
////	arr := []string{"a", "b", "c", "d", "b", "c", "d"}
////	linkedList := linkedlist.NewLinkedList()
////	for _, a := range arr {
////		linkedList.Add(a)
////	}
////	linkedList.Head().Next().Next().Next().Next() = linkedList.Head().Next()
////	hasCycle := linkedList.HasCycle()
////	if hasCycle {
////		t.Log("链表有环")
////	} else {
////		t.Error("链表没有环")
////	}
////}
//
//func TestLinkedList_DeleteReciprocal(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	t.Logf("before: ")
//	linkedList.Print()
//	node := linkedList.DeleteReciprocal(3)
//	if nil != node {
//		t.Log("删除倒数第三个元素成功,该元素是：", node.Value())
//		t.Logf("after: ")
//		linkedList.Print()
//	} else {
//		t.Error("删除倒数第二个元素成功")
//	}
//}
//
//func TestLinkedList_GetMiddle(t *testing.T) {
//	linkedList := NewLinkedList()
//	for i := 0; i < 5; i++ {
//		linkedList.Add(i)
//	}
//	middle := linkedList.Middle()
//	if nil != middle && middle.Value() == 2 {
//		t.Log("获取链表中间节点成功")
//	} else {
//		t.Error("获取链表中间节点失败")
//	}
//}
