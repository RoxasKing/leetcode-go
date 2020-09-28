package main

import (
	"container/heap"
	"sort"
)

/*
  城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。现在，假设您获得了城市风光照片（图A）上显示的所有建筑物的位置和高度，请编写一个程序以输出由这些建筑物形成的天际线（图B）。

  每个建筑物的几何信息用三元组 [Li，Ri，Hi] 表示，其中 Li 和 Ri 分别是第 i 座建筑物左右边缘的 x 坐标，Hi 是其高度。可以保证 0 ≤ Li, Ri ≤ INT_MAX, 0 < Hi ≤ INT_MAX 和 Ri - Li > 0。您可以假设所有建筑物都是在绝对平坦且高度为 0 的表面上的完美矩形。

  例如，图A中所有建筑物的尺寸记录为：[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] 。

  输出是以 [ [x1,y1], [x2, y2], [x3, y3], ... ] 格式的“关键点”（图B中的红点）的列表，它们唯一地定义了天际线。关键点是水平线段的左端点。请注意，最右侧建筑物的最后一个关键点仅用于标记天际线的终点，并始终为零高度。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

  例如，图B中的天际线应该表示为：[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]。

  说明:
    任何输入列表中的建筑物数量保证在 [0, 10000] 范围内。
    输入列表已经按左 x 坐标 Li  进行升序排列。
    输出列表必须按 x 位排序。
    输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/the-skyline-problem
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority queue ( Max Stack )
func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([]int, 0, n*2)
	for _, building := range buildings {
		points = append(points, building[0], building[1]) // 将左右边缘坐标放入数组points中
	}
	sort.Ints(points)
	bh := buildingHeap{}                                   // 对于建筑物，构建以高度为关键字的最大堆
	heap.Push(&bh, building{rBound: 1<<63 - 1, height: 0}) // 哨兵节点，保证最小高度为0
	var index int
	var out [][]int
	for _, point := range points {
		for index < n && buildings[index][0] == point { // 遍历建筑物记录，将左边界等于当前点的建筑物入队
			heap.Push(&bh, building{
				rBound: buildings[index][1],
				height: buildings[index][2],
			})
			index++
		}
		for bh[0].rBound <= point { // 将右边界小于等于当前点的建筑物出队
			heap.Pop(&bh)
		}
		if len(out) == 0 || out[len(out)-1][1] != bh[0].height { // 结果集最后一个高度和当前堆的最大高度不同，将当前点加入结果集
			out = append(out, []int{point, bh[0].height})
		}
	}
	return out
}

type building struct{ rBound, height int }
type buildingHeap []building

func (bh buildingHeap) Len() int            { return len(bh) }
func (bh buildingHeap) Less(i, j int) bool  { return bh[i].height > bh[j].height }
func (bh buildingHeap) Swap(i, j int)       { bh[i], bh[j] = bh[j], bh[i] }
func (bh *buildingHeap) Push(x interface{}) { *bh = append(*bh, x.(building)) }
func (bh *buildingHeap) Pop() interface{} {
	out := (*bh)[bh.Len()-1]
	*bh = (*bh)[:bh.Len()-1]
	return out
}
