package main

// Tags:
// Greedy

func maximumNumber(num string, change []int) string {
	nums := []byte(num)
	var center int
	for i := range num {
		idx := int(num[i] - '0')
		if idx < change[idx] {
			nums[i] = byte(change[idx]) + '0'
			center = i
			break
		}
	}

	for i := center - 1; i >= 0; i-- {
		idx := int(num[i] - '0')
		if idx > change[idx] {
			break
		}
		nums[i] = byte(change[idx]) + '0'
	}

	for i := center + 1; i < len(num); i++ {
		idx := int(num[i] - '0')
		if idx > change[idx] {
			break
		}
		nums[i] = byte(change[idx]) + '0'
	}

	return string(nums)
}
