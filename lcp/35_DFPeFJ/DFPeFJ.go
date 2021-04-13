package main

import "container/heap"

/*
  小明的电动车电量充满时可行驶距离为 cnt，每行驶 1 单位距离消耗 1 单位电量，且花费 1 单位时间。小明想选择电动车作为代步工具。地图上共有 N 个景点，景点编号为 0 ~ N-1。他将地图信息以 [城市 A 编号,城市 B 编号,两城市间距离] 格式整理在在二维数组 paths，表示城市 A、B 间存在双向通路。初始状态，电动车电量为 0。每个城市都设有充电桩，charge[i] 表示第 i 个城市每充 1 单位电量需要花费的单位时间。请返回小明最少需要花费多少单位时间从起点城市 start 抵达终点城市 end。

  示例 1：
    输入：paths = [[1,3,3],[3,2,1],[2,1,3],[0,1,4],[3,0,5]], cnt = 6, start = 1, end = 0, charge = [2,10,4,1]
    输出：43
    解释：最佳路线为：1->3->0。
      在城市 1 仅充 3 单位电至城市 3，然后在城市 3 充 5 单位电，行驶至城市 5。
      充电用时共 3*10 + 5*1= 35
      行驶用时 3 + 5 = 8，此时总用时最短 43。

  示例 2：
    输入：paths = [[0,4,2],[4,3,5],[3,0,5],[0,1,5],[3,2,4],[1,2,8]], cnt = 8, start = 0, end = 2, charge = [4,1,1,3,2]
    输出：38
    解释：最佳路线为：0->4->3->2。
      城市 0 充电 2 单位，行驶至城市 4 充电 8 单位，行驶至城市 3 充电 1 单位，最终行驶至城市 2。
      充电用时 4*2+2*8+3*1 = 27
      行驶用时 2+5+4 = 11，总用时最短 38。

  提示：
    1. 1 <= paths.length <= 200
    2. paths[i].length == 3
    3. 2 <= charge.length == n <= 100
    4. 0 <= path[i][0],path[i][1],start,end < n
    5. 1 <= cnt <= 100
    6. 1 <= path[i][2] <= cnt
    7. 1 <= charge[i] <= 100
    8. 题目保证所有城市相互可以到达

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/DFPeFJ
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dijkstra's algorithm + Priority Queue
func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	edges := make([][][2]int, n)
	for _, p := range paths {
		u, v, w := p[0], p[1], p[2]
		edges[u] = append(edges[u], [2]int{v, w})
		edges[v] = append(edges[v], [2]int{u, w})
	}

	pq := PriorityQueue{}
	heap.Push(&pq, [3]int{0, start, 0})

	minT := make([][]int, n) // 记录到达不同地点不同电量的最少耗时
	for i := range minT {
		minT[i] = make([]int, cnt+1)
		for j := range minT[i] {
			minT[i][j] = 1<<31 - 1
		}
	}
	minT[start][0] = 0

	for pq.Len() > 0 {
		e := heap.Pop(&pq).([3]int)
		t, u, c := e[0], e[1], e[2] // 耗时， 地点， 电量

		if t > minT[u][c] { // 大于已记录的最佳耗时，跳过
			continue
		}

		if u == end {
			return t
		}

		if c < cnt { // 充电
			newT := t + charge[u]
			if newT < minT[u][c+1] {
				minT[u][c+1] = newT
				heap.Push(&pq, [3]int{newT, u, c + 1})
			}
		}

		for _, next := range edges[u] {
			v, w := next[0], next[1]
			if c >= w && t+w < minT[v][c-w] { // 电量足够，可达，并且比记录的耗时低，更新记录，并入队
				minT[v][c-w] = t + w
				heap.Push(&pq, [3]int{t + w, v, c - w})
			}
		}
	}

	return -1
}

type PriorityQueue [][3]int // 耗时， 地点， 电量

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i][0] < pq[j][0] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([3]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
