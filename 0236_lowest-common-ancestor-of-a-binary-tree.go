package leetcode

/*
  给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
  百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
  例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS + Hash
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var signal int
	mark := func(node *TreeNode) {
		if node == p {
			signal |= 1
		}
		if node == q {
			signal |= 2
		}
	}
	link := make(map[*TreeNode]*TreeNode)
	queue := []*TreeNode{root}
	for len(queue) != 0 && signal != 3 {
		if queue[0].Left != nil {
			link[queue[0].Left] = queue[0]
			mark(queue[0].Left)
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			link[queue[0].Right] = queue[0]
			mark(queue[0].Right)
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	exists := make(map[*TreeNode]struct{})
	for p != nil {
		exists[p] = struct{}{}
		p = link[p]
	}
	for q != nil {
		if _, ok := exists[q]; ok {
			return q
		}
		q = link[q]
	}
	return nil
}

// DFS + Recursive
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		} else if node == p || node == q {
			return node
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l != nil && r != nil {
			return node
		} else if l == nil {
			return r
		}
		return l
	}
	return dfs(root)
}
