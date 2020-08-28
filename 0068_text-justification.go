package main

/*
  给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
  你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
  要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
  文本的最后一行应为左对齐，且单词之间不插入额外的空格。
说明:
  单词是指由非空格字符组成的字符序列。
  每个单词的长度大于 0，小于等于 maxWidth。
  输入单词数组 words 至少包含一个单词。
*/

func fullJustify(words []string, maxWidth int) []string {
	out := []string{}
	if len(words) < 1 {
		return out
	}
	tmp := []string{words[0]}
	tmpLen := len(words[0])
	i := 1
	for i < len(words) {
		if tmpLen+1+len(words[i]) <= maxWidth {
			tmp = append(tmp, words[i])
			tmpLen += 1 + len(words[i])
			i++
		} else {
			// 当前行总空格数量
			spaceNum := maxWidth - tmpLen + len(tmp) - 1
			elem := tmp[0]
			if len(tmp) > 1 {
				j := 1
				for ; j <= spaceNum%(len(tmp)-1); j++ {
					for k := 0; k < spaceNum/(len(tmp)-1)+1; k++ {
						elem += " "
					}
					elem += tmp[j]
				}
				for ; j < len(tmp); j++ {
					for k := 0; k < spaceNum/(len(tmp)-1); k++ {
						elem += " "
					}
					elem += tmp[j]
				}
			} else {
				for k := 0; k < spaceNum; k++ {
					elem += " "
				}
			}
			out = append(out, elem)
			tmp = []string{words[i]}
			tmpLen = len(words[i])
			i++
		}
	}
	spaceNum := maxWidth - tmpLen + len(tmp) - 1
	elem := tmp[0]
	for j := 1; j < len(tmp); j++ {
		elem += " " + tmp[j]
		spaceNum--
	}
	for k := 0; k < spaceNum; k++ {
		elem += " "
	}
	out = append(out, elem)

	return out
}

func fullJustify2(words []string, maxWidth int) []string {
	var out []string
	for {
		var tmp []string
		var tmpStrLen int
		for len(words) > 0 && tmpStrLen+len(tmp)+len(words[0]) <= maxWidth {
			tmp = append(tmp, words[0])
			tmpStrLen += len(words[0])
			words = words[1:]
		}
		var str string
		if len(words) == 0 {
			for i := range tmp {
				str += tmp[i]
				if i < len(tmp)-1 {
					str += " "
				}
			}
			for len(str) < maxWidth {
				str += " "
			}
			out = append(out, str)
			break
		}
		if len(tmp)-1 == 0 {
			str += tmp[0]
			for i := 0; i < maxWidth-tmpStrLen; i++ {
				str += " "
			}
		} else {
			spaceAll := maxWidth - tmpStrLen
			var gap string
			for i := 0; i < spaceAll/(len(tmp)-1); i++ {
				gap += " "
			}
			cur := spaceAll % (len(tmp) - 1)
			for i := range tmp {
				str += tmp[i]
				if i == len(tmp)-1 {
					break
				}
				str += gap
				if cur > 0 {
					str += " "
					cur--
				}
			}
		}
		out = append(out, str)
	}
	return out
}
