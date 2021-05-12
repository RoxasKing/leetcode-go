package main

/*
  Let f(x) be the number of zeroes at the end of x!. (Recall that x! = 1 * 2 * 3 * ... * x, and by convention, 0! = 1.)

  For example, f(3) = 0 because 3! = 6 has no zeroes at the end, while f(11) = 2 because 11! = 39916800 has 2 zeroes at the end. Given k, find how many non-negative integers x have the property that f(x) = k.

  Example 1:
    Input: k = 0
    Output: 5
    Explanation: 0!, 1!, 2!, 3!, and 4! end with k = 0 zeroes.

  Example 2:
    Input: k = 5
    Output: 0
    Explanation: There is no x such that x! ends in k = 5 zeroes.

  Note:
    k will be an integer in the range [0, 10^9].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math + Binary Search
func preimageSizeFZF(k int) int {
	l, r := 0, 10*k+1
	for l < r {
		m := (l + r) >> 1
		res := countZero(m)
		if res < k {
			l = m + 1
		} else if res > k {
			r = m - 1
		} else {
			return 5
		}
	}
	return 0
}

func countZero(n int) int {
	out := 0
	for n >= 5 {
		n /= 5
		out += n
	}
	return out
}
