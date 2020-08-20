package codinginterviews

/*
  给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

  百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

  说明:
    所有节点的值都是唯一的。
    p、q 为不同节点且均存在于给定的二叉搜索树中。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursion
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	var search func(*TreeNode) *TreeNode
	search = func(root *TreeNode) *TreeNode {
		if root.Val > q.Val {
			return search(root.Left)
		} else if root.Val < p.Val {
			return search(root.Right)
		}
		return root
	}
	return search(root)
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	var out *TreeNode
	for out == nil {
		if root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val {
			root = root.Right
		} else {
			out = root
		}
	}
	return out
}
