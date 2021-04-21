package main

/*
  Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 ≤ i ≤ j < n.

  Follow up: Could you do this in O(n) runtime?

  Example 1:
    Input: nums = [3,10,5,25,2,8]
    Output: 28
    Explanation: The maximum result is 5 XOR 25 = 28.

  Example 2:
    Input: nums = [0]
    Output: 0

  Example 3:
    Input: nums = [2,4]
    Output: 6

  Example 4:
    Input: nums = [8,10,2]
    Output: 10

  Example 5:
    Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
    Output: 127

  Constraints:
    1. 1 <= nums.length <= 2 * 10^4
    2. 0 <= nums[i] <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie + Queue
func findMaximumXOR(nums []int) int {
	trie := NewTrie()
	for _, num := range nums {
		trie.Insert(num)
	}

	out := -1 << 31
	q := []*triePair{{a: trie, b: trie, num: 0}}
	for len(q) > 0 {
		tp := q[0]
		q = q[1:]
		a, b, num := tp.a, tp.b, tp.num
		if a.isEnd && b.isEnd {
			out = Max(out, num)
			continue
		}

		pushed := false
		if a.child[0] != nil && b.child[1] != nil {
			q = append(q, &triePair{a: a.child[0], b: b.child[1], num: num<<1 + 1})
			pushed = true
		}
		if a.child[1] != nil && b.child[0] != nil {
			q = append(q, &triePair{a: a.child[1], b: b.child[0], num: num<<1 + 1})
			pushed = true
		}
		if pushed {
			continue
		}

		if a.child[0] != nil {
			q = append(q, &triePair{a: a.child[0], b: b.child[0], num: num << 1})
		} else {
			q = append(q, &triePair{a: a.child[1], b: b.child[1], num: num << 1})
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type triePair struct {
	a   *Trie
	b   *Trie
	num int
}

type Trie struct {
	child [2]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(num int) {
	node := t
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if node.child[idx] == nil {
			node.child[idx] = NewTrie()
		}
		node.isEnd = false
		node = node.child[idx]
	}
	node.isEnd = true
}
