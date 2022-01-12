package main

// Difficulty:
// Easy

// Tags:
// Hash

func slowestKey(releaseTimes []int, keysPressed string) byte {
	cnt := [26]int{}
	cnt[keysPressed[0]-'a'] += releaseTimes[0]
	out := keysPressed[0]
	max := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		x := keysPressed[i] - 'a'
		if t := releaseTimes[i] - releaseTimes[i-1]; cnt[x] < t {
			cnt[x] = t
		}
		if max < cnt[x] {
			max = cnt[x]
			out = keysPressed[i]
		} else if max == cnt[x] && out < keysPressed[i] {
			out = keysPressed[i]
		}
	}
	return out
}
