package main

/*
  所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
  编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/repeated-dna-sequences
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Rabin-Karp
func findRepeatedDnaSequences(s string) []string {
	n, L := len(s), 10
	if n <= L {
		return nil
	}
	dict := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	nums := make([]int, n)
	for i := range nums {
		nums[i] = dict[s[i]]
	}
	base := 4
	baseL := Pow(base, L)
	var hash int
	for i := 0; i < L; i++ {
		hash = hash*base + nums[i]
	}
	mark := map[int]bool{hash: true}
	save := make(map[int]bool)
	var out []string
	for i := 1; i < n+1-L; i++ {
		hash = hash*base - nums[i-1]*baseL + nums[i-1+L]
		if mark[hash] && !save[hash] {
			out = append(out, s[i:i+L])
			save[hash] = true
		}
		mark[hash] = true
	}
	return out
}

func Pow(base, t int) int {
	out := 1
	for t > 0 {
		if t&1 == 1 {
			out *= base
		}
		base *= base
		t >>= 1
	}
	return out
}
