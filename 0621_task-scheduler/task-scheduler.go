package main

import "sort"

/*
  给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

  然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

  你需要计算完成所有任务所需要的最短时间。

  提示：
    任务的总个数为 [1, 10000]。
    n 的取值范围为 [0, 100]。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/task-scheduler
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}
	sort.Ints(count)
	intervals := count[25] - 1
	slotCounts := intervals * n
	for i := 24; i >= 0 && count[i] > 0; i-- {
		slotCounts -= Min(count[i], intervals)
	}
	if slotCounts > 0 {
		return slotCounts + len(tasks)
	}
	return len(tasks)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
