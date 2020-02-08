package leetcode

/*
  罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
    字符          数值
    I             1
    V             5
    X             10
    L             50
    C             100
    D             500
    M             1000
  例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做 XXVII, 即为 XX + V + II 。
  通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
    I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
    X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
    C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
  给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
*/

func intToRoman(num int) string {
	var (
		out []rune
		i   int
	)
	for num != 0 {
		switch {
		case num >= 1000:
			i = num / 1000
			num -= i * 1000
			for ; i > 0; i-- {
				out = append(out, 'M')
			}
		case num >= 900:
			num -= 900
			out = append(out, 'C', 'M')
		case num >= 500:
			num -= 500
			out = append(out, 'D')
		case num >= 400:
			num -= 400
			out = append(out, 'C', 'D')
		case num >= 100:
			i = num / 100
			num -= i * 100
			for ; i > 0; i-- {
				out = append(out, 'C')
			}
		case num >= 90:
			num -= 90
			out = append(out, 'X', 'C')
		case num >= 50:
			num -= 50
			out = append(out, 'L')
		case num >= 40:
			num -= 40
			out = append(out, 'X', 'L')
		case num >= 10:
			i = num / 10
			num -= i * 10
			for ; i > 0; i-- {
				out = append(out, 'X')
			}
		case num >= 9:
			num -= 9
			out = append(out, 'I', 'X')
		case num >= 5:
			num -= 5
			out = append(out, 'V')
		case num >= 4:
			num -= 4
			out = append(out, 'I', 'V')
		case num >= 1:
			for ; num > 0; num-- {
				out = append(out, 'I')
			}
		}
	}
	return string(out)
}
