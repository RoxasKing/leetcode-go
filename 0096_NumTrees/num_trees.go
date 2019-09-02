package main

import "fmt"

func numTrees(n int) int {
	out := 1
	for i := 0; i < n; i++ {
		out = out * 2 * (2*i + 1) / (i + 2)
	}
	return out
}

func numTrees2(n int) int {
	save := make([]int, n+1)
	save[0], save[1] = 1, 1

	// i 为节点数
	for i := 2; i <= n; i++ {
		// j 为分段的位置，从第一个开始
		for j := 1; j <= i; j++ {
			// j-1 即为去除根节点，左子树的个数， i-j 即为去除根节点，右子树的个数
			save[i] += save[j-1] * save[i-j]
		}
	}
	return save[n]
}

func main() {
	fmt.Println(numTrees(18))
}
