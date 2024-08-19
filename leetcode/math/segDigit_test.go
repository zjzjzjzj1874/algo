package math

import (
	"reflect"
	"testing"
)

// 2553. 分割数组中数字的数位
func Test_separateDigits(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{name: "separateDigits", args: args{nums: []int{100, 40}}, wantAns: []int{1, 0, 0, 4, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := separateDigits(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("separateDigits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
