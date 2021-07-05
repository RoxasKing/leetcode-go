package main

func gcdOfStrings(str1 string, str2 string) string {
	for str2 != "" {
		if len(str1) < len(str2) {
			str1, str2 = str2, str1
		}
		pre2 := str2
		for str1 != "" && str2 != "" && str1[0] == str2[0] {
			str1 = str1[1:]
			str2 = str2[1:]
		}
		if str2 != "" {
			return ""
		}
		str1, str2 = pre2, str1
	}
	return str1
}
