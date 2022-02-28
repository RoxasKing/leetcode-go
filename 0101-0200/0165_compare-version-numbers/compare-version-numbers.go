package main

// Difficulty:
// Medium

func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)
	for i, j := -1, -1; i < n1 || j < n2; {
		v1, v2 := 0, 0
		for i++; i < n1 && version1[i] != '.'; i++ {
			v1 = v1*10 + int(version1[i]-'0')
		}
		for j++; j < n2 && version2[j] != '.'; j++ {
			v2 = v2*10 + int(version2[j]-'0')
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}
	return 0
}
