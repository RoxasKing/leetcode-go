package main

/*
  There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

  You are giving candies to these children subjected to the following requirements:
    1. Each child must have at least one candy.
    2. Children with a higher rating get more candies than their neighbors.
  Return the minimum number of candies you need to have to distribute the candies to the children.

  Example 1:
    Input: ratings = [1,0,2]
    Output: 5
    Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

  Example 2:
    Input: ratings = [1,2,2]
    Output: 4
    Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
      The third child gets 1 candy because it satisfies the above two conditions.

  Constraints:
    1. n == ratings.length
    2. 1 <= n <= 2 * 10^4
    3. 0 <= ratings[i] <= 2 * 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/candy
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm
func candy(ratings []int) int {
	n := len(ratings)
	dp := make([]int, n)
	q := []int{}

	for i, rt := range ratings {
		if i-1 >= 0 && ratings[i-1] < rt {
			continue
		}
		if i+1 < n && ratings[i+1] < rt {
			continue
		}
		dp[i] = 1
		q = append(q, i)
	}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if i-1 >= 0 && ratings[i-1] > ratings[i] && dp[i-1] < dp[i]+1 {
			dp[i-1] = dp[i] + 1
			q = append(q, i-1)
		}
		if i+1 < n && ratings[i+1] > ratings[i] && dp[i+1] < dp[i]+1 {
			dp[i+1] = dp[i] + 1
			q = append(q, i+1)
		}
	}

	out := 0
	for _, num := range dp {
		out += num
	}
	return out
}
