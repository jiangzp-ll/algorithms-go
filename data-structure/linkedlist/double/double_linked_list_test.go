package double

import (
	"errors"
	errors2 "github.com/zepeng-jiang/go-basic-demo/pkg/errors"
	"testing"
)

var (
	list = initDoubleLinkedList("string")
	l    = initDoubleLinkedList("int")
)

func TestNode_NewNode(t *testing.T) {
	in := "a"
	node := NewNode(in)
	if node.Value() == in && node.Next() == nil {
		t.Logf("init node success")
	} else {
		t.Error("init node failed")
	}
}

func TestNewLinkedList(t *testing.T) {
	ll, err := NewLinkedList("string")
	if err != nil {
		t.Errorf("init a LinkedList has error! error: %v ", err)
		return
	}
	t.Logf("init a LinkedList is success! LinkedList: %v ", ll)
}

func TestNewLinkedList_With_Type_Is_Empty(t *testing.T) {
	_, err := NewLinkedList("")
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be not empty")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Errorf("function NewLinkedList has bug")
	}
}

func TestLinkedList_Add(t *testing.T) {
	defer list.Clear()
	in := "d"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	if err := list.Add(in); err != nil {
		t.Errorf("add value to LinkedList has error! error: %v ", err)
		return
	}
	if list.Head().Value() == elements[0] && list.Tail().Value() == in && list.Len() == len(elements)+1 {
		t.Log("add element to LinkedList with success")
	} else {
		t.Error("function Add has bug")
	}
}

func TestLinkedList_Add_With_LinkedList_Is_Empty(t *testing.T) {
	defer list.Clear()
	in := "a"
	if err := list.Add(in); err != nil {
		t.Errorf("add value to LinkedList has error! error: %v ", err)
		return
	}
	if list.Head().Value() == in && list.Tail().Value() == in && list.Len() == 1 {
		t.Log("add element to LinkedList with success")
	} else {
		t.Error("function Add has bug")
	}
}

func TestLinkedList_Add_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	in := 1
	if err := list.Add(in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Add has bug")
	}
}

func TestLinkedList_AddToHead(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	if err := list.AddToHead(in); err != nil {
		t.Errorf("add value to LinkedList head has error! error: %v ", err)
		return
	}
	if list.Head().Value() == in && list.Len() == len(elements)+1 {
		t.Log("add element to LinkedList head is success")
	} else {
		t.Error("function Add has bug")
	}
}

func TestLinkedList_AddToHead_With_LinkedList_Is_Empty(t *testing.T) {
	defer list.Clear()
	in := "a"
	if err := list.AddToHead(in); err != nil {
		t.Errorf("add value to LinkedList head has error! error: %v ", err)
		return
	}
	if list.Head().Value() == in && list.Tail().Value() == in && list.Len() == 1 {
		t.Log("add element to LinkedList head is success")
	} else {
		t.Error("function Add has bug")
	}
}

func TestLinkedList_AddToHead_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	in := 1
	if err := list.AddToHead(in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Add has bug")
	}
}

