package main

import "sort"

// Binary Search

type TimeMap struct {
	dict map[string][]*val
}

type val struct {
	value     string
	timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		dict: map[string][]*val{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.dict[key] = append(this.dict[key], &val{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.dict[key]
	if len(arr) == 0 {
		return ""
	}

	i := sort.Search(len(arr), func(i int) bool { return arr[i].timestamp > timestamp }) - 1
	if i < 0 {
		return ""
	}
	return arr[i].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
