package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// string 转换成 []byte, 容易取得相应位上的具体值
	bsi := []byte(num1)
	bsj := []byte(num2)

	// temp 的长度只可能为 len(num1)+len(num2) 或 len(num1)+len(num2)-1 先选长的，免得不够位
	temp := make([]int, len(num1)+len(num2))
	for i := 0; i < len(bsi); i++ {
		for j := 0; j < len(bsj); j++ {
			// 每个位上的乘积，可以直接存入 temp 中相应的位置
			temp[i+j+1] += int(bsi[i]-'0') * int(bsj[j]-'0')
		}
	}

	// 统一处理进位
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}

	// num1 和 num2 较小的时候， temp 的首位为 0
	// 为避免输出结果以 0 开头， 需要去掉 temp 的 0 首位
	if temp[0] == 0 {
		temp = temp[1:]
	}

	// 转换结果
	// temp 选用为 []int, 而不是 []byte， 是因为 Go 中， byte 的基础结构是 uint8， 最大值为 255
	// 不考虑进位的话， temp 会溢出
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}

	return string(res)
}

func main() {
	num1 := "12"
	num2 := "12"
	fmt.Println(multiply(num1, num2))
}
