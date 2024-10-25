package greedy

import "testing"

// 910. 最小差值 II
func Test_smallestRangeII(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smallest range", args: args{nums: []int{0, 10}, k: 2}, want: 6},
		{name: "smallest range", args: args{nums: []int{1, 3, 6}, k: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeII(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestRangeII() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 334. 递增的三元子序列
func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "increaseTriplet", args: args{nums: []int{20, 100, 10, 12, 5, 13}}, want: true},
		{name: "increaseTriplet", args: args{nums: []int{5, 4, 3, 2, 1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1647. 字符频次唯一的最小删除次数
func Test_minDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "minDeletions", args: args{s: "ceabaacb"}, wantAns: 2},
		{name: "minDeletions", args: args{s: "aaabbbcc"}, wantAns: 2},
		{name: "minDeletions", args: args{s: "abcabc"}, wantAns: 3},
		{name: "minDeletions", args: args{s: "bbcebab"}, wantAns: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDeletions(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minDeletions() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

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
