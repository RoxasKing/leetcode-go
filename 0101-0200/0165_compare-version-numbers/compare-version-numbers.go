package main

/*
  比较两个版本号 version1 和 version2。
  如果 version1 > version2 返回 1，如果 version1 < version2 返回 -1， 除此之外返回 0。
  你可以假设版本字符串非空，并且只包含数字和 . 字符。
   . 字符不代表小数点，而是用于分隔数字序列。
  例如，2.5 不是“两个半”，也不是“差一半到三”，而是第二版中的第五个小版本。
  你可以假设版本号的每一级的默认修订版号为 0。例如，版本号 3.4 的第一级（大版本）和第二级（小版本）修订号分别为 3 和 4。其第三级和第四级修订号均为 0。

  提示：
    版本字符串由以点 （.） 分隔的数字字符串组成。这个数字字符串可能有前导零。
    版本字符串不以点开始或结束，并且其中不会有两个连续的点。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/compare-version-numbers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)
	l1, r1, l2, r2 := 0, 1, 0, 1
	for l1 < n1 && l2 < n2 {
		for r1 < n1 && version1[r1] != '.' {
			r1++
		}
		for r2 < n2 && version2[r2] != '.' {
			r2++
		}
		if res := compare(trimZero(version1[l1:r1]), trimZero(version2[l2:r2])); res != 0 {
			return res
		}
		l1, r1, l2, r2 = r1+1, r1+2, r2+1, r2+2
	}
	if l1 < n1 && containNonZeroNumber(version1[l1:]) {
		return 1
	} else if l2 < n2 && containNonZeroNumber(version2[l2:]) {
		return -1
	}
	return 0
}

func compare(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return boolToInt(n1 > n2)
	}
	for i := 0; i < n1; i++ {
		if s1[i] != s2[i] {
			return boolToInt(s1[i] > s2[i])
		}
	}
	return 0
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return -1
}

func trimZero(s string) string {
	for len(s) != 0 && s[0] == '0' {
		s = s[1:]
	}
	return s
}

func containNonZeroNumber(s string) bool {
	for i := range s {
		if '1' <= s[i] && s[i] <= '9' {
			return true
		}
	}
	return false
}