func TestLinkedList_AllIndexesOf(t *testing.T) {
	defer list.Clear()
	flag := true
	target := []int{1, 3}
	in := "a"
	elements := []string{in, "b", in}
	addValueToLinkedList(elements)
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

func TestLinkedList_AllIndexesOf_With_LinkedList_Number_Is_Even(t *testing.T) {
	defer list.Clear()
	flag := true
	target := []int{1, 3}
	in := "a"
	elements := []string{in, "b", in, "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_AllIndexesOf_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	flag := true
	target := []int{1}
	in := "a"
	_ = list.Add(in)
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

func TestLinkedList_AllIndexesOf_With_LinkedList_Only_One_Element_And_Value_Not_In_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "a"
	_ = list.Add("b")
	_, err := list.AllIndexesOf(in)
	if err != nil {
		if errors.Is(err, errors2.ValueNotExistError) {
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

func TestLinkedList_AllIndexesOf_With_LinkedList_Is_Empty(t *testing.T) {
	defer l.Clear()
	in := 1
	_, err := l.AllIndexesOf(in)
	if err != nil {
		if errors.Is(err, errors2.LinkedListIsEmptyError) {
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

func TestLinkedList_AllIndexesOf_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	in := 1
	_, err := list.AllIndexesOf(in)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function AllIndexesOf has bug")
	}
}

func TestLinkedList_Check_With_Input_Is_Nil(t *testing.T) {
	defer list.Clear()
	_ = list.Add("a")
	if err := list.Check(nil); err != nil {
		t.Log("input value type not same with LinkedList type.")
	} else {
		t.Error("function Check has bug")
	}
}

func TestLinkedList_Clear(t *testing.T) {
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	list.Clear()
	if list.len == 0 {
		t.Log("clear Double LinkedList is success")
	} else {
		t.Error("clear Double LinkedList is failed")
	}
}

func TestLinkedList_Contain(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	_ = list.Add(in)
	isContain := list.Contain(in)
	if isContain {
		t.Log("the value in LinkedList")
	} else {
		t.Error("the value not in LinkedList")
	}
}

func TestLinkedList_Contain_With_Value_Not_In_The_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "d"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	isContain := list.Contain(in)
	if !isContain {
		t.Log("the value not in LinkedList")
	} else {
		t.Error("function Contain has bug")
	}
}

func TestLinkedList_Contain_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	in := 1
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	isContain := list.Contain(in)
	if !isContain {
		t.Log("the value not in LinkedList")
	} else {
		t.Error("function Contain has bug")
	}
}

func TestLinkedList_Get(t *testing.T) {
	defer list.Clear()
	target := "b"
	index := 2
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_Get_With_Index_Less_Than_Zero(t *testing.T) {
	defer list.Clear()
	index := -1
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	_, err := list.Get(index)
	if err != nil {
		if errors.Is(err, errors2.InvalidIndexError) {
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

func TestLinkedList_Get_With_Index_Out_Of_LinkedList_Range(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	index := len(elements) + 1
	addValueToLinkedList(elements)
	_, err := list.Get(index)
	if err != nil {
		if errors.Is(err, errors2.InvalidIndexError) {
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

func TestLinkedList_HasCycle(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	list.head.next.next.next = list.head.next
	hasCycle := list.HasCycle()
	if hasCycle {
		t.Log("LinkedList has cycle")
	} else {
		t.Error("LinkedList not has cycle")
	}
}

func TestLinkedList_HasCycle_With_Not_Has_Cycle(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	hasCycle := list.HasCycle()
	if !hasCycle {
		t.Log("LinkedList not has cycle")
	} else {
		t.Error("function HasCycle has bug")
	}
}

func TestLinkedList_HasCycle_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	_ = list.Add("a")
	hasCycle := list.HasCycle()
	if !hasCycle {
		t.Log("LinkedList not has cycle")
	} else {
		t.Error("function HasCycle has bug")
	}
}

func TestLinkedList_HasCycle_With_LinkedList_Is_Empty(t *testing.T) {
	hasCycle := l.HasCycle()
	if !hasCycle {
		t.Log("LinkedList not has cycle")
	} else {
		t.Error("function HasCycle has bug")
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	defer list.Clear()
	index := 2
	target := "b"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_IndexOf_With_The_Value_Not_In_LinkedList(t *testing.T) {
	defer list.Clear()
	target := "d"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	_, err := list.IndexOf(target)
	if err != nil {
		if errors.Is(err, errors2.ValueNotExistError) {
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

func TestLinkedList_IndexOf_With_LinkedList_Is_Empty(t *testing.T) {
	defer l.Clear()
	_, err := l.IndexOf(1)
	if err != nil {
		if errors.Is(err, errors2.LinkedListIsEmptyError) {
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

func TestLinkedList_IndexOf_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	_, err := list.IndexOf(1)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function IndexOf has bug")
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_InsertAfter_With_Not_Existed_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	node := NewNode("A")
	if err := list.InsertAfter(node, in); err != nil {
		if errors.Is(err, errors2.NotExistError) {
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

func TestLinkedList_InsertAfter_With_Empty_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	var node *Node
	if err := list.InsertAfter(node, in); err != nil {
		if errors.Is(err, errors2.InputNodeIsEmptyError) {
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

func TestLinkedList_InsertAfter_With_Empty_LinkedList(t *testing.T) {
	defer l.Clear()
	node := NewNode(1)
	in := 2
	if err := l.InsertAfter(node, in); err != nil {
		if errors.Is(err, errors2.LinkedListIsEmptyError) {
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

func TestLinkedList_InsertAfter_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	_ = list.Add("a")
	node := list.Head()
	in := 2
	if err := list.InsertAfter(node, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func TestLinkedList_InsertBefore(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_InsertBefore_With_Head(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_InsertBefore_With_Not_Existed_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	node := NewNode("A")
	if err := list.InsertBefore(node, in); err != nil {
		if errors.Is(err, errors2.NotExistError) {
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

func TestLinkedList_InsertBefore_With_Empty_Node(t *testing.T) {
	defer list.Clear()
	in := "a"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	var node *Node
	if err := list.InsertBefore(node, in); err != nil {
		if errors.Is(err, errors2.InputNodeIsEmptyError) {
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

func TestLinkedList_InsertBefore_With_Empty_LinkedList(t *testing.T) {
	defer l.Clear()
	node := NewNode(1)
	in := 2
	if err := l.InsertBefore(node, in); err != nil {
		if errors.Is(err, errors2.LinkedListIsEmptyError) {
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

func TestLinkedList_InsertBefore_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	_ = list.Add("a")
	node := list.Head()
	in := 1
	if err := list.InsertBefore(node, in); err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function InsertAfter has bug")
	}
}

func TestLinkedList_LastIndexOf(t *testing.T) {
	defer list.Clear()
	ti := 5
	in := "b"
	elements := []string{"a", in, "c", "d", in}
	addValueToLinkedList(elements)
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

func TestLinkedList_LastIndexOf_With_Value_Not_In_LinkedList(t *testing.T) {
	defer list.Clear()
	in := "d"
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	_, err := list.LastIndexOf(in)
	if err != nil {
		if errors.Is(err, errors2.ValueNotExistError) {
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

func TestLinkedList_LastIndexOf_With_LinkedList_Is_Empty(t *testing.T) {
	defer l.Clear()
	in := 1
	_, err := l.LastIndexOf(in)
	if err != nil {
		if errors.Is(err, errors2.LinkedListIsEmptyError) {
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

func TestLinkedList_LastIndexOf_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	_ = list.Add("a")
	in := 1
	_, err := list.LastIndexOf(in)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function LastIndexOf has bug")
	}
}

func TestLinkedList_Len(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	if list.Len() == len(elements) {
		t.Log("get LinkedList len is success")
	} else {
		t.Error("get LinkedList len is failed")
	}
}

func TestLinkedList_Len_With_LinkedList_Is_Empty(t *testing.T) {
	if list.Len() == 0 {
		t.Log("get LinkedList len is success")
	} else {
		t.Error("get LinkedList len is failed")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	defer list.Clear()
	index := 2
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_Remove_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	index := 1
	in := "a"
	_ = list.Add(in)
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

func TestLinkedList_Remove_With_LinkedList_Is_Empty(t *testing.T) {
	defer list.Clear()
	node := NewNode("a")
	if err := list.Remove(node); err != nil {
		if errors.Is(err, errors2.LinkedListIsEmptyError) {
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

func TestLinkedList_Remove_With_Node_Not_Existed(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	node := NewNode("a")
	if err := list.Remove(node); err != nil {
		if errors.Is(err, errors2.NotExistError) {
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

func TestLinkedList_Remove_With_Node_Is_Empty(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	var node *Node
	if err := list.Remove(node); err != nil {
		if errors.Is(err, errors2.InputNodeIsEmptyError) {
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

func TestLinkedList_RemoveOf(t *testing.T) {
	defer list.Clear()
	index := 2
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_RemoveOf_With_LinkedList_Only_One_Element(t *testing.T) {
	defer list.Clear()
	index := 1
	in := "a"
	_ = list.Add(in)
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

func TestLinkedList_RemoveOf_With_Index_Not_In_Range(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	index := len(elements) + 1
	addValueToLinkedList(elements)
	_, err := list.RemoveOf(index)
	if err != nil {
		if errors.Is(err, errors2.InvalidIndexError) {
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

func TestLinkedList_Reverse(t *testing.T) {
	defer list.Clear()
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
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

func TestLinkedList_Set(t *testing.T) {
	defer list.Clear()
	target := "a"
	index := 2
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	old, err := list.Set(index, target)
	if err != nil {
		t.Errorf("set value has error! error: %v \n", err)
		return
	}
	node, _ := list.Get(index)
	if old == "b" && node.Value() == target {
		t.Log("set value is success")
	} else {
		t.Error("set value is failed")
	}
}

func TestLinkedList_Set_With_Invalid_Index(t *testing.T) {
	defer list.Clear()
	target := "a"
	index := -1
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	_, err := list.Set(index, target)
	if err != nil {
		if errors.Is(err, errors2.InvalidIndexError) {
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

func TestLinkedList_Set_With_Different_Type_Value(t *testing.T) {
	defer list.Clear()
	target := 1
	index := 3
	elements := []string{"a", "b", "c"}
	addValueToLinkedList(elements)
	_, err := list.Set(index, target)
	if err != nil {
		if errors.Is(err, errors2.InvalidTypeError) {
			t.Logf("type must be same")
			return
		} else {
			t.Errorf("unkown error! error: %v ", err)
			return
		}
	} else {
		t.Error("function Set has bug")
	}
}

// initDoubleLinkedList ,init a DoubleLinkedList
func initDoubleLinkedList(typeOf string) *LinkedList {
	list, _ := NewLinkedList(typeOf)
	return list
}

// addValueToLinkedList , add value to LinkedList
func addValueToLinkedList(elements []string) {
	for _, e := range elements {
		_ = list.Add(e)
	}
}
