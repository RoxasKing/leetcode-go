package main

import "math/rand"

/*
  打乱一个没有重复元素的数组。

  示例:
    // 以数字集合 1, 2 和 3 初始化数组。
    int[] nums = {1,2,3};
    Solution solution = new Solution(nums);
    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
    solution.shuffle();
    // 重设数组到它的初始状态[1,2,3]。
    solution.reset();
    // 随机返回数组[1,2,3]打乱后的结果。
    solution.shuffle();

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shuffle-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Solution struct {
	origin []int
	array  []int
}

func Constructor(nums []int) Solution {
	array := make([]int, len(nums))
	copy(array, nums)
	return Solution{
		origin: nums,
		array:  array,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.array = this.origin
	this.origin = make([]int, len(this.array))
	copy(this.origin, this.array)
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := range this.array {
		tmp := rand.Intn(len(this.array))
		this.array[i], this.array[tmp] = this.array[tmp], this.array[i]
	}
	return this.array
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
