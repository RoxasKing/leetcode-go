package main

/*
  编写一个算法来判断一个数 n 是不是快乐数。
  「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
  如果 n 是快乐数就返回 True ；不是，则返回 False 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/happy-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isHappy(n int) bool {
	for {
		if n == 1 {
			return true
		}
		if n < 10 && n%2 == 0 {
			return false
		}
		var newN int
		for n != 0 {
			newN += (n % 10) * (n % 10)
			n /= 10
		}
		for newN%10 == 0 {
			newN /= 10
		}
		n = newN
	}
}

func isHappy2(n int) bool {
	dict := map[int]struct{}{
		4:   {},
		16:  {},
		37:  {},
		58:  {},
		89:  {},
		145: {},
		42:  {},
		20:  {},
	}
	for {
		if n == 1 {
			return true
		}
		if _, ok := dict[n]; ok {
			return false
		}
		var newN int
		for n != 0 {
			newN += (n % 10) * (n % 10)
			n /= 10
		}
		n = newN
	}
}
