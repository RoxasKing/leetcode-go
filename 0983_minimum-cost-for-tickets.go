package leetcode

/*
  在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。
  火车票有三种不同的销售方式：
    一张为期一天的通行证售价为 costs[0] 美元；
    一张为期七天的通行证售价为 costs[1] 美元；
    一张为期三十天的通行证售价为 costs[2] 美元。
  通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张为期 7 天的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。
  返回你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-cost-for-tickets
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mincostTickets(days []int, costs []int) int {
	memo := make([]int, 366)
	exists := make(map[int]bool)
	for _, day := range days {
		exists[day] = true
	}
	return dpMincostTickets(1, &costs, &memo, &exists)
}

func dpMincostTickets(day int, costs, memo *[]int, dict *map[int]bool) int {
	if day > 365 {
		return 0
	}
	if (*memo)[day] > 0 {
		return (*memo)[day]
	}
	if (*dict)[day] {
		(*memo)[day] = Min(
			Min(
				dpMincostTickets(day+1, costs, memo, dict)+(*costs)[0],
				dpMincostTickets(day+7, costs, memo, dict)+(*costs)[1],
			),
			dpMincostTickets(day+30, costs, memo, dict)+(*costs)[2],
		)
	} else {
		(*memo)[day] = dpMincostTickets(day+1, costs, memo, dict)
	}
	return (*memo)[day]
}

func dpMincostTickets11(days []int, costs []int) int {
	dayMap := []int{1, 7, 30}
	memo := make([]int, days[len(days)-1]+1)
	exists := make(map[int]struct{})
	for _, day := range days {
		exists[day] = struct{}{}
	}
	var dp func(int) int
	dp = func(day int) int {
		if day > days[len(days)-1] {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if _, ok := exists[day]; ok {
			memo[day] = 1<<31 - 1
			for i := range costs {
				memo[day] = Min(memo[day], dp(day+dayMap[i])+costs[i])
			}
		} else {
			memo[day] = dp(day + 1)
		}
		return memo[day]
	}
	return dp(1)
}

func mincostTickets2(days []int, costs []int) int {
	daysMap := []int{1, 7, 30}
	memo := make([]int, 366)
	return dpMincostTickets2(0, &days, &costs, &daysMap, &memo)
}

func dpMincostTickets2(index int, days, costs, daysMap, memo *[]int) int {
	if index >= len(*days) {
		return 0
	}
	if (*memo)[index] > 0 {
		return (*memo)[index]
	}
	(*memo)[index] = 1<<31 - 1
	for i := 0; i < 3; i++ {
		j := index
		for j < len(*days) && (*days)[j] < (*days)[index]+(*daysMap)[i] {
			j++
		}
		(*memo)[index] = Min(
			(*memo)[index],
			dpMincostTickets2(j, days, costs, daysMap, memo)+(*costs)[i],
		)
	}
	return (*memo)[index]
}

func mincostTickets22(days []int, costs []int) int {
	daysMap := []int{1, 7, 30}
	memo := make([]int, len(days)+1)
	var dp func(int) int
	dp = func(index int) int {
		if index >= len(days) {
			return 0
		}
		if memo[index] > 0 {
			return memo[index]
		}
		memo[index] = 1<<31 - 1
		for i := 0; i < 3; i++ {
			j := index
			for j < len(days) && days[j] < days[index]+daysMap[i] {
				j++
			}
			memo[index] = Min(memo[index], dp(j)+costs[i])
		}
		return memo[index]
	}
	return dp(0)
}

func mincostTickets3(days []int, costs []int) int {
	daysMap := []int{1, 7, 30}
	exists := make(map[int]struct{}, len(days))
	for _, day := range days {
		exists[day] = struct{}{}
	}
	dp := make([]int, days[len(days)-1]+1)
	for i := 1; i <= days[len(days)-1]; i++ {
		if _, ok := exists[i]; ok {
			dp[i] = dp[i-1] + costs[0]
			for j := 1; j < 3; j++ {
				curCosts := costs[j]
				if i > daysMap[j] {
					curCosts += dp[i-daysMap[j]]
				}
				dp[i] = Min(dp[i], curCosts)
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[days[len(days)-1]]
}

func mincostTickets33(days []int, costs []int) int {
	daysMap := []int{1, 7, 30}
	newDays := make([]int, len(days))
	for i := range days {
		newDays[i] = days[i] - days[0] + 1
	}
	exists := make(map[int]int, len(days))
	for i := range newDays {
		exists[newDays[i]] = i + 1
	}
	dp := make([]int, len(days)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + costs[0]
		for j := 1; j < 3; j++ {
			curCost := costs[j]
			if newDays[i-1] > daysMap[j] {
				for k := newDays[i-1] - daysMap[j]; k >= 1; k-- {
					if index, ok := exists[k]; ok {
						curCost += dp[index]
						break
					}
				}
			}
			dp[i] = Min(dp[i], curCost)
		}
	}
	return dp[len(days)]
}
