package main

/*
  给定一副牌，每张牌上都写着一个整数。
  此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
    每组都有 X 张牌。
    组内所有的牌上都写着相同的整数。
  仅当你可选的 X >= 2 时返回 true。
*/

func hasGroupsSizeX(deck []int) bool {
	save := make([]int, 10000)
	for _, v := range deck {
		save[v]++
	}
	for i := 2; i <= len(deck); i++ {
		if len(deck)%i == 0 {
			for _, v := range save {
				if v%i != 0 {
					goto NEXT
				}
			}
			return true
		}
	NEXT:
	}
	return false
}

func hasGroupsSizeX2(deck []int) bool {
	save := make(map[int]int)
	for _, v := range deck {
		save[v]++
	}
	curGcd := -1
	for _, v := range save {
		if curGcd == -1 {
			curGcd = v
		} else {
			curGcd = Gcd(curGcd, v)
		}
	}
	return curGcd >= 2
}
