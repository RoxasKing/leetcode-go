package leetcode

// ListNode : Linked list data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode : Binary tree data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Max : Returns the largest number in a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min : Returns the smallest  number in a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs : Returns the absolute value of n
func Abs(n int) int {
	if n <= 0 {
		return 0 - n
	}
	return n
}

// Gcd : Returns greatest common divisor of a and b
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
