package main

import "sort"

/*
  You are given an array nums consisting of non-negative integers. You are also given a queries array, where queries[i] = [xi, mi].

  The answer to the i^th query is the maximum bitwise XOR value of xi and any element of nums that does not exceed mi. In other words, the answer is max(nums[j] XOR xi) for all j such that nums[j] <= mi. If all elements in nums are larger than mi, then the answer is -1.

  Return an integer array answer where answer.length == queries.length and answer[i] is the answer to the i^th query.

  Example 1:
    Input: nums = [0,1,2,3,4], queries = [[3,1],[1,3],[5,6]]
    Output: [3,3,7]
    Explanation:
      1) 0 and 1 are the only two integers not greater than 1. 0 XOR 3 = 3 and 1 XOR 3 = 2. The larger of the two is 3.
      2) 1 XOR 2 = 3.
      3) 5 XOR 2 = 7.

  Example 2:
    Input: nums = [5,2,4,6,6,3], queries = [[12,4],[8,1],[6,3]]
    Output: [15,-1,5]

  Constraints:
    1. 1 <= nums.length, queries.length <= 10^5
    2. queries[i].length == 2
    3. 0 <= nums[j], xi, mi <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-xor-with-an-element-from-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie + Bit Manipulation
func maximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	n := len(queries)
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool { return queries[idxs[i]][1] < queries[idxs[j]][1] })
	t := &Trie{}
	out := make([]int, n)
	for _, i := range idxs {
		x, m := queries[i][0], queries[i][1]
		for len(nums) > 0 && nums[0] <= m {
			t.Insert(nums[0])
			nums = nums[1:]
		}
		out[i] = t.Query(x)
	}
	return out
}

type Trie struct {
	child [2]*Trie
}

func (t *Trie) Insert(num int) {
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
}

func (t *Trie) Query(num int) int {
	out := 0
	for i := 31; i >= 0; i-- {
		if t == nil {
			return -1
		}
		idx := (num >> i) & 1
		if t.child[1-idx] != nil {
			out |= 1 << i
			t = t.child[1-idx]
		} else {
			t = t.child[idx]
		}
	}
	return out
}
