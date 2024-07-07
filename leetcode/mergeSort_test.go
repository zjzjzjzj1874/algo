package leetcode

import "testing"

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "MergeSort", args: args{nums: []int{10, 10, 7, 7, 8, 8, 9, 9, 1, 1, 5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.nums)
		})
	}
}
