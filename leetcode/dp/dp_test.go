package dp

import "testing"

// LCR 088. 使用最小花费爬楼梯
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

// 70. 爬楼梯
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

// LCR 103. 零钱兑换
func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "coinChange", args: args{coins: []int{1, 2, 5}, amount: 11}, want: 3},
		{name: "coinChange", args: args{coins: []int{2}, amount: 3}, want: -1},
		{name: "coinChange", args: args{coins: []int{1}, amount: 0}, want: 0},
		{name: "coinChange", args: args{coins: []int{186, 419, 83, 408}, amount: 6249}, want: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
