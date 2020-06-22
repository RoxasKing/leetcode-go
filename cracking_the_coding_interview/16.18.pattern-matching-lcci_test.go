package crackingthecodingintervew

import "testing"

func Test_patternMatching(t *testing.T) {
	type args struct {
		pattern string
		value   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abba", "dogcatcatdog"}, true},
		{"2", args{"abba", "dogcatcatfish"}, false},
		{"3", args{"aaaa", "dogcatcatdog"}, false},
		{"4", args{"abba", "dogdogdogdog"}, true},
		{"5", args{"ab", ""}, false},
		{"6", args{"bbbaa", "xxxxxx"}, true},
		{"7", args{"aaaaab", "xahnxdxyaahnxdxyaahnxdxyaahnxdxyaauxuhuo"}, true},
		{"8", args{"a", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := patternMatching(tt.args.pattern, tt.args.value); got != tt.want {
				t.Errorf("patternMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
