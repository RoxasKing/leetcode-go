package main

// Tags:
// Morris Traversal
func findMode(root *TreeNode) []int {
	dict := make(map[int]int)
	var out []int
	cur, count := -1<<31, 0
	handle := func() {
		if count == 0 {
			return
		}
		if len(out) != 0 { // 当前结果集有值
			if dict[out[0]] == count { // 当前统计数字的出现频率和结果集里的一样，添加
				out = append(out, cur)
			} else if dict[out[0]] < count { // 当前统计数字的出现频率大于结果集里的，清空现有结果集并添加当前值
				out = []int{cur}
			}
		} else {
			out = []int{cur}
		}
		dict[cur] = count
	}
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
		}
		if cur == node.Val {
			count++
		} else {
			handle()
			cur, count = node.Val, 1
		}
		node = node.Right
	}
	handle()
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
