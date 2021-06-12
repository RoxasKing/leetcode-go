package main

/*
  Given an array of integers cost and an integer target. Return the maximum integer you can paint under the following rules:

    1. The cost of painting a digit (i+1) is given by cost[i] (0 indexed).
    2. The total cost used must be equal to target.
    3. Integer does not have digits 0.

  Since the answer may be too large, return it as string.

  If there is no way to paint any integer given the condition, return "0".

  Example 1:
    Input: cost = [4,3,2,5,6,7,2,5,5], target = 9
    Output: "7772"
    Explanation:  The cost to paint the digit '7' is 2, and the digit '2' is 3. Then cost("7772") = 2*3+ 3*1 = 9. You could also paint "977", but "7772" is the largest number.
        Digit    cost
          1  ->   4
          2  ->   3
          3  ->   2
          4  ->   5
          5  ->   6
          6  ->   7
          7  ->   2
          8  ->   5
          9  ->   5

  Example 2:
    Input: cost = [7,6,5,5,5,6,8,7,8], target = 12
    Output: "85"
    Explanation: The cost to paint the digit '8' is 7, and the digit '5' is 5. Then cost("85") = 7 + 5 = 12.

  Example 3:
    Input: cost = [2,4,6,2,4,6,4,4,4], target = 5
    Output: "0"
    Explanation: It's not possible to paint any integer with total cost equal to target.

  Example 4:
    Input: cost = [6,10,15,40,40,40,40,40,40], target = 47
    Output: "32211"

  Constraints:
    1. cost.length == 9
    2. 1 <= cost[i] <= 5000
    3. 1 <= target <= 5000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/form-largest-integer-with-digits-that-add-up-to-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
// Knapsack Problem

func largestNumber(cost []int, target int) string {
	f := make([]int, target+1)
	for _, c := range cost {
		for i := c; i <= target; i++ {
			if i-c > 0 && f[i-c] == 0 {
				continue
			}
			f[i] = Max(f[i], f[i-c]+1)
		}
	}

	if f[target] == 0 {
		return "0"
	}

	out := make([]byte, 0, f[target])
	for i, j := 8, target; i >= 0; i-- {
		for c := cost[i]; (j-c == 0 || j-c > 0 && f[j-c] > 0) && f[j] == f[j-c]+1; j -= c {
			out = append(out, byte(i)+'1')
		}
	}
	return string(out)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
