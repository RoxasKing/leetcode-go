package main

// Tags:
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
	n := mountainArr.length()
	top := bSearchTop(mountainArr, 0, n-1)
	if res := bSearch1(target, mountainArr, 0, top); res != -1 {
		return res
	}
	return bSearch2(target, mountainArr, top+1, n-1)
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

func bSearch1(target int, mountainArr *MountainArray, l, r int) int {
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
	return -1
}

func bSearch2(target int, mountainArr *MountainArray, l, r int) int {
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

type MountainArray struct {
	nums []int
}

func Constructor(nums []int) *MountainArray   { return &MountainArray{nums: nums} }
func (this *MountainArray) get(index int) int { return this.nums[index] }
func (this *MountainArray) length() int       { return len(this.nums) }
