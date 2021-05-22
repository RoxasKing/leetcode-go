package main

import "math/rand"

/*
  Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.

  Example 1:
    Input
      ["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
      [[[1, 2, 3]], [], [], [], [], []]
    Output
      [null, 1, 3, 2, 2, 3]
    Explanation
      Solution solution = new Solution([1, 2, 3]);
      solution.getRandom(); // return 1
      solution.getRandom(); // return 3
      solution.getRandom(); // return 2
      solution.getRandom(); // return 2
      solution.getRandom(); // return 3
      // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.

  Constraints:
    1. The number of nodes in the linked list will be in the range [1, 10^4].
    2. -10^4 <= Node.val <= 10^4
    3. At most 104 calls will be made to getRandom.

  Follow up:
    1. What if the linked list is extremely large and its length is unknown to you?
    2. Could you solve this efficiently without using extra space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/linked-list-random-node
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Reservoir Sampling

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	out := 0
	cnt := 0
	for ptr := this.head; ptr != nil; ptr = ptr.Next {
		cnt++
		i := rand.Intn(cnt) + 1
		if i == cnt {
			out = ptr.Val
		}
	}
	return out
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
