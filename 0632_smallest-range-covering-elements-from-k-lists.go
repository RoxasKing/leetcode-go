package main

/*
  你有 k 个升序排列的整数数组。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
  我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。

  注意:
    给定的列表可能包含重复元素，所以在这里升序表示 >= 。
    1 <= k <= 3500
    -105 <= 元素的值 <= 105
    对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func smallestRange(nums [][]int) []int {
	indices := make(map[int][]int)
	min, max := 1<<31-1, -1<<31
	for i := range nums {
		for _, num := range nums[i] {
			indices[num] = append(indices[num], i)
			min = Min(min, num)
			max = Max(max, num)
		}
	}
	var inside int
	count := make([]int, len(nums))
	borderL, borderR := min, max
	for l, r := min, min; r <= max; r++ {
		if len(indices[r]) == 0 {
			continue
		}
		for _, x := range indices[r] {
			count[x]++
			if count[x] == 1 {
				inside++
			}
		}
		for inside == len(nums) {
			if r-l < borderR-borderL {
				borderL, borderR = l, r
			}
			for _, x := range indices[l] {
				count[x]--
				if count[x] == 0 {
					inside--
				}
			}
			l++
		}
	}
	return []int{borderL, borderR}
}
