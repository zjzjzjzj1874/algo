package middle

import (
	"reflect"
	"testing"
)

func Test_findKthLargestWithHeapSort(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "findKthLargest", args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestWithHeapSort(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestWithHeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{name: "smallestK", args: args{arr: []int{1, 3, 5, 7, 2, 4, 6, 8}, k: 4}, wantAns: []int{1, 2, 3, 4}},
		{name: "smallestK", args: args{arr: []int{
			62577, -220, -8737, -22, -6, 59956, 5363, -16699, 0, -10603, 64, -24528, -4818, 96, 5747, 2638, -223, 37663, -390, 35778, -4977, -3834, -56074, 7, -76, 601, -1712, -48874, 31, 3, -9417, -33152, 775, 9396, 60947, -1919, 683, -37092, -524, -8, 1458, 80, -8, 1, 7, -355, 9, 397, -30, -21019, -565, 8762, -4, 531, -211, -23702, 3, 3399, -67, 64542, 39546, 52500, -6263, 4, -16, -1, 861, 5134, 8, 63701, 40202, 43349, -4283, -3, -22721, -6, 42754, -726, 118, 51, 539, 790, -9972, 41752, 0, 31, -23957, -714, -446, 4, -61087, 84, -140, 6, 53, -48496, 9, -15357, 402, 5541, 4, 53936, 6, 3, 37591, 7, 30, -7197, -26607, 202, 140, -4, -7410, 2031, -715, 4, -60981, 365, -23620, -41, 4, -2482, -59, 5, -911, 52, 50068, 38, 61, 664, 0, -868, 8681, -8, 8, 29, 412},
			k: 131},
			wantAns: []int{-61087, -60981, -56074, -48874, -48496, -37092, -33152, -26607, -24528, -23957, -23702, -23620, -22721, -21019, -16699, -15357, -10603, -9972, -9417, -8737, -7410, -7197, -6263, -4977, -4818, -4283, -3834, -2482, -1919, -1712, -911, -868, -726, -715, -714, -565, -524, -446, -390, -355, -223, -220, -211, -140, -76, -67, -59, -41, -30, -22, -16, -8, -8, -8, -6, -6, -4, -4, -3, -1, 0, 0, 0, 1, 3, 3, 3, 4, 4, 4, 4, 4, 5, 6, 6, 7, 7, 7, 8, 8, 9, 9, 29, 30, 31, 31, 38, 51, 52, 53, 61, 64, 80, 84, 96, 118, 140, 202, 365, 397, 402, 412, 531, 539, 601, 664, 683, 775, 790, 861, 1458, 2031, 2638, 3399, 5134, 5363, 5541, 5747, 8681, 8762, 9396, 35778, 37591, 37663, 39546, 40202, 41752, 42754, 43349, 50068, 52500}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := smallestK(tt.args.arr, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("smallestK() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
