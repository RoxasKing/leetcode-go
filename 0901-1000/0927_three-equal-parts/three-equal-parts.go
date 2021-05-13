package main

/*
  You are given an array arr which consists of only zeros and ones, divide the array into three non-empty parts such that all of these parts represent the same binary value.

  If it is possible, return any [i, j] with i + 1 < j, such that:

    1. arr[0], arr[1], ..., arr[i] is the first part,
    2. arr[i + 1], arr[i + 2], ..., arr[j - 1] is the second part, and
    3. arr[j], arr[j + 1], ..., arr[arr.length - 1] is the third part.
    4. All three parts have equal binary values.

  If it is not possible, return [-1, -1].

  Note that the entire part is used when considering what binary value it represents. For example, [1,1,0] represents 6 in decimal, not 3. Also, leading zeros are allowed, so [0,1,1] and [1,1] represent the same value.

  Example 1:
    Input: arr = [1,0,1,0,1]
    Output: [0,3]

  Example 2:
    Input: arr = [1,1,0,1,1]
    Output: [-1,-1]

  Example 3:
    Input: arr = [1,1,0,0,1]
    Output: [0,2]

  Constraints:
    1. 3 <= arr.length <= 3 * 10^4
    2. arr[i] is 0 or 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/three-equal-parts
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math

func threeEqualParts(arr []int) []int {
	cnt := 0
	for _, num := range arr {
		if num == 1 {
			cnt++
		}
	}
	if cnt%3 != 0 {
		return []int{-1, -1}
	}

	target := cnt / 3

	cnt, a, b := 0, -1, -1
	for i, num := range arr {
		if num == 1 {
			cnt++
		}

		if cnt < target {
			continue
		}
		cnt = 0

		if a == -1 {
			a = i
			continue
		}

		b = i

		num3 := 0
		for j := b + 1; j < len(arr); j++ {
			num3 = num3<<1 + arr[j]
			num3 %= 1e9 + 7
		}

		num1 := 0
		for j := 0; j <= a; j++ {
			num1 = num1<<1 + arr[j]
			num1 %= 1e9 + 7
		}
		for arr[a+1] == 0 && num1 != num3 {
			a++
			num1 <<= 1
			num1 %= 1e9 + 7
		}

		if num1 != num3 {
			return []int{-1, -1}
		}

		num2 := 0
		for j := a + 1; j <= b; j++ {
			num2 = num2<<1 + arr[j]
			num2 %= 1e9 + 7
		}
		for arr[b+1] == 0 && num2 != num3 {
			b++
			num2 <<= 1
			num2 %= 1e9 + 7
		}

		if num2 != num3 {
			return []int{-1, -1}
		}

		b++
		break
	}

	return []int{a, b}
}
