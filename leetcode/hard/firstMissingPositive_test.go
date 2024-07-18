package hard

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "missing", args: args{nums: []int{1, 2, 0}}, want: 3},
		{name: "missing", args: args{nums: []int{3, 4, -1, 1}}, want: 2},
		//{name: "missing", args: args{nums: []int{7, 8, 9, 11, 12}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositiveV1(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 135. 分发糖果
func Test_Candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "candy", args: args{ratings: []int{1, 0, 2}}, want: 5},
		{name: "candy", args: args{ratings: []int{1, 2, 2}}, want: 4},
		{name: "candy", args: args{ratings: []int{1, 3, 2, 2, 1}}, want: 7},
		{name: "candy", args: args{ratings: []int{1, 2, 87, 87, 87, 2, 1}}, want: 13},
		{name: "candy", args: args{ratings: []int{1, 3, 4, 5, 2}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
