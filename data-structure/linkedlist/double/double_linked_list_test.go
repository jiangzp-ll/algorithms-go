package double

import (
	"errors"
	"github.com/zepeng-jiang/go-basic-demo/data-structure/linkedlist/single"
	"testing"
)

var list = NewLinkedList()

func Test_Node_NewNode(t *testing.T) {
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
	in := 100
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	list.Add(in)
	if list.Head().Value() == elements[0] && list.Tail().Value() == in && list.Len() == len(elements)+1 {
		t.Log("add element to LinkedList with success")
	} else {
		t.Error("function Add has bug")
	}
}

func Test_LinkedList_Add_With_LinkedList_Is_Empty(t *testing.T) {
	defer list.Clear()
	in := "a"
	list.Add(in)
	if list.Head().Value() == in && list.Tail().Value() == in && list.Len() == 1 {
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

func Test_LinkedList_AddToHead_With_LinkedList_Is_Empty(t *testing.T) {
	defer list.Clear()
	in := "a"
	list.AddToHead(in)
	if list.Head().Value() == in && list.Tail().Value() == in && list.Len() == 1 {
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

func Test_LinkedList_AllIndexesOf_With_LinkedList_Number_Is_Even(t *testing.T) {
	defer list.Clear()
	flag := true
	target := []int{1, 3}
	in := "a"
	elements := []string{in, "b", in, "c"}
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

func Test_LinkedList_AllIndexesOf_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	flag := true
	target := []int{1}
	in := "a"
	list.Add(in)
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

func Test_LinkedList_AllIndexesOf_With_LinkedList_Only_One_Element_And_Value_Not_In_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "a"
	list.Add("b")
	_, err := list.AllIndexesOf(in)
	if err != nil {
		if errors.Is(err, single.ValueNotExistError) {
			t.Log("value not exist")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function AllIndexesOf has bug")
	}
}

func Test_LinkedList_AllIndexesOf_With_LinkedList_Is_Empty(t *testing.T) {
	in := "a"
	l := NewLinkedList()
	_, err := l.AllIndexesOf(in)
	if err != nil {
		if errors.Is(err, single.LinkedListIsEmptyError) {
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

func Test_LinkedList_Clear(t *testing.T) {
	elements := []string{"a", "b", "c"}
	for _, e := range elements {
		list.Add(e)
	}
	list.Clear()
	if list.len == 0 {
		t.Log("clear Double LinkedList is success")
	} else {
		t.Error("clear Double LinkedList is failed")
	}
}

func Test_LinkedList_Contain(t *testing.T) {
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

func Test_LinkedList_Get(t *testing.T) {
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
		if errors.Is(err, single.InvalidIndexError) {
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
		if errors.Is(err, single.InvalidIndexError) {
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

func Test_LinkedList_HasCycle_With_Not_Has_Cycle(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	for _, e := range elements {
		list.Add(e)
	}
	hasCycle := list.HasCycle()
	if !hasCycle {
		t.Log("LinkedList not has cycle")
	} else {
		t.Error("function HasCycle has bug")
	}
}

func Test_LinkedList_HasCycle_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	list.Add("a")
	hasCycle := list.HasCycle()
	if !hasCycle {
		t.Log("LinkedList not has cycle")
	} else {
		t.Error("function HasCycle has bug")
	}
}

func Test_LinkedList_HasCycle_With_LinkedList_Is_Empty(t *testing.T) {
	l := NewLinkedList()
	hasCycle := l.HasCycle()
	if !hasCycle {
		t.Log("LinkedList not has cycle")
	} else {
		t.Error("function HasCycle has bug")
	}
}

func Test_LinkedList_IndexOf(t *testing.T) {
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
		if errors.Is(err, single.ValueNotExistError) {
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
		if errors.Is(err, single.LinkedListIsEmptyError) {
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
	if n.value == in && n.Prev() == node && list.Len() == len(elements)+1 {
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
		if errors.Is(err, single.NodeNotExistError) {
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
		if errors.Is(err, single.InputNodeIsEmptyError) {
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
		if errors.Is(err, single.LinkedListIsEmptyError) {
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

func Test_LinkedList_InsertBefore_With_Head(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	index := 1
	node, _ := list.Get(index)
	if err := list.InsertBefore(node, in); err != nil {
		t.Errorf("insert value has error! error: %v \n", err)
		return
	}
	if list.Head().value == in && list.Len() == len(elements)+1 {
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
		if errors.Is(err, single.NodeNotExistError) {
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
		if errors.Is(err, single.InputNodeIsEmptyError) {
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
		if errors.Is(err, single.LinkedListIsEmptyError) {
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
		if errors.Is(err, single.ValueNotExistError) {
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
		if errors.Is(err, single.LinkedListIsEmptyError) {
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

func Test_LinkedList_Len(t *testing.T) {
	defer list.Clear()
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	if list.Len() == len(elements) {
		t.Log("get LinkedList len is success")
	} else {
		t.Error("get LinkedList len is failed")
	}
}

func Test_LinkedList_Len_With_LinkedList_Is_Empty(t *testing.T) {
	if list.Len() == 0 {
		t.Log("get LinkedList len is success")
	} else {
		t.Error("get LinkedList len is failed")
	}
}

func Test_LinkedList_Remove(t *testing.T) {
	defer list.Clear()
	index := 2
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	node, _ := list.Get(index)
	if err := list.Remove(node); err != nil {
		t.Errorf("remove node has error! error: %v \n", err)
		return
	}
	if list.Len() == len(elements)-1 {
		t.Log("remove node is success")
	} else {
		t.Error("remove node is failed")
	}
}

func Test_LinkedList_Remove_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	index := 1
	in := "a"
	list.Add(in)
	node, _ := list.Get(index)
	if err := list.Remove(node); err != nil {
		t.Errorf("remove node has error! error: %v \n", err)
		return
	}
	if list.Len() == 0 {
		t.Log("remove node is success")
	} else {
		t.Error("remove node is failed")
	}
}

func Test_LinkedList_Remove_With_LinkedList_Is_Empty(t *testing.T) {
	defer list.Clear()
	node := NewNode("a")
	if err := list.Remove(node); err != nil {
		if errors.Is(err, single.LinkedListIsEmptyError) {
			t.Log("LinkedList is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Remove has bug")
	}
}

func Test_LinkedList_Remove_With_Node_Not_Existed(t *testing.T) {
	defer list.Clear()
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	node := NewNode("a")
	if err := list.Remove(node); err != nil {
		if errors.Is(err, single.NodeNotExistError) {
			t.Log("node not exited")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Remove has bug")
	}
}

func Test_LinkedList_Remove_With_Node_Is_Empty(t *testing.T) {
	defer list.Clear()
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	var node *Node
	if err := list.Remove(node); err != nil {
		if errors.Is(err, single.InputNodeIsEmptyError) {
			t.Log("input node is empty")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Remove has bug")
	}
}

func Test_LinkedList_RemoveOf(t *testing.T) {
	defer list.Clear()
	index := 2
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	node, err := list.RemoveOf(index)
	if err != nil {
		t.Errorf("remove the specified index node has error! error: %v \n", err)
		return
	}
	if list.Len() == len(elements)-1 && node.Value() == elements[index-1] {
		t.Log("remove the specified index node is success")
	} else {
		t.Error("remove the specified index node is failed")
	}
}

func Test_LinkedList_RemoveOf_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	index := 1
	in := "a"
	list.Add(in)
	node, err := list.RemoveOf(index)
	if err != nil {
		t.Errorf("remove the specified index node has error! error: %v \n", err)
		return
	}
	if list.Len() == 0 && node.Value() == in {
		t.Log("remove the specified index node is success")
	} else {
		t.Error("remove the specified index node is failed")
	}
}

func Test_LinkedList_RemoveOf_With_Index_Not_In_Range(t *testing.T) {
	defer list.Clear()
	elements := []int{1, 2, 3}
	index := len(elements) + 1
	for _, e := range elements {
		list.Add(e)
	}
	_, err := list.RemoveOf(index)
	if err != nil {
		if errors.Is(err, single.InvalidIndexError) {
			t.Log("invalid index")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function RemoveOf has bug")
	}
}

func Test_LinkedList_Reverse(t *testing.T) {
	defer list.Clear()
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	list.Reverse()
	cur := list.head
	flag := true
	for i := len(elements) - 1; i >= 0; i-- {
		if cur.value != elements[i] {
			flag = false
			break
		}
		cur = cur.next
	}
	if flag && list.Len() == len(elements) {
		t.Log("reverse LinkedList is success")
	} else {
		t.Error("reverse LinkedList is failed")
	}
}

func Test_LinkedList_Set(t *testing.T) {
	defer list.Clear()
	target := "a"
	index := 2
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	old, err := list.Set(index, target)
	if err != nil {
		t.Errorf("set value has error! error: %v \n", err)
		return
	}
	node, _ := list.Get(index)
	if old == 2 && node.Value() == target {
		t.Log("set value is success")
	} else {
		t.Error("set value is failed")
	}
}

func Test_LinkedList_Set_With_Invalid_Index(t *testing.T) {
	defer list.Clear()
	target := "a"
	index := -1
	elements := []int{1, 2, 3}
	for _, e := range elements {
		list.Add(e)
	}
	_, err := list.Set(index, target)
	if err != nil {
		if errors.Is(err, single.InvalidIndexError) {
			t.Log("invalid index")
			return
		} else {
			t.Errorf("unknown error! error: %v \n", err)
			return
		}
	} else {
		t.Error("function Set has bug")
	}
}
