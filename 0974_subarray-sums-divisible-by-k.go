package main

/*
  给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。

  提示：
    1 <= A.length <= 30000
    -10000 <= A[i] <= 10000
    2 <= K <= 10000
*/

// Hash + Congruence theorem
func subarraysDivByK(A []int, K int) int {
	var out, sum int
	sumMap := map[int]int{0: 1}
	for _, a := range A {
		sum += a
		modulus := (sum%K + K) % K // if sum%K < 0 , add K makes it > 0
		out += sumMap[modulus]
		sumMap[modulus]++
	}
	return out
}

func subarraysDivByK2(A []int, K int) int {
	var out, sum int
	sumMap := map[int]int{0: 1}
	for _, a := range A {
		sum += a
		modulus := (sum%K + K) % K
		sumMap[modulus]++
	}
	for _, s := range sumMap {
		out += s * (s - 1) >> 1
	}
	return out
}
