package main

// Difficulty:
// Medium

// Tags:
// Boyer–Moore Majority Vote Algorithm

func majorityElement(nums []int) []int {
	n1, n2 := 0, 0
	v1, v2 := 0, 0
	for _, num := range nums {
		if v1 > 0 && num == n1 {
			v1++
		} else if v2 > 0 && num == n2 {
			v2++
		} else if v1 == 0 {
			n1 = num
			v1++
		} else if v2 == 0 {
			n2 = num
			v2++
		} else {
			v1--
			v2--
		}
	}

	c1, c2 := 0, 0
	for _, num := range nums {
		if v1 > 0 && num == n1 {
			c1++
		}
		if v2 > 0 && num == n2 {
			c2++
		}
	}

	out := []int{}
	k := len(nums) / 3
	if v1 > 0 && c1 > k {
		out = append(out, n1)
	}
	if v2 > 0 && c2 > k {
		out = append(out, n2)
	}
	return out
}
