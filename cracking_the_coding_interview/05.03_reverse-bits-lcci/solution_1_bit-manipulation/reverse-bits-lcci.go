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

// Bit Manipulation

func reverseBits(num int) int {
	arr := []int{}
	cnt := 0
	x := uint(num)
	for i := 0; i <= 31; i++ {
		if (x>>i)&1 == 1 {
			cnt++
		} else {
			arr = append(arr, cnt, 0)
			cnt = 0
		}
	}
	arr = append(arr, cnt)
	out := 0
	for i := range arr {
		if arr[i] != 0 {
			out = Max(out, arr[i])
			continue
		}
		cnt := 1
		if i > 0 {
			cnt += arr[i-1]
		}
		if i < len(arr)-1 {
			cnt += arr[i+1]
		}
		out = Max(out, cnt)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
