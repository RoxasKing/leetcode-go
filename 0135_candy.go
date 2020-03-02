package leetcode

/*
  老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
  你需要按照以下要求，帮助老师给这些孩子分发糖果：
    每个孩子至少分配到 1 个糖果。
    相邻的孩子中，评分高的孩子必须获得更多的糖果。
  那么这样下来，老师至少需要准备多少颗糖果呢？
*/

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	count := func(num int) int {
		return (num * (num + 1)) >> 1
	}
	compare := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}
	var candies int
	var up, down int
	var pre, cur int
	for i := 1; i < len(ratings); i++ {
		pre, cur = cur, compare(ratings[i], ratings[i-1])
		if pre > 0 && cur == 0 || pre < 0 && cur >= 0 {
			candies += count(up) + count(down) + Max(up, down)
			up, down = 0, 0
		}
		if cur > 0 {
			up++
		} else if cur < 0 {
			down++
		} else {
			candies++
		}
	}
	candies += count(up) + count(down) + Max(up, down) + 1
	return candies
}

// two array
func candy2(ratings []int) int {
	var sum int
	ltor, rtol := make([]int, len(ratings)), make([]int, len(ratings))
	for i := range ratings {
		ltor[i] = 1
		rtol[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			ltor[i] = ltor[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rtol[i] = rtol[i+1] + 1
		}
	}
	for i := range ratings {
		sum += Max(ltor[i], rtol[i])
	}
	return sum
}

// one array
func candy3(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range ratings {
		candies[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	sum := candies[len(candies)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = Max(candies[i], candies[i+1]+1)
		}
		sum += candies[i]
	}
	return sum
}
