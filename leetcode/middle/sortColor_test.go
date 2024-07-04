package middle

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "sortColors", args: args{nums: []int{2, 0, 2, 1, 1, 0}}},
		{name: "sortColors", args: args{nums: []int{2, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
