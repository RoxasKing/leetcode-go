package main

/*
  小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves， 字符串 leaves 仅包含小写字符 r 和 y， 其中字符 r 表示一片红叶，字符 y 表示一片黄叶。
  出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。每部分树叶数量可以不相等，但均需大于等于 1。每次调整操作，小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。

  示例 1：
    输入：leaves = "rrryyyrryyyrr"
    输出：2
    解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"

  示例 2：
    输入：leaves = "ryr"
    输出：0
    解释：已符合要求，不需要额外操作

  提示：
    1. 3 <= leaves.length <= 10^5
    2. leaves 中只包含字符 'r' 和字符 'y'

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/UlBDOe
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// TODO

// Dynamic Programming
func minimumOperations(leaves string) int {
	n := len(leaves)
	g := -1
	if leaves[0] == 'y' {
		g = 1
	}
	gmin := g
	min := 1<<31 - 1
	for i := 1; i < n-1; i++ {
		g += 2*isYellow(leaves[i] == 'y') - 1
		min = Min(min, gmin-g)
		gmin = Min(gmin, g)
	}
	if leaves[n-1] == 'y' {
		g++
	}
	return min + (g+n)/2
}

func isYellow(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
