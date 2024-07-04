package dichotomy

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "findPeakElement", args: args{nums: []int{1, 2, 3, 1}}, want: 2},
		{name: "findPeakElement", args: args{nums: []int{1, 2, 4, 6, 13}}, want: 4},
		{name: "findPeakElement", args: args{nums: []int{1, 2, 1, 3, 5, 6, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElementV1(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
