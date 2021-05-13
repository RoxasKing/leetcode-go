package main

import "sort"

/*
  Create a timebased key-value store class TimeMap, that supports two operations.

  1. set(string key, string value, int timestamp)
    Stores the key and value, along with the given timestamp.

  2. get(string key, int timestamp)
    1. Returns a value such that set(key, value, timestamp_prev) was called previously, with timestamp_prev <= timestamp.
    2. If there are multiple such values, it returns the one with the largest timestamp_prev.
    3. If there are no values, it returns the empty string ("").

  Example 1:
    Input: inputs = ["TimeMap","set","get","get","set","get","get"], inputs = [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
    Output: [null,null,"bar","bar",null,"bar2","bar2"]
    Explanation:
      TimeMap kv;
      kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1
      kv.get("foo", 1);  // output "bar"
      kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"
      kv.set("foo", "bar2", 4);
      kv.get("foo", 4); // output "bar2"
      kv.get("foo", 5); //output "bar2"

  Example 2:
    Input: inputs = ["TimeMap","set","set","get","get","get","get","get"], inputs = [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
    Output: [null,null,null,"","high","high","low","low"]

  Note:
    1. All key/value strings are lowercase.
    2. All key/value strings have length in the range [1, 100]
    3. The timestamps for all TimeMap.set operations are strictly increasing.
    4. 1 <= timestamp <= 10^7
    5. TimeMap.set and TimeMap.get functions will be called a total of 120000 times (combined) per test case.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/time-based-key-value-store
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
