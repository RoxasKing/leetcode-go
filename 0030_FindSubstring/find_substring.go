package main

import "fmt"

func findSubstring(s string, words []string) []int {
	sizes := len(s)
	out := make([]int, 0, sizes)
	sizews := len(words)
	if sizes == 0 || sizews == 0 {
		return out
	}
	sizew := len(words[0])
	if sizes < sizews*sizew {
		return out
	}
	// 记录 words 中每个单词总共出现的次数
	recw := make(map[string]int, sizews)
	for _, w := range words {
		if len(w) == sizew {
			recw[w]++
		} else {
			return out
		}
	}
	// 记录当前单词出现次数
	var recc map[string]int
	// 记录符合条件的单词数
	var count int
	// 选择器最左和最右单词起始位置
	left, right := 0, 0
	// 移位循环，循环次数视单词长度而定
	for i := 0; i < sizew; i++ {
		// 初始化选择器
		left, right = i, i
		// 初始化单词出现次数统计
		recc = make(map[string]int)
		// 初始化符合次数统计
		count = 0
		// 若当前可选择最大范围超过 words 字符串总长
		for sizes-left >= sizews*sizew {
			word := s[right : right+sizew]
			if max, ok := recw[word]; !ok {
				// 如果单词不在记录中,从当前点往后一个单词位置重新开始检查
				left, right = right+sizew, right+sizew
				// 初始化单词出现次数统计
				recc = make(map[string]int)
				// 初始化符合次数统计
				count = 0
			} else if recc[word]+1 > max {
				// 如果单词出现次数超过words记录
				// left位置单词记录数减 1
				recc[s[left:left+sizew]]--
				// 符合条件单词数减 1
				count--
				// 选择器左边位置向右移动一个单词长度
				left += sizew
			} else {
				// 如果符合增加单词记录的条件,单词计数加 1
				recc[word]++
				// 符合条件单词数加 1
				count++
				// 选择器左边位置向右移动一个单词位置
				right += sizew
				// 如果符合条件单词计数等于 words 长度
				if count == sizews {
					out = append(out, left)
					// left位置单词记录数减 1
					recc[s[left:left+sizew]]--
					// 符合条件单词数减 1
					count--
					// 选择器左边位置向右移动一个单词长度
					left += sizew
				}
			}
		}
	}
	return out
}

func main() {
	//s := "barfoothefoobarman"
	//words := []string{"foo", "bar"}
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	//s := ""
	//words := []string{}
	fmt.Println(findSubstring(s, words))
}
