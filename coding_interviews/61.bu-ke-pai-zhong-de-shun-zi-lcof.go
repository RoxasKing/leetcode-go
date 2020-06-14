package codinginterviews

import "sort"

/*
  从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

  限制：
    数组长度为 5
    数组的数取值为 [0, 13] .

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isStraight(nums []int) bool {
	sort.Ints(nums)
	var index, countZero, countEmpty int
	for nums[index] == 0 {
		countZero++
		index++
	}
	for index < len(nums)-1 {
		if nums[index+1] == nums[index] {
			return false
		}
		countEmpty += (nums[index+1] - nums[index] - 1)
		index++
	}
	return countZero >= countEmpty
}
