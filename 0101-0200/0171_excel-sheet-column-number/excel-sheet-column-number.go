package main

/*
  给定一个Excel表格中的列名称，返回其相应的列序号。

  例如，
        A -> 1
        B -> 2
        C -> 3
        ...
        Z -> 26
        AA -> 27
        AB -> 28
        ...

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/excel-sheet-column-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func titleToNumber(s string) int {
	var out int
	base := 1
	for s != "" {
		last := s[len(s)-1]
		out += base * int(last-'A'+1)
		s = s[:len(s)-1]
		base *= 26
	}
	return out
}
