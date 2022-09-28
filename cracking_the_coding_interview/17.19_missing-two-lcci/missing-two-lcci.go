package main

// Difficulty:
// Hard

// Tags:
// Hash

func missingTwo(nums []int) []int {
	f := [3e4 + 3]bool{}
	for _, x := range nums {
		f[x] = true
	}
	o := make([]int, 0, 2)
	for i := 1; i <= 3e4+2; i++ {
		if ok := f[i]; !ok {
			if o = append(o, i); len(o) == 2 {
				break
			}
		}
	}
	return o
}
