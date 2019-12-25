package My_LeetCode_In_Go

func maxProfit(prices []int) int {
	var out int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			out += prices[i] - prices[i-1]
		}
	}
	return out
}
