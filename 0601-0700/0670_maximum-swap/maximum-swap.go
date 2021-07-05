package main

// Tags:
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
