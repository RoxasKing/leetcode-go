package main

func largestOddNumber(num string) string {
	out := ""
	for i := range num {
		if int(num[i]-'0')&1 == 1 {
			out = num[:i+1]
		}
	}
	return out
}
