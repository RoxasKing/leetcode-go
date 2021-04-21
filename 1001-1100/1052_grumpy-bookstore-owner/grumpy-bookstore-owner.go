package main

/*
  Today, the bookstore owner has a store open for customers.length minutes.  Every minute, some number of customers (customers[i]) enter the store, and all those customers leave after the end of that minute.

  On some minutes, the bookstore owner is grumpy.  If the bookstore owner is grumpy on the i-th minute, grumpy[i] = 1, otherwise grumpy[i] = 0.  When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise they are satisfied.

  The bookstore owner knows a secret technique to keep themselves not grumpy for X minutes straight, but can only use it once.

  Return the maximum number of customers that can be satisfied throughout the day.

  Example 1:
    Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
    Output: 16
    Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes.
    The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.

  Note:
    1. 1 <= X <= customers.length == grumpy.length <= 20000
    2. 0 <= customers[i] <= 1000
    3. 0 <= grumpy[i] <= 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/grumpy-bookstore-owner
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func maxSatisfied(customers []int, grumpy []int, X int) int {
	out, max := 0, 0
	for i, cur := 0, 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			out += customers[i]
		} else {
			cur += customers[i]
		}

		if i > X-1 && grumpy[i-X] == 1 {
			cur -= customers[i-X]
		}

		if i >= X-1 && cur > max {
			max = cur
		}
	}
	out += max
	return out
}
