package leetcode

import "testing"

func TestBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "BubbleSort", args: args{nums: []int{7, 8, 9, 1, 8, 9, 10, 10, 7, 1, 5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.nums)
		})
	}
}
