package sort

func RadixSort(array []int) {
	max := array[0]
	for i := range array {
		if array[i] > max {
			max = array[i]
		}
	}
	var maxLen int
	for i := max; i != 0; i /= 10 {
		maxLen++
	}
	mod := 10
	for i := 0; i < maxLen; i, mod = i+1, mod*10 {
		dp := make([][]int, mod)
		for j := 0; j < len(array); j++ {
			bucket := array[j] % mod
			dp[bucket] = append(dp[bucket], array[j])
		}
		var index int
		for j := range dp {
			for k := range dp[j] {
				array[index] = dp[j][k]
				index++
			}
		}
	}
}

// Considering negative cases
func RadixSort2(array []int) {
	max := array[0]
	for i := range array {
		if array[i] > max {
			max = array[i]
		}
	}
	var maxLen int
	for i := max; i != 0; i /= 10 {
		maxLen++
	}
	mod, dev := 10, 1
	for i := 0; i < maxLen; i, dev, mod = i+1, dev*10, mod*10 {
		counter := make([][]int, mod*2)
		for j := 0; j < len(array); j++ {
			bucket := ((array[j] % mod) / dev) + mod
			counter[bucket] = append(counter[bucket], array[j])
		}
		var pos int
		for k := range counter {
			for l := range counter[k] {
				array[pos] = counter[k][l]
				pos++
			}
		}
	}
}
