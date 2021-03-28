package main

/*
  Write an algorithm to determine if a number n is happy.

  A happy number is a number defined by the following process:

  Starting with any positive integer, replace the number by the sum of the squares of its digits.
  Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
  Those numbers for which this process ends in 1 are happy.
  Return true if n is a happy number, and false if not.

  Example 1:
    Input: n = 19
    Output: true
    Explanation:
    12 + 92 = 82
    82 + 22 = 68
    62 + 82 = 100
    12 + 02 + 02 = 1

  Example 2:
    Input: n = 2
    Output: false

  Constraints:
    1 <= n <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/happy-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func isHappy(n int) bool {
	slow := getNext(n)
	fast := getNext(slow)
	for fast != 1 {
		if slow == fast {
			return false
		}
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return true
}

func getNext(n int) int {
	var res int
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

func isHappy2(n int) bool {
	dict := map[int]struct{}{
		4:   {},
		16:  {},
		37:  {},
		58:  {},
		89:  {},
		145: {},
		42:  {},
		20:  {},
	}
	for {
		if n == 1 {
			return true
		}
		if _, ok := dict[n]; ok {
			return false
		}
		var newN int
		for n != 0 {
			newN += (n % 10) * (n % 10)
			n /= 10
		}
		n = newN
	}
}
