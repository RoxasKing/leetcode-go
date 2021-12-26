package main

// Dfficulty:
// Medium

// Tags:
// Recursion

func decodeString(s string) string {
	arr := []byte{}
	n := len(s)
	freq := 0
	for i := 0; i < n; i++ {
		if '0' <= s[i] && s[i] <= '9' {
			freq *= 10
			freq += int(s[i] - '0')
		} else if s[i] == '[' {
			l, r := i, i+1
			cnt := 1
			for ; r < len(s) && cnt > 0; r++ {
				if s[r] == '[' {
					cnt++
				} else if s[r] == ']' {
					cnt--
				}
			}
			p := []byte(decodeString(s[l+1 : r-1]))
			for ; freq > 0; freq-- {
				arr = append(arr, p...)
			}
			i = r - 1
		} else {
			arr = append(arr, s[i])
		}
	}
	return string(arr)
}
