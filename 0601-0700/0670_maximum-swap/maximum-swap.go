package main

/*
  Given a non-negative integer, you could swap two digits at most once to get the maximum valued number. Return the maximum valued number you could get.

  Example 1:
    Input: 2736
    Output: 7236
    Explanation: Swap the number 2 and the number 7.

  Example 2:
    Input: 9973
    Output: 9973
    Explanation: No swap.

  Note:
    The given number is in the range [0, 108]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-swap
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math + Greedy Algorithm
func maximumSwap(num int) int {
	arr := []int{}
	for ; num > 0; num /= 10 {
		arr = append(arr, num%10)
	}

	n := len(arr)
	for i := 0; i < n>>1; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	for i := 0; i < n-1; i++ {
		max, j := arr[i], i
		for k := n - 1; k > i; k-- {
			if arr[k] > max {
				max, j = arr[k], k
			}
		}
		if max > arr[i] {
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}

	out := 0
	for _, num := range arr {
		out = out*10 + num
	}
	return out
}
