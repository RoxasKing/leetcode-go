package codinginterviews

/*
  输入一个字符串，打印出该字符串中字符的所有排列。
  你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

  限制：
    1 <= s 的长度 <= 8
*/

func permutation(s string) []string {
	if s == "" {
		return nil
	}
	bytes := []byte(s)
	var out []string
	var dfs func(int)
	dfs = func(ptr int) {
		if ptr == len(s)-1 {
			out = append(out, string(bytes))
			return
		}
		dict := make(map[byte]bool)
		for i := ptr; i < len(bytes); i++ {
			if dict[bytes[i]] {
				continue
			}
			dict[bytes[i]] = true
			bytes[ptr], bytes[i] = bytes[i], bytes[ptr]
			dfs(ptr + 1)
			bytes[ptr], bytes[i] = bytes[i], bytes[ptr]
		}
	}
	dfs(0)
	return out
}
