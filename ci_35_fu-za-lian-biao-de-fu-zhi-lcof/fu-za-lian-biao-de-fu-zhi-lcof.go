package main

/*
  请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

  示例 1：
    输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
    输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

  示例 2：
    输入：head = [[1,1],[2,1]]
    输出：[[1,1],[2,1]]

  示例 3：
    输入：head = [[3,null],[3,0],[3,null]]
    输出：[[3,null],[3,0],[3,null]]

  示例 4：
    输入：head = []
    输出：[]
    解释：给定的链表为空（空指针），因此返回 null。

  提示：
    1. -10000 <= Node.val <= 10000
    2. Node.random 为空（null）或指向链表中的节点。
    3. 节点数目不超过 1000 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
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
	for n := head; n != nil; n = n.Next.Next {
		node := &Node{Val: n.Val, Next: n.Next}
		n.Next = node
	}
	for n := head; n != nil; n = n.Next.Next {
		if n.Random != nil {
			n.Next.Random = n.Random.Next
		}
	}
	outPre := &Node{}
	ptr := outPre
	for n := head; n != nil; n = n.Next {
		ptr.Next = n.Next
		ptr = ptr.Next
		n.Next = n.Next.Next
	}
	return outPre.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
