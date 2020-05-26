package codinginterviews

import (
	"math/rand"
	"sort"
)

/*
  输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
  示例 1：
    输入：arr = [3,2,1], k = 2
    输出：[1,2] 或者 [2,1]
  示例 2：
    输入：arr = [0,1,2,1], k = 1
    输出：[0]
*/

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return nil
	}
	sort.Ints(arr)
	return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	out := make([]int, k)
	var index int
	var qSortK func(int, int, int)
	qSortK = func(l, r, k int) {
		if l == r {
			out[index] = arr[l]
			index++
			return
		}
		pivotIdx := rand.Intn(r+1-l) + l
		arr[r], arr[pivotIdx] = arr[pivotIdx], arr[r]
		m := l
		for i := l; i < r; i++ {
			if arr[i] < arr[r] {
				arr[i], arr[m] = arr[m], arr[i]
				m++
			}
		}
		arr[m], arr[r] = arr[r], arr[m]
		if m+1-l < k {
			copy(out[index:index+1+m+1-l], arr[l:m+1])
			index += m + 1 - l
			qSortK(m+1, r, k-(m+1-l))
		} else if m+1-l > k {
			qSortK(l, m, k)
		} else {
			copy(out[index:index+m+1-l], arr[l:m+1])
		}
	}
	qSortK(0, len(arr)-1, k)
	return out
}
