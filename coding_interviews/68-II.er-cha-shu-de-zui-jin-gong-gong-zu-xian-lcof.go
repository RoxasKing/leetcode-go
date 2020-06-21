package codinginterviews

/*
  给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

  百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

  说明:
    所有节点的值都是唯一的。
    p、q 为不同节点且均存在于给定的二叉树中。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
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

func lowestCommonAncestorII2(root, p, q *TreeNode) *TreeNode {
	var signal int
	queue := []*TreeNode{root}
	pathes := make(map[*TreeNode]*TreeNode)
	mark := func(node *TreeNode) {
		if node == p {
			signal |= 1
		} else if node == q {
			signal |= 2
		}
	}
	handleNode := func(father, son *TreeNode) {
		pathes[son] = father
		mark(son)
		queue = append(queue, son)
	}
	for len(queue) != 0 && signal != 3 {
		root, queue = queue[0], queue[1:]
		if root.Left != nil {
			handleNode(root, root.Left)
		}
		if root.Right != nil {
			handleNode(root, root.Right)
		}
	}
	exists := make(map[*TreeNode]struct{})
	for node := p; node != nil; node = pathes[node] {
		exists[node] = struct{}{}
	}
	for node := q; node != nil; node = pathes[node] {
		if _, ok := exists[node]; ok {
			return node
		}
	}
	return nil
}
