package monitonicStack

import "testing"

func Test_maxWidthRamp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxwidth", args: args{nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}}, want: 7},
		{name: "maxwidth", args: args{nums: []int{2, 2, 1}}, want: 1},
		{name: "maxwidth", args: args{nums: []int{6, 0, 8, 2, 1, 5}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.args.nums); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
