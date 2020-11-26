package lru

import (
	"github.com/zepeng-jiang/go-basic-demo/main/algorithm/lru"
	"testing"
)

func TestLRUCache_Put_And_Get(t *testing.T) {
	l := lru.NewLRUCache(2)
	l.Put(1, 1)
	l.Put(2, 2)
	v1 := l.Get(1)
	if v1 == 1 {
		t.Log("put and get success")
	} else {
		t.Error("put and get failed")
	}
}

func TestLRUCache_Delete(t *testing.T) {
	l := lru.NewLRUCache(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	val := l.Get(1)
	if val == -1 {
		t.Log("delete last use element success")
	} else {
		t.Error("l realize failed")
	}
}