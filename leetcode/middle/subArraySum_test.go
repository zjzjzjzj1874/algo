package middle

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "", args: args{nums: []int{1, 1, 1}, k: 2}, want: 2},
		//{name: "", args: args{nums: []int{1, 2, 3}, k: 3}, want: 2},
		{name: "", args: args{nums: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumV1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
