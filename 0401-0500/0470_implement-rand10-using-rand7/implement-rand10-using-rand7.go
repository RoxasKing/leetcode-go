package main

import "math/rand"

/*
  Given the API rand7() that generates a uniform random integer in the range [1, 7], write a function rand10() that generates a uniform random integer in the range [1, 10]. You can only call the API rand7(), and you shouldn't call any other API. Please do not use a language's built-in random API.

  Each test case will have one internal argument n, the number of times that your implemented function rand10() will be called while testing. Note that this is not an argument passed to rand10().

  Follow up:
    1. What is the expected value for the number of calls to rand7() function?
    2. Could you minimize the number of calls to rand7()?

  Example 1:
    Input: n = 1
    Output: [2]

  Example 2:
    Input: n = 2
    Output: [2,8]

  Example 3:
    Input: n = 3
    Output: [3,8,10]

  Constraints:
    1 <= n <= 105

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-rand10-using-rand7
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math + Bit Manipulation
func rand10() int {
	a, b := 7, 7

	for a == 7 { // [1, 6]
		a = rand7()
	}
	for b > 5 { // [1, 5]
		b = rand7()
	}

	if a&1 == 1 { // 50% [0, 1] * 20% [1, 5]
		return b
	}
	return b + 5
}

func rand7() int {
	return rand.Intn(7) + 1
}
