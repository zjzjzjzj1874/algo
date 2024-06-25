package easy

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test_max_profit", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "test_max_profit", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
		{name: "test_max_profit", args: args{prices: []int{2, 1, 4}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
