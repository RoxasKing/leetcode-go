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
