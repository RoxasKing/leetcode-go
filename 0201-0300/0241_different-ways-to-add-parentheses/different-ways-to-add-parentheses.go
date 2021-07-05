package main

// Tags:
// Dynamic Programming

func diffWaysToCompute(expression string) []int {
	nums := []int{}
	ops := []byte{}
	for i := 0; i < len(expression); i++ {
		num := 0
		for ; i < len(expression) && '0' <= expression[i] && expression[i] <= '9'; i++ {
			num = num*10 + int(expression[i]-'0')
		}
		nums = append(nums, num)
		if i < len(expression) {
			ops = append(ops, expression[i])
		}
	}

	n := len(nums)
	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, n)
		f[i][i] = []int{nums[i]}
	}

	for l := 2; l <= n; l++ {
		for j := 0; j <= n-l; j++ {
			for k := j + 1; k < j+l; k++ {
				for _, num1 := range f[j][k-1] {
					for _, num2 := range f[k][j+l-1] {
						switch ops[k-1] {
						case '+':
							f[j][j+l-1] = append(f[j][j+l-1], num1+num2)
						case '-':
							f[j][j+l-1] = append(f[j][j+l-1], num1-num2)
						case '*':
							f[j][j+l-1] = append(f[j][j+l-1], num1*num2)
						}
					}
				}
			}
		}
	}

	return f[0][n-1]
}
