package main

// Tags:
// Hash
// Morris Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
	n := len(trees)
	if n == 1 {
		return trees[0]
	}

	pVal := [5e4 + 1]int{}
	tMap := [5e4 + 1]*TreeNode{}
	for _, t := range trees {
		if t.Left != nil {
			if pVal[t.Val] == t.Left.Val || pVal[t.Left.Val] != 0 {
				return nil
			}
			pVal[t.Left.Val] = t.Val
		}
		if t.Right != nil {
			if pVal[t.Val] == t.Right.Val || pVal[t.Right.Val] != 0 {
				return nil
			}
			pVal[t.Right.Val] = t.Val
		}
		tMap[t.Val] = t
	}

	var root *TreeNode
	for _, t := range trees {
		if pVal[t.Val] == 0 {
			if root != nil {
				return nil
			}
			root = t
		}
	}

	if root == nil {
		return nil
	}

	var linked int
	wakled := [5e4 + 1]bool{}
	var prev *TreeNode
	for ptr := root; ptr != nil; {
		wakled[ptr.Val] = true
		if ptr.Left != nil {
			if tl := tMap[ptr.Left.Val]; tl != nil && ptr.Left != tl {
				linked++
				ptr.Left = tl
			}
			pre := ptr.Left
			for pre.Right != nil && pre.Right != ptr {
				if tr := tMap[pre.Right.Val]; tr != nil && pre.Right != tr {
					linked++
					pre.Right = tr
				}
				pre = pre.Right
			}
			if pre.Right != ptr {
				pre.Right = ptr
				ptr = ptr.Left
				continue
			}
			pre.Right = nil
		}
		if prev != nil && prev.Val >= ptr.Val {
			return nil
		}
		prev = ptr
		if ptr.Right != nil && !wakled[ptr.Right.Val] {
			if tr := tMap[ptr.Right.Val]; tr != nil && ptr.Right != tr {
				linked++
				ptr.Right = tr
			}
		}
		ptr = ptr.Right
	}

	if linked < n-1 {
		return nil
	}

	return root
}
