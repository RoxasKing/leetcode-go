package main

import (
	"My_LeetCode_In_Go/Algorithm"
	"fmt"
)

func main() {
	fmt.Println(Algorithm.Fibonacci1(10))
	fmt.Println(Algorithm.Fibonacci2(20))
	data := []int{9, 7, 23, 1, 45, 3, 11, 99, 32}
	fmt.Println(data)
	Algorithm.QuickSort(data)
	fmt.Println(data)
	data2 := []int{6, 7, 13, 1, 45, 3, 11, 29, 32}
	fmt.Println(data2)
	Algorithm.BubbleSort(data2)
	fmt.Println(data2)

	var balance = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	queen(balance, 0)
}

func queen(a [8]int, cur int) {
	// 如果 cur 超过数组 a 的下标，则打印
	if cur == len(a) {
		fmt.Println(a)
		fmt.Println()
		return
	}
	for i := 0; i < len(a); i++ {
		a[cur] = i
		flag := true
		// 判断 i 是否合法
		for j := 0; j < cur; j++ {
			ab := i - a[j]
			temp := 0
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			// 如果下标为 j 的数等于 i，或者差值的绝对值与下标差值相同
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}
		// 如果当前 i 合法，则继续计算下一列
		if flag {
			queen(a, cur+1)
		}
	}
}
