package main

/*
  In a row of trees, the i-th tree produces fruit with type tree[i].

  You start at any tree of your choice, then repeatedly perform the following steps:

  Add one piece of fruit from this tree to your baskets.  If you cannot, stop.
  Move to the next tree to the right of the current tree.  If there is no tree to the right, stop.
  Note that you do not have any choice after the initial choice of starting tree: you must perform step 1, then step 2, then back to step 1, then step 2, and so on until you stop.

  You have two baskets, and each basket can carry any quantity of fruit, but you want each basket to only carry one type of fruit each.

  What is the total amount of fruit you can collect with this procedure?

  Example 1:
    Input: [1,2,1]
    Output: 3
    Explanation: We can collect [1,2,1].

  Example 2:
    Input: [0,1,2,2]
    Output: 3
    Explanation: We can collect [1,2,2].
    If we started at the first tree, we would only collect [0, 1].

  Example 3:
    Input: [1,2,3,2,2]
    Output: 4
    Explanation: We can collect [2,3,2,2].
    If we started at the first tree, we would only collect [1, 2].

  Example 4:
    Input: [3,3,3,1,2,1,1,2,3,3,4]
    Output: 5
    Explanation: We can collect [1,2,1,1,2].
    If we started at the first tree or the eighth tree, we would only collect 4 fruits.

  Note:
    1 <= tree.length <= 40000
    0 <= tree[i] < tree.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fruit-into-baskets
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func totalFruit(tree []int) int {
	out := 0
	l1, l2, r1, r2 := 0, 0, 0, 0
	for i := range tree {
		if tree[i] == tree[l1] {
			l2 = i
		} else if tree[i] == tree[r1] {
			r2 = i
		} else if tree[l1] == tree[r1] {
			r1, r2 = i, i
		} else {
			out = Max(out, Max(l2, r2)+1-l1)
			l1, l2, r1, r2 = Min(l2+1, r2+1), Max(l2, r2), i, i
		}
	}
	out = Max(out, Max(l2, r2)+1-l1)
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
