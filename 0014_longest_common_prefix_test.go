package My_LeetCode_In_Go

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 4", args{[]string{"aa", "a"}}, "a"},
		{"test 3", args{[]string{"", ""}}, ""},
		{"test 1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"test 2", args{[]string{"dog", "racecar", "car"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
