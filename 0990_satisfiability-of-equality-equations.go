package leetcode

/*
  给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。

  只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

  提示：
    1 <= equations.length <= 500
    equations[i].length == 4
    equations[i][0] 和 equations[i][3] 是小写字母
    equations[i][1] 要么是 '='，要么是 '!'
    equations[i][2] 是 '='

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/satisfiability-of-equality-equations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}
	for _, equation := range equations {
		if equation[1] == '=' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			union(parent, index1, index2)
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			index1 := int(equation[0] - 'a')
			index2 := int(equation[3] - 'a')
			if find(parent, index1) == find(parent, index2) {
				return false
			}
		}
	}
	return true
}

func union(parent []int, index1, index2 int) {
	parent[find(parent, index1)] = find(parent, index2)
}

func find(parent []int, index int) int {
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}
