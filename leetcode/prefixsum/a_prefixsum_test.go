package prefixsum

import (
	"testing"
)

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
}
