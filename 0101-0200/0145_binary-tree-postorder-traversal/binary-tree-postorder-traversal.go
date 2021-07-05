package main

// Tags:
// DFS
func postorderTraversal(root *TreeNode) []int {
	var out []int
	dfs(root, &out)
	return out
}

func dfs(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, out)
	dfs(node.Right, out)
	*out = append(*out, node.Val)
}

// Stack
func postorderTraversal2(root *TreeNode) []int {
	out := []int{}
	s := TreeStack{}
	prev := &TreeNode{}
	for root != nil || s.Len() > 0 {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		top := s.Top()
		if top.Right != nil && top.Right != prev {
			root = top.Right
		} else {
			s.Pop()
			out = append(out, top.Val)
			prev = top
		}
	}
	return out
}

type TreeStack []*TreeNode

func (s TreeStack) Len() int          { return len(s) }
func (s TreeStack) Top() *TreeNode    { return s[s.Len()-1] }
func (s *TreeStack) Push(x *TreeNode) { *s = append(*s, x) }
func (s *TreeStack) Pop() *TreeNode {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

// Morris Traversal
func postorderTraversal3(root *TreeNode) []int {
	out := []int{}
	for ptr := root; ptr != nil; {
		if pre := ptr.Left; pre != nil {
			for pre.Right != nil && pre.Right != ptr {
				pre = pre.Right
			}
			if pre.Right != ptr {
				pre.Right = ptr
				ptr = ptr.Left
				continue
			}
			pre.Right = nil
			addPath(ptr.Left, &out)
		}
		ptr = ptr.Right
	}
	addPath(root, &out)
	return out
}

func addPath(root *TreeNode, out *[]int) {
	size := len(*out)
	for ; root != nil; root = root.Right {
		*out = append(*out, root.Val)
	}
	reverse((*out)[size:])
}

func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n>>1; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
