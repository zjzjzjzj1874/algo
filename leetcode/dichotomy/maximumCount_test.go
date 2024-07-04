package dichotomy

import "testing"

func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maximum", args: args{nums: []int{-2, -1, -1, 1, 2, 3}}, want: 3},
		{name: "maximum", args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2}}, want: 3},
		{name: "maximum", args: args{nums: []int{5, 20, 66, 1314}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
