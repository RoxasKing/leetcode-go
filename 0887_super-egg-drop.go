package leetcode

/*
  你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
  每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
  你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
  每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
  你的目标是确切地知道 F 的值是多少。
  无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
*/

func superEggDrop(K int, N int) int {
	dict := make([][]int, N+1)
	for i := range dict {
		dict[i] = make([]int, K+1)
	}
	return dpSuperEggDrop(N, K, &dict)
}

func dpSuperEggDrop(N, K int, dict *[][]int) int {
	if N == 0 || K == 1 {
		return N
	}
	if val := (*dict)[N][K]; val != 0 {
		return val
	}
	res := 1<<31 - 1
	l, r := 1, N
	for l <= r {
		mid := (l + r) / 2
		broken := dpSuperEggDrop(mid-1, K-1, dict)
		notBroken := dpSuperEggDrop(N-mid, K, dict)
		if broken > notBroken {
			r = mid - 1
			res = Min(res, broken+1)
		} else {
			l = mid + 1
			res = Min(res, notBroken+1)
		}
	}
	(*dict)[N][K] = res
	return res
}

func superEggDrop2(K int, N int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	for i := 1; i < N; i++ {
		for j := 1; j <= K; j++ {
			dp[i][j] = dp[i-1][j-1] + 1 + dp[i-1][j]
			if dp[i][j] >= N {
				return i
			}
		}
	}
	return N
}

func superEggDrop3(K int, N int) int {
	dp := make([]int, K+1)
	for n := 1; n < N; n++ {
		for k := K; k >= 1; k-- {
			if k != 1 {
				dp[k] += dp[k-1] + 1
			} else {
				dp[k] = n
			}
			if dp[k] >= N {
				return n
			}
		}
	}
	return N
}

func superEggDrop4(K int, N int) int {
	T := 1
	for recurSuperEggDrop(T, K) < N {
		T++
	}
	return T
}

func recurSuperEggDrop(T, K int) int {
	if T == 1 || K == 1 {
		return T
	}
	return recurSuperEggDrop(T-1, K-1) + 1 + recurSuperEggDrop(T-1, K)
}
