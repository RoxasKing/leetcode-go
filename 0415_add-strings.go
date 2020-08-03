package leetcode

/*
  给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

  注意：
    num1 和num2 的长度都小于 5100.
    num1 和num2 都只包含数字 0-9.
    num1 和num2 都不包含任何前导零。
    你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/add-strings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addStrings(num1 string, num2 string) string {
	nums := make([]byte, 0, Max(len(num1), len(num2)))
	var temp int
	for num1 != "" && num2 != "" {
		sum := temp
		sum += int(num1[len(num1)-1]-'0') + int(num2[len(num2)-1]-'0')
		nums = append(nums, byte(sum%10)+'0')
		temp = sum / 10
		num1, num2 = num1[:len(num1)-1], num2[:len(num2)-1]
	}
	for num1 != "" {
		sum := temp
		sum += int(num1[len(num1)-1] - '0')
		nums = append(nums, byte(sum%10)+'0')
		temp = sum / 10
		num1 = num1[:len(num1)-1]
	}
	for num2 != "" {
		sum := temp
		sum += int(num2[len(num2)-1] - '0')
		nums = append(nums, byte(sum%10)+'0')
		temp = sum / 10
		num2 = num2[:len(num2)-1]
	}
	if temp > 0 {
		nums = append(nums, byte(temp)+'0')
	}
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return string(nums)
}
