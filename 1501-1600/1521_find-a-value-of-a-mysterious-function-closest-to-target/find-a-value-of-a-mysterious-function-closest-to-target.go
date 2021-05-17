package main

/*
  Winston was given the above mysterious function func. He has an integer array arr and an integer target and he wants to find the values l and r that make the value |func(arr, l, r) - target| minimum possible.

  Return the minimum possible value of |func(arr, l, r) - target|.

  Notice that func should be called with the values l and r where 0 <= l, r < arr.length.

  Example 1:
    Input: arr = [9,12,3,7,15], target = 5
    Output: 2
    Explanation: Calling func with all the pairs of [l,r] = [[0,0],[1,1],[2,2],[3,3],[4,4],[0,1],[1,2],[2,3],[3,4],[0,2],[1,3],[2,4],[0,3],[1,4],[0,4]], Winston got the following results [9,12,3,7,15,8,0,3,7,0,0,3,0,0,0]. The value closest to 5 is 7 and 3, thus the minimum difference is 2.

  Example 2:
    Input: arr = [1000000,1000000,1000000], target = 1
    Output: 999999
    Explanation: Winston called the func with all possible values of [l,r] and he always got 1000000, thus the min difference is 999999.

  Example 3:
    Input: arr = [1,2,4,8,16], target = 0
    Output: 0

  Constraints:
    1. 1 <= arr.length <= 10^5
    2. 1 <= arr[i] <= 10^6
    3. 0 <= target <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-a-value-of-a-mysterious-function-closest-to-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Segment Tree
// Sliding Window

func closestToTarget(arr []int, target int) int {
	n := len(arr)
	f := make([]int, 4*n)
	build(f, 1, 0, n-1, arr)

	out := 1<<31 - 1
	for l, r := 0, 0; r < n; {
		res := query(f, 1, 0, n-1, l, r)
		if res < target {
			l++
			r = Max(l, r)
		} else {
			r++
		}
		out = Min(out, Abs(res-target))
	}
	return out
}

func build(f []int, i, s, t int, arr []int) {
	if s == t {
		f[i] = arr[s]
		return
	}
	m := (s + t) >> 1
	build(f, i<<1, s, m, arr)
	build(f, i<<1+1, m+1, t, arr)
	f[i] = f[i<<1] & f[i<<1+1]
}

func query(f []int, i, s, t, l, r int) int {
	if t < l || s > r {
		return 1<<63 - 1
	}
	if l <= s && t <= r {
		return f[i]
	}
	m := (s + t) >> 1
	return query(f, i<<1, s, m, l, r) & query(f, i<<1+1, m+1, t, l, r)
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
