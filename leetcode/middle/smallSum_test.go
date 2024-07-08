package middle

import "testing"

func Test_smallSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "small sum", args: args{nums: []int{1, 3, 4, 2, 5}}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallSum(tt.args.nums); got != tt.want {
				t.Errorf("smallSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
