package zuo

import "testing"

func Test_findMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "missing number", args: args{nums: []int{0, 1, 2, 3, 5, 6, 7, 8, 9}}, want: 4},
		{name: "missing number", args: args{nums: []int{0, 1, 3, 4, 5, 6}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingNumberDichotomy(tt.args.nums); got != tt.want {
				t.Errorf("findMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
