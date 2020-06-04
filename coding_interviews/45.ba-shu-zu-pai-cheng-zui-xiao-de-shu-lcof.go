package codinginterviews

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

/*
  输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

  提示:
    0 < nums.length <= 100

  说明:
    输出结果可能非常大，所以你需要返回一个字符串而不是整数
    拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Using stdlib sort package
func minNumber(nums []int) string {
	strs := convertIntSliceToStringSlice(nums)
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})
	return strings.Join(strs, "")
}

func convertIntSliceToStringSlice(nums []int) []string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	return strs
}

// Not use stdlib sort package
func minNumber2(nums []int) string {
	strs := convertIntSliceToStringSlice(nums)
	quickSortForminNuber(strs, 0, len(strs)-1)
	return strings.Join(strs, "")
}

func quickSortForminNuber(strs []string, l, r int) {
	if l >= r {
		return
	}
	pivotIndex := rand.Intn(r+1-l) + l
	strs[r], strs[pivotIndex] = strs[pivotIndex], strs[r]
	index := l
	for i := l; i < r; i++ {
		if strs[i]+strs[r] < strs[r]+strs[i] {
			if strs[i] != strs[index] {
				strs[i], strs[index] = strs[index], strs[i]
			}
			index++
		}
	}
	strs[index], strs[r] = strs[r], strs[index]
	quickSortForminNuber(strs, l, index-1)
	quickSortForminNuber(strs, index+1, r)
}
