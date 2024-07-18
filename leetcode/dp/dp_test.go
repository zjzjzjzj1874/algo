package dp

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minCost", args: args{cost: []int{10, 15, 20}}, want: 15},
		{name: "minCost", args: args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "climbStairs", args: args{n: 3}, want: 3},
		{name: "climbStairs", args: args{n: 8}, want: 34},
		{name: "climbStairs", args: args{n: 10}, want: 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsWithOptimize(tt.args.n); got != tt.want {
				//if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
