package codinginterviews

/*
  数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

  请写一个函数，求任意第n位对应的数字。

  限制：
    0 <= n < 2^31

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findNthDigit(n int) int {
	if n < 0 {
		return -1
	}
	digits := 1
	numbers := countOfIntegers(digits)
	for n >= numbers*digits {
		n -= numbers * digits
		digits++
		numbers = countOfIntegers(digits)
	}
	return digitAtIndex(n, digits)
}
func countOfIntegers(digits int) int {
	if digits == 1 {
		return 10
	}
	return 9 * Pow(10, digits-1)
}

func digitAtIndex(index, digits int) int {
	number := beginNuber(digits) + index/digits
	indexFromRight := digits - index%digits
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func beginNuber(digits int) int {
	if digits == 1 {
		return 10
	}
	return Pow(10, digits-1)
}
