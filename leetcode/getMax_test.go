package leetcode

import "testing"

func TestGetMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "getMax", args: args{nums: []int{10, 10, 7, 7, 8, 8, 9, 9, 1, 1, 5, 5}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMax(tt.args.nums); got != tt.want {
				t.Errorf("GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
