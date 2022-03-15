package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation

func validUtf8(data []int) bool {
	n := len(data)
	for i := 0; i < n; i++ {
		head := data[i]
		cnt := 0
		for k := 7; k >= 0 && (head>>k)&1 == 1; k-- {
			cnt++
		}
		if cnt == 1 || cnt > 4 || cnt > n-i {
			return false
		}
		for ; cnt > 1; cnt-- {
			if i++; data[i]&0b11000000 == 0b10000000 {
				continue
			}
			return false
		}
	}
	return true
}
