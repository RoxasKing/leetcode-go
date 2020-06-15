package codinginterviews

/*
  给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

  提示：
    所有元素乘积之和不会溢出 32 位整数
    a.length <= 100000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func constructArr(a []int) []int {
	out := make([]int, len(a))
	var countZero, zeroIndex int
	product := 1
	for i, num := range a {
		if num != 0 {
			product *= num
			continue
		}
		countZero++
		if countZero > 1 {
			return out
		}
		zeroIndex = i
	}
	if countZero == 1 {
		out[zeroIndex] = product
	} else {
		for i := range a {
			out[i] = product / a[i]
		}
	}
	return out
}
