package middle

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "search", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		{name: "search", args: args{nums: []int{1}, target: 0}, want: -1},
		{name: "search", args: args{nums: []int{5, 1, 2, 3, 4}, target: 1}, want: 1},
		{name: "search", args: args{nums: []int{8, 1, 2, 3, 4, 5, 6, 7}, target: 6}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
