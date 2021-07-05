package main

// Tags:
// Bit Manipulation
// Prefix Sum

func wonderfulSubstrings(word string) int64 {
	cnt := map[int64]int64{0: 1}
	var out, sum int64
	for i := range word {
		sum ^= 1 << int(word[i]-'a')
		out += cnt[sum]
		for j := 0; j < 10; j++ {
			out += cnt[sum^(1<<j)]
		}
		cnt[sum]++
	}
	return out
}
