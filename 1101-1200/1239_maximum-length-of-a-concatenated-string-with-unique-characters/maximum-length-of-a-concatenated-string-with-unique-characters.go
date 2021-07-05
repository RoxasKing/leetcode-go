package main

// Tags:
// Backtracking

func maxLength(arr []string) int {
	n := len(arr)
	nums := make([]int, n)
	for i := range arr {
		cnt := [26]int{}
		for j := range arr[i] {
			cnt[arr[i][j]-'a']++
			if cnt[arr[i][j]-'a'] > 1 {
				nums[i] = -1
				break
			}
			nums[i] |= 1 << int(arr[i][j]-'a')
		}
	}
	out := 0
	backtrack(arr, nums, n, 0, 0, 0, &out)
	return out
}

func backtrack(arr []string, nums []int, n, i, curVal, curLen int, out *int) {
	if i == n {
		return
	}

	backtrack(arr, nums, n, i+1, curVal, curLen, out)

	if nums[i] != -1 && curVal^nums[i] == curVal+nums[i] {
		curVal += nums[i]
		curLen += len(arr[i])
		*out = Max(*out, curLen)
		backtrack(arr, nums, n, i+1, curVal, curLen, out)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
