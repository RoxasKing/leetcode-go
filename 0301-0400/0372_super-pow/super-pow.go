package main

/*
  Your task is to calculate ab mod 1337 where a is a positive integer and b is an extremely large positive integer given in the form of an array.

  Example 1:
    Input: a = 2, b = [3]
    Output: 8

  Example 2:
    Input: a = 2, b = [1,0]
    Output: 1024

  Example 3:
    Input: a = 1, b = [4,3,3,8,5,2]
    Output: 1

  Example 4:
    Input: a = 2147483647, b = [2,0,0]
    Output: 1198

  Constraints:
    1. 1 <= a <= 2^31 - 1
    2. 1 <= b.length <= 2000
    3. 0 <= b[i] <= 9
    4. b doesn't contain leading zeros.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/super-pow
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func superPow(a int, b []int) int {
	a %= mod
	out := quickMul(a, b[0])
	for _, n := range b[1:] {
		out = quickMul(out, 10) * quickMul(a, n)
		out %= mod
	}
	return out
}

func quickMul(x, n int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out *= x
			out %= mod
		}
		x *= x
		x %= mod
		n >>= 1
	}
	return out
}

var mod = 1337
