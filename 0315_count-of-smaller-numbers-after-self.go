package leetcode

/*
  给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	counter := make([]int, len(nums))
	tmpIdxs := make([]int, len(nums))
	indexes := make([]int, len(nums))
	for i := range indexes {
		indexes[i] = i
	}
	var merge func(int, int)
	merge = func(l, r int) {
		if l == r {
			return
		}
		m := l + (r-l)>>1
		merge(l, m)
		merge(m+1, r)
		if nums[indexes[m]] <= nums[indexes[m+1]] {
			return
		}
		copy(tmpIdxs[l:r+1], indexes[l:r+1])
		i, j := l, m+1
		for k := l; k <= r; k++ {
			switch {
			case i > m:
				indexes[k] = tmpIdxs[j]
				j++
			case j > r:
				indexes[k] = tmpIdxs[i]
				i++
				counter[indexes[k]] += r - m
			case nums[tmpIdxs[i]] <= nums[tmpIdxs[j]]:
				indexes[k] = tmpIdxs[i]
				i++
				counter[indexes[k]] += j - m - 1
			default:
				indexes[k] = tmpIdxs[j]
				j++
			}
		}
	}
	merge(0, len(nums)-1)
	return counter
}
