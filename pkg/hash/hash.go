package hash

import "hash/crc32"

// Hash ,get the position of the key in the array
func Hash(length int, key string) int64 {
	if key == "" {
		return 0
	}
	hc := hashCode(key)
	return (hc ^ (hc >> 16)) % int64(length)
}

// hashCode ,get the Hash code of a string
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
