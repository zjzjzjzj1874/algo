package leetcode

import "testing"

func Test_selectionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "selectionSort", args: args{nums: []int{10, 10, 7, 7, 8, 8, 9, 9, 1, 1, 5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectionSort(tt.args.nums)
		})
	}
}
