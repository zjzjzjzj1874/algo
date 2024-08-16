package greedy

import "testing"

// 2091. 从数组中移除最大值和最小值
func Test_minimumDeletions(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "minimumDel", args: args{nums: []int{2, 10, 7, 5, 4, 1, 8, 6}}, wantAns: 5},
		{name: "minimumDel", args: args{nums: []int{41, -47, -26, 57, 82, -23, 47, 52, 42, 79, 2, 77, 0, -4, 1, -99, -57, 72, -95}}, wantAns: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumDeletionsWithOptimize(tt.args.nums); gotAns != tt.wantAns {
				//if gotAns := minimumDeletions(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minimumDeletions() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}