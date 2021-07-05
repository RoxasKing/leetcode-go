package main

import "math"

// Trie
func respace(dictionary []string, sentence string) int {
	type Trie struct {
		next  [26]*Trie
		isEnd bool
	}
	trie := new(Trie)
	for _, word := range dictionary {
		cur := trie
		for i := len(word) - 1; i >= 0; i-- {
			index := word[i] - 'a'
			if cur.next[index] == nil {
				cur.next[index] = new(Trie)
			}
			cur = cur.next[index]
		}
		cur.isEnd = true
	}
	dp := make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		cur := trie
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			index := sentence[j] - 'a'
			if cur.next[index] == nil {
				break
			} else if cur.next[index].isEnd {
				dp[i] = Min(dp[i], dp[j])
			}
			cur = cur.next[index]
		}
	}
	return dp[len(sentence)]
}

// Rabin-Karp
func respace2(dictionary []string, sentence string) int {
	const (
		P    = math.MaxInt32
		BASE = 41
	)
	getHash := func(s string) int {
		var hashVal int
		for i := len(s) - 1; i >= 0; i-- {
			index := int(s[i]-'a') + 1
			hashVal = (hashVal*BASE + index) % P
		}
		return hashVal
	}
	hashVals := make(map[int]bool)
	for _, word := range dictionary {
		hashVals[getHash(word)] = true
	}
	dp := make([]int, len(sentence)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = len(sentence)
	}
	for i := 1; i <= len(sentence); i++ {
		var hashVal int
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			index := int(sentence[j]-'a') + 1
			hashVal = (hashVal*BASE + index) % P
			if hashVals[hashVal] {
				dp[i] = Min(dp[i], dp[j])
			}
		}
	}
	return dp[len(sentence)]
}
