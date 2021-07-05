package main

// Tags:
// Dynamic Programming
func maxHappyGroups(batchSize int, groups []int) int {
	freq0 := make([]int, 9)
	freq := make([]int, 9)
	w := make([]int, 9)

	for _, g := range groups {
		freq0[g%batchSize]++
	}

	msum := 1
	for i := 1; i < batchSize; i++ {
		msum *= (freq0[i] + 1)
	}

	w[1] = 1
	for i := 2; i < batchSize; i++ {
		w[i] = w[i-1] * (freq0[i-1] + 1)
	}

	f := make([]int, msum)
	for fmask := 1; fmask < msum; fmask++ {
		last := 0
		for i := 1; i < batchSize; i++ {
			freq[i] = (fmask / w[i]) % (freq0[i] + 1)
			last = (last + (freq0[i]-freq[i])*i) % batchSize
		}
		for c := 1; c < batchSize; c++ {
			if freq[c] == 0 {
				continue
			}
			cnt := f[fmask-w[c]]
			if last == 0 {
				cnt++
			}
			f[fmask] = Max(f[fmask], cnt)
		}
	}
	return f[msum-1] + freq0[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
