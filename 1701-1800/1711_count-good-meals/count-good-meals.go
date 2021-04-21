package main

import "sort"

/*
  A good meal is a meal that contains exactly two different food items with a sum of deliciousness equal to a power of two.

  You can pick any two different foods to make a good meal.

  Given an array of integers deliciousness where deliciousness[i] is the deliciousness of the i​​​​​​th​​​​​​​​ item of food, return the number of different good meals you can make from this list modulo 109 + 7.

  Note that items with different indices are considered different even if they have the same deliciousness value.

  Example 1:
    Input: deliciousness = [1,3,5,7,9]
    Output: 4
    Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
    Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.

  Example 2:
    Input: deliciousness = [1,1,1,3,3,3,7]
    Output: 15
    Explanation: The good meals are (1,1) with 3 ways, (1,3) with 9 ways, and (1,7) with 3 ways.

  Constraints:
    1. 1 <= deliciousness.length <= 10^5
    2. 0 <= deliciousness[i] <= 2^20

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-good-meals
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func countPairs(deliciousness []int) int {
	arr := []int{}
	cnt := make(map[int]int)
	for _, d := range deliciousness {
		if cnt[d] == 0 {
			arr = append(arr, d)
		}
		cnt[d]++
	}
	sort.Ints(arr)

	out := 0
	for good := 1; good <= 1<<21; good <<= 1 {
		for l, r := 0, len(arr)-1; l <= r; {
			sum := arr[l] + arr[r]
			if sum < good {
				l++
			} else if sum > good {
				r--
			} else {
				if l == r {
					out = (out + cnt[arr[l]]*(cnt[arr[l]]-1)/2) % (1e9 + 7)
				} else {
					out = (out + cnt[arr[l]]*cnt[arr[r]]) % (1e9 + 7)
				}
				l++
				r--
			}
		}
	}
	return out
}
