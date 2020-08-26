package main

/*
  给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
    struct Node {
      int val;
      Node *left;
      Node *right;
      Node *next;
    }
  填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

  初始状态下，所有 next 指针都被设置为 NULL。

  提示：
    你只能使用常量级额外空间。
    使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Recursion
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var bfs func(*Node)
	bfs = func(node *Node) {
		if node.Left == nil && node.Right == nil {
			return
		}
		nodeL, nodeR := node.Left, node.Right
		for nodeL != nil && nodeR != nil {
			nodeL.Next = nodeR
			nodeL, nodeR = nodeL.Right, nodeR.Left
		}
		bfs(node.Left)
		bfs(node.Right)
	}
	bfs(root)
	return root
}

// Morris Traversal
func connect2(root *Node) *Node {
	node := root
	for node != nil {
		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}
			if pre.Right != node {
				pre.Right = node
				node = node.Left
				continue
			}
			pre.Right = nil
			nodeL, nodeR := node.Left, node.Right
			for nodeL != nil && nodeR != nil {
				nodeL.Next = nodeR
				nodeL, nodeR = nodeL.Right, nodeR.Left
			}
		}
		node = node.Right
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
