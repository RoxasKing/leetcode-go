package main

/*
  在一排树中，第 i 棵树产生 tree[i] 型的水果。
  你可以从你选择的任何树开始，然后重复执行以下步骤：

  把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
  移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
  请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

  你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
  用这个程序你能收集的水果总量是多少？

  提示：
    1 <= tree.length <= 40000
    0 <= tree[i] < tree.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fruit-into-baskets
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Double Pointer
func totalFruit(tree []int) int {
	var out, start, p1, p2 int
	for p2 < len(tree) && tree[p2] == tree[0] {
		p2++
	}
	if p2 == len(tree) {
		return len(tree)
	}
	p1 = p2 - 1
	for i := p2 + 1; i < len(tree); i++ {
		switch tree[i] {
		case tree[p1]:
			p1 = i
		case tree[p2]:
			p2 = i
		default:
			out = Max(out, Max(p1, p2)+1-start)
			start = Min(p1, p2) + 1
			p1, p2 = i-1, i
		}
	}
	out = Max(out, Max(p1, p2)+1-start)
	return out
}

// Sliding Window
func totalFruit2(tree []int) int {
	var out, l, r int
	count := make(map[int]int)
	for r = range tree {
		count[tree[r]]++
		for len(count) > 2 {
			count[tree[l]]--
			if count[tree[l]] == 0 {
				delete(count, tree[l])
			}
			l++
		}
		out = Max(out, r+1-l)
	}
	return out
}
