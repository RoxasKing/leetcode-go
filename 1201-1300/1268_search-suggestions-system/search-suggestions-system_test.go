package main

import (
	"reflect"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1",
			args{[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"},
			[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			"2",
			args{[]string{"havana"}, "havana"},
			[][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			"3",
			args{[]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"},
			[][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"},
				{"bags"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := suggestedProducts(tt.args.products, tt.args.searchWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("suggestedProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
