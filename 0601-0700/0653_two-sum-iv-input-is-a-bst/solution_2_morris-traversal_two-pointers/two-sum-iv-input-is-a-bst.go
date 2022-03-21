package main

// Difficulty:
// Easy

// Tags:
// Morris Traversal
// Two Pointers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr = [10001]int{}

func findTarget(root *TreeNode, k int) bool {
	tail := -1
	for ptr := root; ptr != nil; {
		if ptr.Left != nil {
			pre := ptr.Left
			for ; pre.Right != nil && pre.Right != ptr; pre = pre.Right {
			}
			if pre.Right != ptr {
				pre.Right = ptr
				ptr = ptr.Left
				continue
			}
			pre.Right = nil
		}
		tail++
		arr[tail] = ptr.Val
		ptr = ptr.Right
	}
	for l, r := 0, tail; l < r; {
		sum := arr[l] + arr[r]
		if sum < k {
			l++
		} else if sum > k {
			l--
		} else {
			return true
		}
	}
	return false
}
