package main

import (
	"math/rand"
)

/*
  输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

  示例 1：
    输入：arr = [3,2,1], k = 2
    输出：[1,2] 或者 [2,1]

  示例 2：
    输入：arr = [0,1,2,1], k = 1
    输出：[0]

  限制：
    1. 0 <= k <= arr.length <= 10000
    2. 0 <= arr[i] <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Quick Sort
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	qSort(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func qSort(arr []int, k, l, r int) {
	if l == r {
		return
	}
	pivotIdx := l + rand.Intn(r+1-l)
	arr[r], arr[pivotIdx] = arr[pivotIdx], arr[r]
	m := l
	for i := l; i < r; i++ {
		if arr[i] < arr[r] {
			arr[i], arr[m] = arr[m], arr[i]
			m++
		}
	}
	arr[m], arr[r] = arr[r], arr[m]
	if m > k-1 {
		qSort(arr, k, l, m-1)
	} else if m < k-1 {
		qSort(arr, k, m+1, r)
	}
}
