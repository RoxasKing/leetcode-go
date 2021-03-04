package main

import "testing"

func Test_countMatches(t *testing.T) {
	type args struct {
		items     [][]string
		ruleKey   string
		ruleValue string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}},
				"color",
				"silver",
			},
			1,
		},
		{
			"2",
			args{
				[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}},
				"type",
				"phone",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatches(tt.args.items, tt.args.ruleKey, tt.args.ruleValue); got != tt.want {
				t.Errorf("countMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
