package leetcode

/*
  给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
  说明：解集不能包含重复的子集
*/

func subsets(nums []int) [][]int {
	var out [][]int
	var recur func(int, int, []int)
	recur = func(index, k int, array []int) {
		if len(array) == k {
			out = append(out, append(make([]int, 0, k), array...))
			return
		}
		for i := index; i < len(nums); i++ {
			array = append(array, nums[i])
			recur(i+1, k, array)
			array = array[:len(array)-1]
		}
	}
	for i := 0; i <= len(nums); i++ {
		recur(0, i, make([]int, 0, i))
	}
	return out
}
