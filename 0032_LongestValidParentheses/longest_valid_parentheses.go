package main

import "fmt"

func longestValidParentheses(s string) int {
	// 定义一个数组 slice 记录字符
	var slice []byte
	// 定义一个数组 position 记录字符的下标
	var position []int
	var out int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 如果是 '(' ,则压入栈中
			slice = append(slice, s[i])
			position = append(position, i)
		} else if s[i] == ')' && len(slice) != 0 && slice[len(slice)-1] == '(' {
			// 如果是 ’)' ，且数组 slice 长度大于 0 ，且 slice 最后一个元素 是 '('
			// 数组 slice 和 position 长度都减 1
			slice = slice[:len(slice)-1]
			position = position[:len(position)-1]
		} else {
			// 否则将 下标压入 position 中
			position = append(position, i)
		}
	}
	if len(position) != 0 {
		// 如果数组 position 长度为 1 ，且此下标为 0 或者是 len(s)-1
		if len(position) == 1 && (position[0] == 0 || position[0] == len(s)-1) {
			return len(s) - 1
		}
		// 初始化指针
		pointer := -1
		// 循环遍历存储下标的数组
		for k, v := range position {
			// size用于记录当前最大间隔
			size := 0
			if pointer != -1 {
				size = v - pointer - 1
			} else {
				size = v
			}
			// 修改输出最大值
			if size > out {
				out = size
			}
			// 指针指向当前 v
			pointer = v
			// 如果遍历到末尾，且最大间隔大于当前 out
			if k == len(position)-1 && out < len(s)-1-pointer {
				// 修改 out 值
				out = len(s) - 1 - pointer
			}
		}
		return out
	} else {
		// 如果 position 数组为 0 ，直接输出整个数组长度
		return len(s)
	}
}

func main() {
	//s := "()()("
	s := "())())"
	fmt.Println(longestValidParentheses(s))
}
