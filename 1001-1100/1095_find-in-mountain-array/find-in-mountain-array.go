package main

/*
  (This problem is an interactive problem.)

  You may recall that an array A is a mountain array if and only if:
      1. A.length >= 3
      2. There exists some i with 0 < i < A.length - 1 such that:
         1. A[0] < A[1] < ... A[i-1] < A[i]
         2. A[i] > A[i+1] > ... > A[A.length - 1]
  Given a mountain array mountainArr, return the minimum index such that mountainArr.get(index) == target.  If such an index doesn't exist, return -1.

  You can't access the mountain array directly.  You may only access the array using a MountainArray interface:
    1. MountainArray.get(k) returns the element of the array at index k (0-indexed).
    2. MountainArray.length() returns the length of the array.
  Submissions making more than 100 calls to MountainArray.get will be judged Wrong Answer.  Also, any solutions that attempt to circumvent the judge will result in disqualification.

  Example 1:
    Input: array = [1,2,3,4,5,3,1], target = 3
    Output: 2
    Explanation: 3 exists in the array, at index=2 and index=5. Return the minimum index, which is 2.

  Example 2:
    Input: array = [0,1,2,4,2,1], target = 3
    Output: -1
    Explanation: 3 does not exist in the array, so we return -1.

  Constraints:
    1. 3 <= mountain_arr.length() <= 10000
    2. 0 <= target <= 10^9
    3. 0 <= mountain_arr.get(index) <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-in-mountain-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	top := bSearchTop(mountainArr, l, r)
	L, R := l, r
	l, r = L, top
	for l <= r {
		m := (l + r) >> 1
		res := mountainArr.get(m)
		if res < target {
			l = m + 1
		} else if res > target {
			r = m - 1
		} else {
			return m
		}
	}
	l, r = top+1, R
	for l <= r {
		m := (l + r) >> 1
		res := mountainArr.get(m)
		if res > target {
			l = m + 1
		} else if res < target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

func bSearchTop(mountainArr *MountainArray, l, r int) int {
	for l+1 < r {
		m := (l + r) >> 1
		a := mountainArr.get(m - 1)
		b := mountainArr.get(m)
		c := mountainArr.get(m + 1)
		if a < b && b > c {
			return m
		}
		if a < b && b < c {
			l = m
		} else {
			r = m
		}
	}
	return -1
}

type MountainArray struct {
	nums []int
}

func Constructor(nums []int) *MountainArray   { return &MountainArray{nums: nums} }
func (this *MountainArray) get(index int) int { return this.nums[index] }
func (this *MountainArray) length() int       { return len(this.nums) }
