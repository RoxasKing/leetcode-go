package main

/*
  You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to find the length of the longest sequence of 1s you could create.

  Example 1:
    Input: num = 1775(110111011112)
    Output: 8

  Example 2:
    Input: num = 7(01112)
    Output: 4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-bits-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming

func reverseBits(num int) int {
	out := 0
	x := uint(num)
	f := [32][2]int{}
	f[0] = [2]int{int(x & 1), 1}
	for i := 1; i <= 31; i++ {
		if (x>>i)&1 == 1 {
			f[i] = [2]int{f[i-1][0] + 1, f[i-1][1] + 1}
		} else {
			f[i][1] = f[i-1][0] + 1
		}
		out = Max(out, Max(f[i][0], f[i][1]))
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
