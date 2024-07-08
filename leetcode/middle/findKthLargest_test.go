package middle

import "testing"

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
