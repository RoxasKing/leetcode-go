package leetcode

import "testing"

func Test_getMaxRepetitions(t *testing.T) {
	type args struct {
		s1 string
		n1 int
		s2 string
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"1", args{"acb", 4, "ab", 2}, 2},
		// {"2", args{"aaa", 3, "aa", 1}, 4},
		// {"3", args{"baba", 11, "baab", 1}, 7},
		{
			"4",
			args{
				"phqghumeaylnlfdxfircvscxggbwkfnqduxwfnfozvsrtkjprepggxrpnrvystmwcysyycqpevikeffmznimkkasvwsrenzkycxf", 1000000,
				"xtlsgypsfadpooefxzbcoejuvpvaboygpoeylfpbnpljvrvipyamyehwqnqrqpmxujjloovaowuxwhmsncbxcoksfzkvatxdknly", 100,
			},
			303,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxRepetitions(tt.args.s1, tt.args.n1, tt.args.s2, tt.args.n2); got != tt.want {
				t.Errorf("getMaxRepetitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
