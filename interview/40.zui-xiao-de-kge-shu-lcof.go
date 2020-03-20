package interview

import "math/rand"

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
	help := func(lastIndex int) {
		for i := lastIndex / 2; i >= 0; i-- {
			son := i*2 + 1
			if son > lastIndex {
				continue
			}
			if son < lastIndex && arr[son+1] < arr[son] {
				son++
			}
			if arr[son] < arr[i] {
				arr[son], arr[i] = arr[i], arr[son]
			}
		}
	}
	for i := 0; i < k; i++ {
		help(len(arr) - 1 - i)
		arr[0], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[0]
	}
	return arr[len(arr)-k:]
}

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	out := make([]int, 0, k)
	quickSortK(arr, k, &out)
	return out
}

func quickSortK(arr []int, k int, out *[]int) {
	if len(arr) <= 1 {
		*out = append(*out, arr[0])
		return
	}
	pivotIdx := rand.Intn(len(arr))
	arr[len(arr)-1], arr[pivotIdx] = arr[pivotIdx], arr[len(arr)-1]
	pivot, mid := arr[len(arr)-1], 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			arr[i], arr[mid] = arr[mid], arr[i]
			mid++
		}
	}
	arr[mid], arr[len(arr)-1] = arr[len(arr)-1], arr[mid]
	if mid < k-1 {
		*out = append(*out, arr[:mid+1]...)
		quickSortK(arr[mid+1:], k-1-mid, out)
	} else if mid > k-1 {
		quickSortK(arr[:mid+1], k, out)
	} else {
		*out = append(*out, arr[:mid+1]...)
	}
}
