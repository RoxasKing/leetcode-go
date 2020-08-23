package main

/*
  统计所有小于非负整数 n 的质数的数量。
*/

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	var count int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
