package codinginterviews

/*
  请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

  提示：
    -10000 <= Node.val <= 10000
    Node.random 为空（null）或指向链表中的节点。
    节点数目不超过 1000 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func copyRandomList(head *Node) *Node {
	cloneNodes(head)
	cloneRadomLinks(head)
	return splitAndReorganize(head)
}

func cloneNodes(head *Node) {
	for n := head; n != nil; n = n.Next.Next {
		clone := &Node{Val: n.Val}
		clone.Next = n.Next
		n.Next = clone
	}
}

func cloneRadomLinks(head *Node) {
	for n := head; n != nil; n = n.Next.Next {
		if n.Random != nil {
			n.Next.Random = n.Random.Next
		}
	}
}

func splitAndReorganize(head *Node) *Node {
	pre := &Node{}
	cur := pre
	for n := head; n != nil; n = n.Next {
		cur.Next = n.Next
		cur = n.Next
		n.Next = n.Next.Next
	}
	return pre.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
