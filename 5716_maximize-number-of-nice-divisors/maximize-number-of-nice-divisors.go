package main

/*
  You are given a positive integer primeFactors. You are asked to construct a positive integer n that satisfies the following conditions:

  The number of prime factors of n (not necessarily distinct) is at most primeFactors.
  The number of nice divisors of n is maximized. Note that a divisor of n is nice if it is divisible by every prime factor of n. For example, if n = 12, then its prime factors are [2,2,3], then 6 and 12 are nice divisors, while 3 and 4 are not.
  Return the number of nice divisors of n. Since that number can be too large, return it modulo 109 + 7.

  Note that a prime number is a natural number greater than 1 that is not a product of two smaller natural numbers. The prime factors of a number n is a list of prime numbers such that their product equals n.

  Example 1:
    Input: primeFactors = 5
    Output: 6
    Explanation: 200 is a valid value of n.
    It has 5 prime factors: [2,2,2,5,5], and it has 6 nice divisors: [10,20,40,50,100,200].
    There is not other value of n that has at most 5 prime factors and more nice divisors.

  Example 2:
    Input: primeFactors = 8
    Output: 18

  Constraints:
    1 <= primeFactors <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximize-number-of-nice-divisors
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm
func maxNiceDivisors(primeFactors int) int {
	if primeFactors == 1 {
		return 1
	}
	timesOf3 := primeFactors / 3
	if primeFactors-timesOf3*3 == 1 {
		timesOf3--
	}
	timesOf2 := (primeFactors - timesOf3*3) / 2
	return QuickMul(3, timesOf3) * QuickMul(2, timesOf2) % (1e9 + 7)
}

func QuickMul(x, n int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out = (out * x) % (1e9 + 7)
		}
		x = (x * x) % (1e9 + 7)
		n >>= 1
	}
	return out
}
