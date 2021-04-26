package main

/*
  Given a (0-indexed) integer array nums and two integers low and high, return the number of nice pairs.
  A nice pair is a pair (i, j) where 0 <= i < j < nums.length and low <= (nums[i] XOR nums[j]) <= high.

  Example 1:
    Input: nums = [1,4,2,7], low = 2, high = 6
    Output: 6
    Explanation: All nice pairs (i, j) are as follows:
        - (0, 1): nums[0] XOR nums[1] = 5
        - (0, 2): nums[0] XOR nums[2] = 3
        - (0, 3): nums[0] XOR nums[3] = 6
        - (1, 2): nums[1] XOR nums[2] = 6
        - (1, 3): nums[1] XOR nums[3] = 3
        - (2, 3): nums[2] XOR nums[3] = 5

  Example 2:
    Input: nums = [9,8,4,2,1], low = 5, high = 14
    Output: 8
    Explanation: All nice pairs (i, j) are as follows:
    ​​​​​    - (0, 2): nums[0] XOR nums[2] = 13
        - (0, 3): nums[0] XOR nums[3] = 11
        - (0, 4): nums[0] XOR nums[4] = 8
        - (1, 2): nums[1] XOR nums[2] = 12
        - (1, 3): nums[1] XOR nums[3] = 10
        - (1, 4): nums[1] XOR nums[4] = 9
        - (2, 3): nums[2] XOR nums[3] = 6
        - (2, 4): nums[2] XOR nums[4] = 5


  Constraints:
    1. 1 <= nums.length <= 2 * 10^4
    2. 1 <= nums[i] <= 2 * 10^4
    3. 1 <= low <= high <= 2 * 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-pairs-with-xor-in-a-range
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!!

// Trie + Bit Manipulation
func countPairs(nums []int, low int, high int) int {
	trie := &Trie{}
	out := 0
	for _, num := range nums {
		out += trie.Lower(num, high+1) - trie.Lower(num, low)
		trie.Insert(num)
	}
	return out
}

type Trie struct {
	child [2]*Trie
	count int
}

func (t *Trie) Insert(num int) {
	node := t
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if node.child[idx] == nil {
			node.child[idx] = &Trie{}
		}
		node = node.child[idx]
		node.count++
	}
}

func (t *Trie) Lower(num, low int) int {
	out := 0
	node := t
	for i := 31; i >= 0 && node != nil; i-- {
		lowIdx := (low >> i) & 1
		numIdx := (num >> i) & 1
		if lowIdx == 1 {
			if node.child[numIdx] != nil {
				out += node.child[numIdx].count
			}
			node = node.child[1-numIdx]
		} else {
			node = node.child[numIdx]
		}
	}
	return out
}
