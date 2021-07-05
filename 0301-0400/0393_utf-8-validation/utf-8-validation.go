package main

// Tags:
// Bit Manipulation

func validUtf8(data []int) bool {
	for i := 0; i < len(data); i++ {
		t := 0
		if (data[i]>>3)^30 == 0 {
			t = 3
		} else if (data[i]>>4)^14 == 0 {
			t = 2
		} else if (data[i]>>5)^6 == 0 {
			t = 1
		}

		if t == 0 && (data[i]>>7) != 0 {
			return false
		}

		for ; t > 0; t-- {
			if i++; i == len(data) || (data[i]>>6)^2 != 0 {
				return false
			}
		}
	}
	return true
}
