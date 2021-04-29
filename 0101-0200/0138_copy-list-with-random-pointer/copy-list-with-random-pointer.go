package main

/*
  A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

  Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

  For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

  Return the head of the copied linked list.

  The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:
    1. val: an integer representing Node.val
    2. random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
  Your code will only be given the head of the original linked list.

  Example 1:
    Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
    Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

  Example 2:
    Input: head = [[1,1],[2,1]]
    Output: [[1,1],[2,1]]

  Example 3:
    Input: head = [[3,null],[3,0],[3,null]]
    Output: [[3,null],[3,0],[3,null]]

  Example 4:
    Input: head = []
    Output: []
    Explanation: The given linked list is empty (null pointer), so return null.

  Constraints:
    1. 0 <= n <= 1000
    2. -10000 <= Node.val <= 10000
    3. Node.random is null or is pointing to some node in the linked list.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	for node := head; node != nil; node = node.Next {
		temp := &Node{Val: node.Val}
		temp.Next = node.Next
		node.Next = temp
		node = temp
	}

	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	outPre := &Node{}
	ptr := outPre
	for node := head; node != nil; node = node.Next {
		ptr.Next = node.Next
		ptr = ptr.Next
		node.Next = node.Next.Next
	}
	ptr.Next = nil
	return outPre.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
