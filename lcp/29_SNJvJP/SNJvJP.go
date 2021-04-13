package main

/*
  某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。

  为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。例如当 num = 5 时，站位如图所示

  请返回位于场地坐标 [Xpos,Ypos] 的成员所持乐器编号。

  示例 1：
    输入：num = 3, Xpos = 0, Ypos = 2
    输出：3
    解释：

  示例 2：
    输入：num = 4, Xpos = 1, Ypos = 2
    输出：5
    解释：

  提示：
    1. 1 <= num <= 10^9
    2. 0 <= Xpos, Ypos < num

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/SNJvJP
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func orchestraLayout(num int, xPos int, yPos int) int {
	i := Max(Max(num-1-xPos, xPos), Max(num-1-yPos, yPos))
	u, d, l, r := num-1-i, i, num-1-i, i
	idx := (num*num - (r+1-l)*(r+1-l)) % 9
	if xPos == u {
		idx = (idx + yPos + 1 - l) % 9
		if idx == 0 {
			return 9
		}
		return idx
	} else {
		idx = (idx + r + 1 - l) % 9
		u++
	}
	if yPos == r {
		idx = (idx + xPos + 1 - u) % 9
		if idx == 0 {
			return 9
		}
		return idx
	} else {
		idx = (idx + d + 1 - u) % 9
		r--
	}
	if xPos == d {
		idx = (idx + r + 1 - yPos) % 9
		if idx == 0 {
			return 9
		}
		return idx
	} else {
		idx = (idx + r + 1 - l) % 9
		d--
	}
	idx = (idx + d + 1 - xPos) % 9
	if idx == 0 {
		return 9
	}
	return idx
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
