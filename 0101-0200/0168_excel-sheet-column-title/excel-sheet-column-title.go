package main

/*
  Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

  For example:
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...

  Example 1:
    Input: columnNumber = 1
    Output: "A"

  Example 2:
    Input: columnNumber = 28
    Output: "AB"

  Example 3:
    Input: columnNumber = 701
    Output: "ZY"

  Example 4:
    Input: columnNumber = 2147483647
    Output: "FXSHRXW"

  Constraints:
    1 <= columnNumber <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/excel-sheet-column-title
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convertToTitle(columnNumber int) string {
	arr := []byte{}
	for ; columnNumber > 0; columnNumber /= 26 {
		i := columnNumber%26 - 1
		if i < 0 {
			i += 26
			columnNumber -= 26
		}
		arr = append(arr, byte(i)+'A')
	}
	for i := 0; i < len(arr)>>1; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}
