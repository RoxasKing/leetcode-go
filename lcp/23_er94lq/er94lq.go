package main

// Tags:
// Simulation

func isMagic(target []int) bool {
	n := len(target)
	k := 0
	for k+1 < n && target[k]+2 == target[k+1] {
		k++
	}
	if target[k] == n/2*2 && k+1 < n && target[k+1] == 1 {
		k++
		for k+1 < n && target[k]+2 == target[k+1] {
			k++
		}
	}
	k++

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}

	for {
		tmp := make([]int, 0, len(arr))
		for i := 1; i < len(arr); i += 2 {
			tmp = append(tmp, arr[i])
		}
		for i := 0; i < len(arr); i += 2 {
			tmp = append(tmp, arr[i])
		}
		for i := 0; i < k && i < len(arr); i++ {
			if tmp[i] != target[i] {
				return false
			}
		}

		if k >= len(arr) {
			break
		}

		arr = tmp[k:]
		target = target[k:]
	}

	return true
}
