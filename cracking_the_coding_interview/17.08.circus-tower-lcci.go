package crackingthecodingintervew

import (
	"math/rand"
)

/*
  有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

  提示：
    height.length == weight.length <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/circus-tower-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func bestSeqAtIndex(height []int, weight []int) int {
	var res []int
	qSortBestSeqAtIndex(height, weight, 0, len(height)-1)
	for i := range weight {
		l, r := 0, len(res)
		for l < r {
			m := l + (r-l)>>1
			if height[i] > height[res[m]] && weight[i] > weight[res[m]] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l < len(res) {
			res[l] = i
		} else {
			res = append(res, i)
		}
	}
	return len(res)
}

func qSortBestSeqAtIndex(a, b []int, l, r int) {
	if l >= r {
		return
	}
	pivotIndex := l + rand.Intn(r+1-l)
	a[pivotIndex], a[r] = a[r], a[pivotIndex]
	b[pivotIndex], b[r] = b[r], b[pivotIndex]
	index := l
	for i := l; i <= r; i++ {
		if a[i] < a[r] || a[i] == a[r] && b[i] > b[r] {
			a[i], a[index] = a[index], a[i]
			b[i], b[index] = b[index], b[i]
			index++
		}
	}
	a[index], a[r] = a[r], a[index]
	b[index], b[r] = b[r], b[index]
	qSortBestSeqAtIndex(a, b, l, index-1)
	qSortBestSeqAtIndex(a, b, index+1, r)
}
