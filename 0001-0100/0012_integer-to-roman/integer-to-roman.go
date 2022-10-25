package main

// Difficulty:
// Medium

// Tags:
// Math

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	out, n := "", len(nums)
	for i := 0; num >= 0 && i < n; i++ {
		for nums[i] <= num {
			out += strs[i]
			num -= nums[i]
		}
	}
	return out
}
