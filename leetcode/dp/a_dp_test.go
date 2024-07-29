package dp

import "testing"

// 2684. 矩阵中移动的最大次数
func Test_maxMoves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		//[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]
		{name: "maxMoves", args: args{grid: [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}}, wantAns: 3},
		{name: "maxMoves", args: args{grid: [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}}, wantAns: 0},
		//,[],[],[],[]]
		{name: "maxMoves", args: args{grid: [][]int{
			//{187, 167, 209, 251, 152, 236, 263, 128, 135},
			//{267, 249, 251, 285, 73,  204, 70,  207, 74},
			//{189, 159, 235, 66,  84,  89,  153, 111, 189},
			//{120, 81,  210, 7,   2,   231, 92,  128, 218},
			//{193, 131, 244, 293, 284, 175, 226, 205, 245},
			{187, 167, 209, 251, 152, 236, 263, 128, 135},
			{267, 249, 251, 285, 73, 204, 70, 207, 74},
			{189, 159, 235, 66, 84, 89, 153, 111, 189},
			{120, 81, 210, 7, 2, 231, 92, 128, 218},
			{193, 131, 244, 293, 284, 175, 226, 205, 245},
		}}, wantAns: 3},
		// [0 0 2 3 0 5 6 7 8]
		// [0 1 2 3 4 5 0 7 0]
		// [0 1 2 0 4 5 6 7 8]
		// [0 0 2 0 0 5 6 7 8]
		// [0 1 2 3 4 5 6 7 8]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxMoves(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("maxMoves() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 1289. 下降路径最小和 II
func Test_minFallingPathSumHard(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		//[[-37,51,-36,34,-22],[82,4,30,14,38],[-68,-52,-92,65,-85],[-49,-3,-77,8,-19],[-60,-71,-21,-62,-73]]
		{name: "minFallingPathSumHard", args: args{
			grid: [][]int{
				{-37, 51, -36, 34, -22},     // -37 0
				{82, 4, 30, 14, 38},         // 4 1
				{-68, -52, -92, 65, -85},    // -92 2
				{-49, -3, -77, 8, -19},      // -49 0
				{-60, -71, -21, -62, -73}}}, // -73 4
			wantAns: -268},
		{name: "minFallingPathSumHard", args: args{
			grid: [][]int{
				{1, 2, 3}, //
				{4, 5, 6}, // 4 1
				{7, 8, 9}, // -92 2
			},
		},
			wantAns: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minFallingPathSumHard(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("minFallingPathSumHard() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 931. 下降路径最小和
func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minFallingPathSum", args: args{matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 100. 三角形最小路径和
func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "minimumTotal", args: args{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
