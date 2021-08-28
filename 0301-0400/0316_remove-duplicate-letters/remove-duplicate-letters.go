package main

// Tags:
// Greedy
// Stack

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	mark := [26]bool{}
	stack := []byte{}
	for i := range s {
		if mark[s[i]-'a'] {
			count[s[i]-'a']--
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] >= s[i] && count[stack[len(stack)-1]-'a'] > 1 {
			index := stack[len(stack)-1] - 'a'
			mark[index] = false
			count[index]--
			stack = stack[:len(stack)-1]
		}
		mark[s[i]-'a'] = true
		stack = append(stack, s[i])
	}
	return string(stack)
}
