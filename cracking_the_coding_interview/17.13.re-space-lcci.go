package crackingthecodingintervew

import "math"

/*
  哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。

  注意：本题相对原题稍作改动，只需返回未识别的字符数

  提示：
    0 <= len(sentence) <= 1000
    dictionary中总字符数不超过 150000。
    你可以认为dictionary和sentence中只包含小写字母。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/re-space-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
		for j := i; j >= 1; j-- {
			index := int(sentence[j-1]-'a') + 1
			hashVal = (hashVal*BASE + index) % P
			if hashVals[hashVal] {
				dp[i] = Min(dp[i], dp[j-1])
			}
		}
	}
	return dp[len(sentence)]
}
