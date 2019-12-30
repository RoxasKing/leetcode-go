package sort

func CountingSort(array []int) {
	max := array[0]
	for i := range array {
		if array[i] > max {
			max = array[i]
		}
	}
	dp := make([]int, max+1)
	for i := range array {
		dp[array[i]]++
	}
	var index int
	for i := range dp {
		for dp[i] > 0 {
			array[index] = i
			index++
			dp[i]--
		}
	}
}
