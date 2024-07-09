package leetcode

import "testing"

func TestRadixSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "radixSort", args: args{nums: []int{170, 45, 75, 90, 802, 24, 2, 66}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RadixSort(tt.args.nums)
		})
	}
}
