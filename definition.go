package leetcode

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Max ...
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min ...
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs ...
func Abs(n int) int {
	if n <= 0 {
		return 0 - n
	}
	return n
}
