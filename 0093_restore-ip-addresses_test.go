package leetcode

import (
	"reflect"
	"testing"
)

func Test_restoreIPAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{"0000"}, []string{"0.0.0.0"}},
		{"", args{"25525511135"}, []string{"255.255.11.135", "255.255.111.35"}},
		{
			"",
			args{"172162541"},
			[]string{
				"17.216.25.41",
				"17.216.254.1",
				"172.16.25.41",
				"172.16.254.1",
				"172.162.5.41",
				"172.162.54.1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIPAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIPAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
