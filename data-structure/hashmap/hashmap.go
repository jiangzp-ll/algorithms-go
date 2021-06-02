package hashmap

import (
	"errors"
	"hash/crc32"
)

var (
	DefaultCapacity   = 16
	DefaultLoadFactor = 0.75
)

// HashMap ,like Java HashMap
type HashMap struct {
	data       []*Entry
	len        int
	loadFactor float64
}

// Entry ,store key and value
type Entry struct {
	Key   string
	Value interface{}
	Next  *Entry
}

// NewHashMap ,init a HashMap with default
func NewHashMap() *HashMap {
	return &HashMap{
		data:       make([]*Entry, DefaultCapacity),
		len:        0,
		loadFactor: DefaultLoadFactor,
	}
}

// NewHashMapWithCap ,init a HashMap with the capacity is input value
func NewHashMapWithCap(c int) (*HashMap, error) {
	if c < 0 {
		return nil, errors.New("illegal initial capacity")
	}
	return &HashMap{
		data:       make([]*Entry, c),
		len:        0,
		loadFactor: DefaultLoadFactor,
	}, nil
}

// Clear ,clear the HashMap
func (h *HashMap) Clear() {
	h.data = make([]*Entry, len(h.data))
	h.len = 0
}

// hash ,get the position of the key in the array of HashMap
func (h *HashMap) hash(key string) int64 {
	if key == "" {
		return 0
	}
	hc := hashCode(key)
	return (hc ^ (hc >> 16)) % int64(len(h.data))
}

// hashCode ,get the hash code of a string
func hashCode(key string) int64 {
	v := int64(crc32.ChecksumIEEE([]byte(key)))
	if v >= 0 {
		return v
	}
	if -v > 0 {
		return -v
	}
	// v == MinInt
	return 0
}
