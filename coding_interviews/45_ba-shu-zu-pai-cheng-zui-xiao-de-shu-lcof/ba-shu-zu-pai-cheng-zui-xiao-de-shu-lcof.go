package main

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

/*
  输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

  示例 1:
    输入: [10,2]
    输出: "102"

  示例 2:
    输入: [3,30,34,5,9]
    输出: "3033459"

  提示:
    0 < nums.length <= 100

  说明:
    1. 输出结果可能非常大，所以你需要返回一个字符串而不是整数
    2. 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Quick Sort
func minNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool { return strs[i]+strs[j] < strs[j]+strs[i] })
	return strings.Join(strs, "")
}

func minNumber2(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	qSort(strs, 0, n-1)
	return strings.Join(strs, "")
}

func qSort(strs []string, l, r int) {
	if l >= r {
		return
	}
	pivotIdx := l + rand.Intn(r+1-l)
	strs[pivotIdx], strs[r] = strs[r], strs[pivotIdx]
	i := l
	for j := l; j < r; j++ {
		if isSmaller(strs[j], strs[r]) {
			strs[j], strs[i] = strs[i], strs[j]
			i++
		}
	}
	strs[i], strs[r] = strs[r], strs[i]
	qSort(strs, l, i-1)
	qSort(strs, i+1, r)
}

func isSmaller(a, b string) bool {
	a, b = a+b, b+a
	return a < b
}
