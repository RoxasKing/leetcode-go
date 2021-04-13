package main

/*
  力扣决定给一个刷题团队发LeetCoin作为奖励。同时，为了监控给大家发了多少LeetCoin，力扣有时候也会进行查询。

  该刷题团队的管理模式可以用一棵树表示：
    1. 团队只有一个负责人，编号为1。除了该负责人外，每个人有且仅有一个领导（负责人没有领导）；
    2. 不存在循环管理的情况，如A管理B，B管理C，C管理A。

  力扣想进行的操作有以下三种：
    1. 给团队的一个成员（也可以是负责人）发一定数量的LeetCoin；
    2. 给团队的一个成员（也可以是负责人），以及他/她管理的所有人（即他/她的下属、他/她下属的下属，……），发一定数量的LeetCoin；
    3. 查询某一个成员（也可以是负责人），以及他/她管理的所有人被发到的LeetCoin之和。

  输入：
    1. N表示团队成员的个数（编号为1～N，负责人为1）；
    2. leadership是大小为(N - 1) * 2的二维数组，其中每个元素[a, b]代表b是a的下属；
    3. operations是一个长度为Q的二维数组，代表以时间排序的操作，格式如下：
      1. operations[i][0] = 1: 代表第一种操作，operations[i][1]代表成员的编号，operations[i][2]代表LeetCoin的数量；
      2. operations[i][0] = 2: 代表第二种操作，operations[i][1]代表成员的编号，operations[i][2]代表LeetCoin的数量；
      3. operations[i][0] = 3: 代表第三种操作，operations[i][1]代表成员的编号；

  输出：
    返回一个数组，数组里是每次查询的返回值（发LeetCoin的操作不需要任何返回值）。由于发的LeetCoin很多，请把每次查询的结果模1e9+7 (1000000007)。


  示例 1：
    输入：N = 6, leadership = [[1, 2], [1, 6], [2, 3], [2, 5], [1, 4]], operations = [[1, 1, 500], [2, 2, 50], [3, 1], [2, 6, 15], [3, 1]]
    输出：[650, 665]
    解释：团队的管理关系见下图。
    第一次查询时，每个成员得到的LeetCoin的数量分别为（按编号顺序）：500, 50, 50, 0, 50, 0;
    第二次查询时，每个成员得到的LeetCoin的数量分别为（按编号顺序）：500, 50, 50, 0, 50, 15.

  限制：
    1. 1 <= N <= 50000
    2. 1 <= Q <= 50000
    3. operations[i][0] != 3 时，1 <= operations[i][2] <= 5000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/coin-bonus
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Segment Tree
func bonus(n int, leadership [][]int, operations [][]int) []int {
	edges := make([][]int, n+1)
	for _, e := range leadership {
		edges[e[0]] = append(edges[e[0]], e[1])
	}

	idxMap := make([][]int, n+1)
	idx := 0
	build(edges, idxMap, 1, &idx)

	sum := make([]int, n*4)
	lazySum := make([]int, n*4)

	out := []int{}
	for _, op := range operations {
		switch op[0] {
		case 1:
			update(sum, lazySum, idxMap[op[1]][0], idxMap[op[1]][0], op[2], 0, n-1, 0)
		case 2:
			update(sum, lazySum, idxMap[op[1]][0], idxMap[op[1]][1], op[2], 0, n-1, 0)
		case 3:
			res := query(sum, lazySum, idxMap[op[1]][0], idxMap[op[1]][1], 0, n-1, 0)
			out = append(out, res)
		}
	}
	return out
}

func build(edges [][]int, idxMap [][]int, member int, idx *int) {
	idxMap[member] = []int{*idx, *idx}
	*idx++
	for _, n := range edges[member] {
		build(edges, idxMap, n, idx)
		idxMap[member][1] = idxMap[n][1]
	}
}

const mod = 1e9 + 7

func update(sum, lazySum []int, opL, opR, opVal int, l, r int, idx int) {
	pushDown(sum, lazySum, l, r, idx)
	if r < opL || opR < l {
		return
	}
	if opL <= l && r <= opR {
		lazySum[idx] = opVal
		pushDown(sum, lazySum, l, r, idx)
		return
	}
	mid := (l + r) >> 1
	update(sum, lazySum, opL, opR, opVal, l, mid, idx<<1+1)
	update(sum, lazySum, opL, opR, opVal, mid+1, r, idx<<1+2)
	sum[idx] = (sum[idx<<1+1] + sum[idx<<1+2]) % mod
}

func query(sum, lazySum []int, opL, opR int, l, r int, idx int) int {
	if r < opL || opR < l {
		return 0
	}
	pushDown(sum, lazySum, l, r, idx)
	if opL <= l && r <= opR {
		return sum[idx]
	}
	mid := (l + r) >> 1
	lSum := query(sum, lazySum, opL, opR, l, mid, idx<<1+1)
	rSum := query(sum, lazySum, opL, opR, mid+1, r, idx<<1+2)
	return (lSum + rSum) % mod
}

func pushDown(sum, lazySum []int, l, r int, idx int) {
	if lazySum[idx] == 0 {
		return
	}
	sum[idx] = (sum[idx] + lazySum[idx]*(r-l+1)) % mod
	if l < r {
		lazySum[idx<<1+1] = (lazySum[idx<<1+1] + lazySum[idx]) % mod
		lazySum[idx<<1+2] = (lazySum[idx<<1+2] + lazySum[idx]) % mod
	}
	lazySum[idx] = 0
}
