package main

/*
  Given a string expression of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. You may return the answer in any order.

  Example 1:
    Input: expression = "2-1-1"
    Output: [0,2]
    Explanation:
      ((2-1)-1) = 0
      (2-(1-1)) = 2

  Example 2:
    Input: expression = "2*3-4*5"
    Output: [-34,-14,-10,-10,10]
    Explanation:
      (2*(3-(4*5))) = -34
      ((2*3)-(4*5)) = -14
      ((2*(3-4))*5) = -10
      (2*((3-4)*5)) = -10
      (((2*3)-4)*5) = 10

  Constraints:
    1. 1 <= expression.length <= 20
    2. expression consists of digits and the operator '+', '-', and '*'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/different-ways-to-add-parentheses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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
