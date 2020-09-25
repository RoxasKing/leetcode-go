package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
  给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

  说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
