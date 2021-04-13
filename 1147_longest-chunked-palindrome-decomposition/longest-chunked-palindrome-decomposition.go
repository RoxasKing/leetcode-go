package main

/*
  You are given a string text. You should split it to k substrings (subtext1, subtext2, ..., subtextk) such that:
    1. subtexti is a non-empty string.
    2. The concatenation of all the substrings is equal to text (i.e., subtext1 + subtext2 + ... + subtextk == text).
    3. subtexti == subtextk - i + 1 for all valid values of i (i.e., 1 <= i <= k).
  Return the largest possible value of k.

  Example 1:
    Input: text = "ghiabcdefhelloadamhelloabcdefghi"
    Output: 7
    Explanation: We can split the string on "(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)".

  Example 2:
    Input: text = "merchant"
    Output: 1
    Explanation: We can split the string on "(merchant)".

  Example 3:
    Input: text = "antaprezatepzapreanta"
    Output: 11
    Explanation: We can split the string on "(a)(nt)(a)(pre)(za)(tpe)(za)(pre)(a)(nt)(a)".

  Example 4:
    Input: text = "aaa"
    Output: 3
    Explanation: We can split the string on "(a)(a)(a)".

  Constraints:
    1. 1 <= text.length <= 1000
    2. text consists only of lowercase English characters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-chunked-palindrome-decomposition
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Dynamic Programming
func longestDecomposition(text string) int {
	if text == "" {
		return 0
	}
	n := len(text)
	out := 1
	for size := 1; size*2 <= n; size++ {
		if text[:size] != text[n-size:] {
			continue
		}
		i := 0
		for (i+1)*size <= n-(i+1)*size && text[i*size:(i+1)*size] == text[:size] &&
			text[i*size:(i+1)*size] == text[n-(i+1)*size:n-i*size] {
			i++
		}
		i--
		out = Max(out, 2*(i+1)+longestDecomposition(text[(i+1)*size:n-(i+1)*size]))
		break
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS + Dynamic Programming + Hash
func longestDecomposition2(text string) int {
	if text == "" {
		return 0
	}
	out := 1
	lHash, rHash := 0, 0
	n := len(text)
	l, r := 0, n-1
	base := 1
	for l < r {
		lHash = lHash*26 + int(text[l]-'a')
		rHash += base * int(text[r]-'a')
		base *= 26
		if lHash != rHash {
			l++
			r--
			continue
		}
		size := l + 1
		i := 0
		for (i+1)*size <= n-(i+1)*size && text[i*size:(i+1)*size] == text[:size] &&
			text[i*size:(i+1)*size] == text[n-(i+1)*size:n-i*size] {
			i++
		}
		i--
		out = Max(out, 2*(i+1)+longestDecomposition2(text[(i+1)*size:n-(i+1)*size]))
		break
	}
	return out
}
