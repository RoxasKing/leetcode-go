package main

/*
  由 n 个连接的字符串 s 组成字符串 S，记作 S = [s,n]。例如，["abc",3]=“abcabcabc”。
  如果我们可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。例如，根据定义，"abc" 可以从 “abdbec” 获得，但不能从 “acbbe” 获得。
  现在给你两个非空字符串 s1 和 s2（每个最多 100 个字符长）和两个整数 0 ≤ n1 ≤ 106 和 1 ≤ n2 ≤ 106。现在考虑字符串 S1 和 S2，其中 S1=[s1,n1] 、S2=[s2,n2] 。
  请你找出一个可以满足使[S2,M] 从 S1 获得的最大整数 M 。
*/

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1cnt, index, s2cnt := 0, 0, 0
	recall := make(map[int][]int)
	var preLoop, inLoop []int
	for {
		s1cnt++
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					s2cnt, index = s2cnt+1, 0
				}
			}
		}
		if s1cnt == n1 {
			return s2cnt / n2
		}
		if _, ok := recall[index]; ok {
			preLoop = recall[index]
			inLoop = []int{s1cnt - recall[index][0], s2cnt - recall[index][1]}
			break
		} else {
			recall[index] = []int{s1cnt, s2cnt}
		}
	}
	ans := preLoop[1] + (n1-preLoop[0])/inLoop[0]*inLoop[1]
	rest := (n1 - preLoop[0]) % inLoop[0]
	for i := 0; i < rest; i++ {
		for j := range s1 {
			if s1[j] == s2[index] {
				index++
				if index == len(s2) {
					ans, index = ans+1, 0
				}
			}
		}
	}
	return ans / n2
}
