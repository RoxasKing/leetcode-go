package main

/*
  写一个程序，输出从 1 到 n 数字的字符串表示。
    1. 如果 n 是3的倍数，输出“Fizz”；
    2. 如果 n 是5的倍数，输出“Buzz”；
    3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fizz-buzz
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fizzBuzz(n int) []string {
	var out []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out = append(out, "FizzBuzz")
		} else if i%3 == 0 {
			out = append(out, "Fizz")
		} else if i%5 == 0 {
			out = append(out, "Buzz")
		} else {
			var tmp []byte
			for j := i; j != 0; j /= 10 {
				tmp = append(tmp, byte(j%10)+'0')
			}
			for i := 0; i < len(tmp)>>1; i++ {
				tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
			}
			out = append(out, string(tmp))
		}
	}
	return out
}
