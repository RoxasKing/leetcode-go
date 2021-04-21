package main

/*
  You are given two jugs with capacities jug1Capacity and jug2Capacity liters. There is an infinite amount of water supply available. Determine whether it is possible to measure exactly targetCapacity liters using these two jugs.

  If targetCapacity liters of water are measurable, you must have targetCapacity liters of water contained within one or both buckets by the end.

  Operations allowed:
    1. Fill any of the jugs with water.
    2. Empty any of the jugs.
    3. Pour water from one jug into another till the other jug is completely full, or the first jug itself is empty.

  Example 1:
    Input: jug1Capacity = 3, jug2Capacity = 5, targetCapacity = 4
    Output: true
    Explanation: The famous Die Hard example

  Example 2:
    Input: jug1Capacity = 2, jug2Capacity = 6, targetCapacity = 5
    Output: false

  Example 3:
    Input: jug1Capacity = 1, jug2Capacity = 2, targetCapacity = 3
    Output: true

  Constraints:
    1 <= jug1Capacity, jug2Capacity, targetCapacity <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/water-and-jug-problem
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}
	if jug1Capacity == 0 || jug2Capacity == 0 {
		return targetCapacity == 0 || jug1Capacity+jug2Capacity == targetCapacity
	}
	return targetCapacity%Gcd(jug1Capacity, jug2Capacity) == 0
}

func Gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
