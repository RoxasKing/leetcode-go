package main

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
		{
			"1",
			args{"25525511135"},
			[]string{
				"255.255.11.135",
				"255.255.111.35",
			},
		},
		{
			"2",
			args{"0000"},
			[]string{"0.0.0.0"},
		},
		{
			"3",
			args{"1111"},
			[]string{"1.1.1.1"},
		},
		{
			"4",
			args{"010010"},
			[]string{"0.10.0.10", "0.100.1.0"},
		},
		{
			"5",
			args{"101023"},
			[]string{
				"1.0.10.23",
				"1.0.102.3",
				"10.1.0.23",
				"10.10.2.3",
				"101.0.2.3"},
		},
		{
			"6",
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
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIPAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
