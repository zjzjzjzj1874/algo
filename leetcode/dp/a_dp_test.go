package dp

import "testing"

// LCR 098. 不同路径
func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "uniquePath", args: args{m: 3, n: 7}, want: 28},
		{name: "uniquePath", args: args{m: 7, n: 3}, want: 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 099. 最小路径和
func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minPathSum", args: args{grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

// LCR 090. 打家劫舍 II
func Test_Rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "rob", args: args{nums: []int{2, 3, 2}}, want: 3},
		{name: "rob", args: args{nums: []int{1, 2, 3, 1}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 091. 粉刷房子
func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minCost", args: args{costs: [][]int{{7, 6, 2}}}, want: 2},
		{name: "minCost", args: args{costs: [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.costs); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCoin(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minCoins", args: args{coins: []int{2, 3, 100}, target: 5}, want: 2},
		{name: "minCoins", args: args{coins: []int{7, 4, 5, 2, 3, 2}, target: 10}, want: 2},
		{name: "minCoins", args: args{coins: []int{17, 17, 10, 5, 5}, target: 10}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := minCoins(tt.args.coins, tt.args.target); got != tt.want {
			if got := minCoinsWithDP(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("minCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
