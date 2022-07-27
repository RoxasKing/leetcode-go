package main

// Difficulty:
// Medium

// Tags:
// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
	q    []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	for ; len(q) > 0; q = q[1:] {
		t := q[0]
		if t.Left == nil {
			break
		}
		q = append(q, t.Left)
		if t.Right == nil {
			break
		}
		q = append(q, t.Right)
	}
	return CBTInserter{root, q}
}

func (this *CBTInserter) Insert(val int) int {
	t := this.q[0]
	if t.Left == nil {
		t.Left = &TreeNode{Val: val}
		this.q = append(this.q, t.Left)
		return t.Val
	}
	t.Right = &TreeNode{Val: val}
	this.q = append(this.q, t.Right)
	this.q = this.q[1:]
	return t.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
