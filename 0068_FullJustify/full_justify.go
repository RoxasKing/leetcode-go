package main

import "fmt"

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
				// 若当前行不止一个单词
				j := 1
				// 5/3时，前两个空格按 2 算
				for ; j <= spaceNum%(len(tmp)-1); j++ {
					for k := 0; k < spaceNum/(len(tmp)-1)+1; k++ {
						elem += " "
					}
					elem += tmp[j]
				}
				// 后一个空格按 1 算
				for ; j < len(tmp); j++ {
					for k := 0; k < spaceNum/(len(tmp)-1); k++ {
						elem += " "
					}
					elem += tmp[j]
				}
			} else {
				// 只有一个单词时
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
	// 最后一行的单词之间空格数皆为 1
	for j := 1; j < len(tmp); j++ {
		elem += " " + tmp[j]
		spaceNum--
	}
	// 剩余空格全部补在最后
	for k := 0; k < spaceNum; k++ {
		elem += " "
	}
	out = append(out, elem)

	return out
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	newWords := fullJustify(words, maxWidth)
	for i := 0; i < len(newWords); i++ {
		fmt.Println(newWords[i])
	}
}
