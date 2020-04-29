package leetcode

/*
  （这是一个 交互式问题 ）
  给你一个 山脉数组 mountainArr，请你返回能够使得 mountainArr.get(index) 等于 target 最小 的下标 index 值。
  如果不存在这样的下标 index，就请返回 -1。
  何为山脉数组？如果数组 A 是一个山脉数组的话，那它满足如下条件：
  首先，A.length >= 3
  其次，在 0 < i < A.length - 1 条件下，存在 i 使得：
    A[0] < A[1] < ... A[i-1] < A[i]
    A[i] > A[i+1] > ... > A[A.length - 1]
  你将 不能直接访问该山脉数组，必须通过 MountainArray 接口来获取数据：
    MountainArray.get(k) - 会返回数组中索引为k 的元素（下标从 0 开始）
    MountainArray.length() - 会返回该数组的长度
  注意：
  对 MountainArray.get 发起超过 100 次调用的提交将被视为错误答案。此外，任何试图规避判题系统的解决方案都将会导致比赛资格被取消。
  为了帮助大家更好地理解交互式问题，我们准备了一个样例 “答案”：https://leetcode-cn.com/playground/RKhe3ave，请注意这 不是一个正确答案。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-in-mountain-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

// MountainArray ...
type MountainArray struct {
	nums []int
}

func newMountainArray(capability int) *MountainArray {
	return &MountainArray{
		nums: make([]int, 0, capability),
	}
}

func (m *MountainArray) put(num int) { m.nums = append(m.nums, num) }

func (m *MountainArray) get(index int) int { return m.nums[index] }
func (m *MountainArray) length() int       { return len(m.nums) }

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1
	var top int
	for {
		m := l + (r-l)>>1
		mLeftVal := mountainArr.get(m - 1)
		mVal := mountainArr.get(m)
		mRightVal := mountainArr.get(m + 1)
		if mLeftVal < mVal && mVal > mRightVal {
			top = m
			break
		} else if mLeftVal > mVal {
			r = m
		} else if mVal < mRightVal {
			l = m
		}
	}
	binarySearchA := func(l, r int) int {
		for l <= r {
			m := l + (r-l)>>1
			mVal := mountainArr.get(m)
			if mVal < target {
				l = m + 1
			} else if mVal > target {
				r = m - 1
			} else {
				return m
			}
		}
		return -1
	}
	binarySearchB := func(l, r int) int {
		for l <= r {
			m := l + (r-l)>>1
			mVal := mountainArr.get(m)
			if mVal > target {
				l = m + 1
			} else if mVal < target {
				r = m - 1
			} else {
				return m
			}
		}
		return -1
	}
	lVal := binarySearchA(0, top)
	rVal := binarySearchB(top+1, mountainArr.length()-1)
	if lVal != -1 {
		return lVal
	}
	return rVal
}
