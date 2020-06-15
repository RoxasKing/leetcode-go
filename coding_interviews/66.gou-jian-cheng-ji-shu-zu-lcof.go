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
	if len(a) < 2 {
		return nil
	}
	l, r := make([]int, len(a)), make([]int, len(a))
	l[0], r[len(a)-1] = 1, 1
	for i := 1; i < len(a); i++ {
		l[i] = l[i-1] * a[i-1]
	}
	for i := len(a) - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i+1]
	}
	out := make([]int, len(a))
	for i := range out {
		out[i] = l[i] * r[i]
	}
	return out
}
