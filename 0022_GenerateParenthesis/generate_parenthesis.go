package main

import "fmt"

// 思路：下一循环就是对上一循环产生的数据的再加工，如：第一次循环产生的为 () ,
// 则第二次循环就是把 () 插入到 上一次产生的 () 之中
func generateParenthesis(n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		var tmp []string
		if len(out) == 0 {
			// 如果是第一次循环，则生成初始的 ()
			tmp = append(tmp, "()")
		} else {
			// 以上一次魂环产生的 []string为基础数据，加工每一个成员
			for j := 0; j < len(out); j++ {
				// 对每一个成员生成所有可能情况
				for k := 0; k < len(out[j]); k++ {
					if k == 0 {
						// 插首位
						tmp = append(tmp, "()"+out[j])
					} else if out[j][k-1] == '(' && out[j][k] == ')' || out[j][k-1] == '(' && out[j][k] == '(' {
						// 插正向括号 或者 反向括号 的中间位
						tmp = append(tmp, out[j][:k]+"()"+out[j][k:])
					} else {
						break
					}
				}
			}
		}
		out = tmp
	}
	return out
}

func main() {
	i := 3
	fmt.Println(generateParenthesis(i))
}
