package main

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

// Important!

// Trie
// Online Algorithm
// Bit Manipulation

func maximizeXor(nums []int, queries [][]int) []int {
	t := MakeTrie()
	for _, num := range nums {
		t.Insert(num)
	}

	n := len(queries)
	out := make([]int, n)

	for i, q := range queries {
		out[i] = t.Query(q[0], q[1])
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Trie struct {
	child [2]*Trie
	min   int
}

func MakeTrie() *Trie {
	return &Trie{min: 1<<31 - 1}
}

func (t *Trie) Insert(num int) {
	for b := 31; b >= 0; b-- {
		i := (num >> b) & 1
		if t.child[i] == nil {
			t.child[i] = MakeTrie()
		}
		t = t.child[i]
		t.min = Min(t.min, num)
	}
}

func (t *Trie) Query(x, m int) int {
	out := 0
	for b := 31; b >= 0; b-- {
		i := (x >> b) & 1
		if t.child[1-i] != nil && t.child[1-i].min <= m {
			out |= 1 << b
			t = t.child[1-i]
		} else if t.child[i].min <= m {
			t = t.child[i]
		} else {
			return -1
		}
	}
	return out
}
