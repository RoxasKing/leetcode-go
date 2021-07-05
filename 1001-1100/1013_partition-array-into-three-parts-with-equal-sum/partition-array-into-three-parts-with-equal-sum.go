package main

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	avg := sum / 3
	cnt := 0
	sum = 0
	for _, num := range arr {
		sum += num
		if sum == avg {
			cnt++
			sum = 0
		}
	}
	return cnt >= 3
}
