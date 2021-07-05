package main

func maxDepthexp1afterSplit(seq string) []int {
	out := make([]int, len(seq))
	depth := 0
	for i := range seq {
		if seq[i] == '(' {
			depth++
			out[i] = depth & 1
		}
		if seq[i] == ')' {
			out[i] = depth & 1
			depth--
		}
	}
	return out
}

func maxDepthexp1afterSplit2(seq string) []int {
	out := make([]int, len(seq))
	for i := range seq {
		if seq[i] == '(' {
			out[i] = 1 - i&1
		} else {
			out[i] = i & 1
		}
	}
	return out
}
