package main

/*
  Given an integer array nums, handle multiple queries of the following types:
    1. Update the value of an element in nums.
    2. Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

  Implement the NumArray class:
    1. NumArray(int[] nums) Initializes the object with the integer array nums.
    2. void update(int index, int val) Updates the value of nums[index] to be val.
    3. int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

  Example 1:
    Input
      ["NumArray", "sumRange", "update", "sumRange"]
      [[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
    Output
      [null, 9, null, 8]
    Explanation
      NumArray numArray = new NumArray([1, 3, 5]);
      numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
      numArray.update(1, 2);   // nums = [1, 2, 5]
      numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8

  Constraints:
    1. 1 <= nums.length <= 3 * 10^4
    2. -100 <= nums[i] <= 100
    3. 0 <= index < nums.length
    4. -100 <= val <= 100
    5. 0 <= left <= right < nums.length
    6. At most 3 * 10^4 calls will be made to update and sumRange.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/range-sum-query-mutable
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Segment Tree

type NumArray struct {
	f    []int
	s, t int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	f := make([]int, 4*n)
	build(f, 1, nums, 0, n-1)
	return NumArray{f: f, s: 0, t: n - 1}
}

func (this *NumArray) Update(index int, val int) {
	update(this.f, 1, this.s, this.t, index, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return query(this.f, 1, this.s, this.t, left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func build(f []int, i int, nums []int, s, t int) {
	if s == t {
		f[i] = nums[s]
		return
	}
	m := (s + t) >> 1
	build(f, i<<1, nums, s, m)
	build(f, i<<1+1, nums, m+1, t)
	f[i] = f[i<<1] + f[i<<1+1]
}

func update(f []int, i, s, t, l, r, v int) {
	if t < l || s > r {
		return
	}
	if l <= s && t <= r {
		f[i] = v
		return
	}
	m := (s + t) >> 1
	update(f, i<<1, s, m, l, r, v)
	update(f, i<<1+1, m+1, t, l, r, v)
	f[i] = f[i<<1] + f[i<<1+1]
}

func query(f []int, i, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		return f[i]
	}
	m := (s + t) >> 1
	return query(f, i<<1, s, m, l, r) + query(f, i<<1+1, m+1, t, l, r)
}
