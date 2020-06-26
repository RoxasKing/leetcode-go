package leetcode

/*
  给定一个排序好的数组，两个整数 k 和 x，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。如果有两个数与 x 的差值一样，优先选择数值较小的那个数。

  说明:
    k 的值为正数，且总是小于给定排序数组的长度。
    数组不为空，且长度不超过 104
    数组里的每个元素与 x 的绝对值不超过 104

  更新(2017/9/19):
  这个参数 arr 已经被改变为一个整数数组（而不是整数列表）。 请重新加载代码定义以获取最新更改。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-k-closest-elements
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findClosestElements(arr []int, k int, x int) []int {
	var index int
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)>>1
		if Abs(arr[m]-x) < Abs(arr[index]-x) {
			index = m
		}
		if arr[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	l, r = index, index+1
	for r-l < k {
		if l == 0 {
			return arr[:k]
		} else if r == len(arr) {
			return arr[len(arr)-k:]
		} else {
			if Abs(arr[l-1]-x) > Abs(arr[r]-x) {
				r++
			} else {
				l--
			}
		}
	}
	return arr[l:r]
}
