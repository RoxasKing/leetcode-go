package main

/*
  For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].

  Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.

  Example 1:
    Input: A = [1,2,0,0], K = 34
    Output: [1,2,3,4]
    Explanation: 1200 + 34 = 1234

  Example 2:
    Input: A = [2,7,4], K = 181
    Output: [4,5,5]
    Explanation: 274 + 181 = 455

  Example 3:
    Input: A = [2,1,5], K = 806
    Output: [1,0,2,1]
    Explanation: 215 + 806 = 1021

  Example 4:
    Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
    Output: [1,0,0,0,0,0,0,0,0,0,0]
    Explanation: 9999999999 + 1 = 10000000000

  Note：
    1. 1 <= A.length <= 10000
    2. 0 <= A[i] <= 9
    3. 0 <= K <= 10000
    4. If A.length > 1, then A[0] != 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/add-to-array-form-of-integer
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addToArrayForm(A []int, K int) []int {
	reverse(A)
	B := []int{}
	for K > 0 {
		B = append(B, K%10)
		K /= 10
	}
	C := make([]int, 0, Max(len(A), len(B))+1)
	remain := 0
	for len(A) > 0 || len(B) > 0 {
		if len(A) > 0 {
			remain += A[0]
			A = A[1:]
		}
		if len(B) > 0 {
			remain += B[0]
			B = B[1:]
		}
		C = append(C, remain%10)
		remain /= 10
	}
	if remain > 0 {
		C = append(C, remain)
	}
	reverse(C)
	return C
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
