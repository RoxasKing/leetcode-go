package main

// Difficulty:
// Medium

// Tags:
// Hash
// Math

func originalDigits(s string) string {
	freq := [128]int{}
	for i := range s {
		freq[s[i]]++
	}
	cnt := [10]int{}
	cnt[0] = freq['z']
	cnt[2] = freq['w']
	cnt[4] = freq['u']
	cnt[6] = freq['x']
	cnt[8] = freq['g']
	cnt[3] = freq['h'] - cnt[8]
	cnt[5] = freq['f'] - cnt[4]
	cnt[7] = freq['s'] - cnt[6]
	cnt[1] = freq['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = freq['i'] - cnt[5] - cnt[6] - cnt[8]
	out := make([]byte, 0, cnt[0]+cnt[1]+cnt[2]+cnt[3]+cnt[4]+cnt[5]+cnt[6]+cnt[7]+cnt[8]+cnt[9])
	for i := 0; i < 10; i++ {
		for ; cnt[i] > 0; cnt[i]-- {
			out = append(out, byte(i)+'0')
		}
	}
	return string(out)
}
