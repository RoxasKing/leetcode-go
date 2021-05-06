package main

/*
  There is a room with n bulbs labeled from 1 to n that all are turned on initially, and four buttons on the wall. Each of the four buttons has a different functionality where:

    1. Button 1: Flips the status of all the bulbs.
    2. Button 2: Flips the status of all the bulbs with even labels (i.e., 2, 4, ...).
    3. Button 3: Flips the status of all the bulbs with odd labels (i.e., 1, 3, ...).
    4. Button 4: Flips the status of all the bulbs with a label j = 3k + 1 where k = 0, 1, 2, ... (i.e., 1, 4, 7, 10, ...).

  You will press one of the four mentioned buttons exactly presses times.

  Given the two integers n and presses, return the number of different statuses after pressing the four buttons exactly presses times.

  Example 1:
    Input: n = 1, presses = 1
    Output: 2
    Explanation: Status can be: [on], [off].

  Example 2:
    Input: n = 2, presses = 1
    Output: 3
    Explanation: Status can be: [on, off], [off, on], [off, off].

  Example 3:
    Input: n = 3, presses = 1
    Output: 4
    Explanation: Status can be: [off, on, off], [on, off, on], [off, off, off], [off, on, on].

  Constraints:
    1. 1 <= n <= 1000
    2. 0 <= presses <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/bulb-switcher-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// ?

// Math
func flipLights(n int, presses int) int {
	if n > 3 {
		n = 3
	}
	if presses == 0 {
		return 1
	} else if presses == 1 {
		return []int{2, 3, 4}[n-1]
	} else if presses == 2 {
		return []int{2, 4, 7}[n-1]
	}
	return []int{2, 4, 8}[n-1]
}
