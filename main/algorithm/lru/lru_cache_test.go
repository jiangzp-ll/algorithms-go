package lru

import "testing"

func TestLRUCache_Put_And_Get(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	v1 := lru.Get(1)
	if v1 == 1 {
		t.Log("put and get success")
	} else {
		t.Error("put and get failed")
	}
}

func TestLRUCache_Delete(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	val := lru.Get(1)
	if val == -1 {
		t.Log("delete last use element success")
	} else {
		t.Error("lru realize failed")
	}
}