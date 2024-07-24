package backtrack

import (
	"reflect"
	"testing"
)

// LCR 083. 全排列
func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "permute", args: args{nums: []int{1, 2, 3}}, wantAns: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := permute(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("permute() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// LCR 084. 全排列II
func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{name: "permuteUnique", args: args{nums: []int{1, 1, 2}}, wantAns: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := permuteUnique(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("permuteUnique() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
