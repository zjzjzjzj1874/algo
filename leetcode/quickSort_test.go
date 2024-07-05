package leetcode

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		//{name: "quickSort", args: args{nums: []int{10, 10, 7, 7, 8, 8, 9, 9, 1, 1, 5, 5}}},
		//{name: "quickSort", args: args{nums: []int{5, 2, 4, 1, 3, 6, 0}}},
		//{name: "quickSort", args: args{nums: []int{5, 2, 3, 1}}},
		{name: "quickSort", args: args{nums: []int{110, 100, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums)
		})
	}
}
