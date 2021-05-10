package main

/*
  You are given a 2D integer array logs where each logs[i] = [birthi, deathi] indicates the birth and death years of the ith person.

  The population of some year x is the number of people alive during that year. The ith person is counted in year x's population if x is in the inclusive range [birthi, deathi - 1]. Note that the person is not counted in the year that they die.

  Return the earliest year with the maximum population.

  Example 1:
    Input: logs = [[1993,1999],[2000,2010]]
    Output: 1993
    Explanation: The maximum population is 1, and 1993 is the earliest year with this population.

  Example 2:
    Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
    Output: 1960
    Explanation:
      The maximum population is 2, and it had happened in years 1960 and 1970.
      The earlier year between them is 1960.

  Constraints:
    1 <= logs.length <= 100
    1950 <= birthi < deathi <= 2050

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-population-year
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximumPopulation(logs [][]int) int {
	min, max := 1<<31-1, 0
	for _, log := range logs {
		min = Min(min, log[0])
		max = Max(max, log[1])
	}

	out, cnt := 0, 0
	for i := min; i <= max; i++ {
		tmp := 0
		for _, log := range logs {
			if log[0] <= i && i <= log[1]-1 {
				tmp++
			}
		}
		if tmp > cnt {
			out, cnt = i, tmp
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
