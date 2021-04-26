package main

import "sort"

/*
  Winter is coming! During the contest, your first job is to design a standard heater with a fixed warm radius to warm all the houses.

  Every house can be warmed, as long as the house is within the heater's warm radius range.

  Given the positions of houses and heaters on a horizontal line, return the minimum radius standard of heaters so that those heaters could cover all houses.

  Notice that all the heaters follow your radius standard, and the warm radius will the same.

  Example 1:
    Input: houses = [1,2,3], heaters = [2]
    Output: 1
    Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.

  Example 2:
    Input: houses = [1,2,3,4], heaters = [1,4]
    Output: 1
    Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.

  Example 3:
    Input: houses = [1,5], heaters = [2]
    Output: 3

  Constraints:
    1. 1 <= houses.length, heaters.length <= 3 * 10^4
    2. 1 <= houses[i], heaters[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/heaters
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findRadius(houses []int, heaters []int) int {
	m, n := len(houses), len(heaters)
	sort.Ints(houses)
	sort.Ints(heaters)
	l, r := 0, int(1e9)
	for l < r {
		radius := (l + r) >> 1
		if !isValid(houses, heaters, m, n, radius) {
			l = radius + 1
		} else {
			r = radius
		}
	}
	return l
}

func isValid(houses, heaters []int, m, n, radius int) bool {
	i, j := 0, 0
	for i < m && j < n {
		l, r := heaters[j]-radius, heaters[j]+radius
		if l <= houses[i] && houses[i] <= r {
			i++
		} else {
			j++
		}
	}
	return i == m
}
