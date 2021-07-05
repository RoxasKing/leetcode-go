package main

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
