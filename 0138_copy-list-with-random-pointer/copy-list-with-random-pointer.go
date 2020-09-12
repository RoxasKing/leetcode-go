package main

/*
  给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
  要求返回这个链表的 深拷贝。
  我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
  val：一个表示 Node.val 的整数。
  random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。

  提示：
    -10000 <= Node.val <= 10000
    Node.random 为空（null）或指向链表中的节点。
    节点数目不超过 1000 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	node := head
	for node != nil {
		tmp := &Node{Val: node.Val}
		tmp.Next = node.Next
		node.Next = tmp

		node = node.Next.Next
	}

	node = head
	for node != nil {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}

		node = node.Next.Next
	}

	out := head.Next
	ptr := out

	head.Next = head.Next.Next
	node = head.Next
	for node != nil {
		ptr.Next = node.Next
		ptr = ptr.Next

		node.Next = node.Next.Next
		node = node.Next
	}

	return out
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
