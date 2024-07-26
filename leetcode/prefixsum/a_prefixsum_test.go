package prefixsum

import (
	"fmt"
	"testing"
)

// 930. 和相同的二元子数组
func Test_numSubArraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "numSubArrays", args: args{nums: []int{1, 0, 1, 0, 1}, goal: 2}, wantAns: 4},
		{name: "numSubArrays", args: args{nums: []int{0, 0, 0, 0, 0}, goal: 0}, wantAns: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSubArraysWithSum(tt.args.nums, tt.args.goal); gotAns != tt.wantAns {
				t.Errorf("numSubArraysWithSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 303. 区域和检索 - 数组不可变
func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		nums []int
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantAns int
	}{
		{name: "SumRange", fields: fields{nums: []int{-2, 0, 3, -5, 2, -1}}, args: args{left: 0, right: 2}, wantAns: 1},
		{name: "SumRange", fields: fields{nums: []int{-2, 0, 3, -5, 2, -1}}, args: args{left: 2, right: 5}, wantAns: -1},
		{name: "SumRange", fields: fields{nums: []int{-2, 0, 3, -5, 2, -1}}, args: args{left: 0, right: 5}, wantAns: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &NumArray{
				nums: tt.fields.nums,
			}
			if gotAns := a.SumRange(tt.args.left, tt.args.right); gotAns != tt.wantAns {
				t.Errorf("SumRange() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}

	//307. 区域和检索 - 数组可修改
	t.Run("307.区域和检索 - 数组可修改", func(t *testing.T) {
		a := &NumArray{
			nums: []int{1, 3, 5},
		}
		fmt.Println(a.SumRange(0, 2))
		a.Update(1, 2)
		fmt.Println(a.SumRange(0, 2))
	})
}
