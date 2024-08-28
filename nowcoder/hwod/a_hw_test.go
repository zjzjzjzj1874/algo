package hwod

import (
	"reflect"
	"testing"
)

// 最多等和不相交连续子序列
func Test_maxConSeq(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "maxConSeq", args: args{nums: []int{-1, 0, 4, -3, 6, 5, -6, 5, -7, -3}}, wantAns: 3},
		{name: "maxConSeq", args: args{nums: []int{8, 9, 1, 9, 6, 3, 9, 1, 0}}, wantAns: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxConSeq(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxConSeq() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 租车骑绿道
func Test_minBikeNum(t *testing.T) {
	type args struct {
		m      int
		n      int
		weight []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "minBikeNum", args: args{m: 3, n: 4, weight: []int{3, 2, 2, 1}}, wantAns: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minBikeNum(tt.args.m, tt.args.n, tt.args.weight); gotAns != tt.wantAns {
				t.Errorf("minBikeNum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 双十一
func Test_maxValue(t *testing.T) {
	type args struct {
		prices []int
		money  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "maxValue", args: args{prices: []int{23, 26, 36, 27}, money: 78}, wantAns: 76},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxValue(tt.args.prices, tt.args.money); gotAns != tt.wantAns {
				t.Errorf("maxValue() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 整理扑克牌
func Test_sortPoker(t *testing.T) {
	type args struct {
		poker string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "sortPoker", args: args{poker: "1 3 3 3 2 1 5"}},
		{name: "sortPoker", args: args{poker: "4 4 2 1 2 1 3 3 3 4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortPoker(tt.args.poker)
		})
	}
}

// 找数字
func Test_findNum(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantTarget [][]int
	}{
		{name: "findNum", args: args{nums: [][]int{{0, 3, 5, 4, 2}, {2, 5, 7, 8, 3}, {2, 5, 4, 2, 4}}},
			wantTarget: [][]int{{-1, 4, 2, 3, 3}, {1, 1, -1, -1, 4}, {1, 1, 2, 3, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTarget := findNum(tt.args.nums); !reflect.DeepEqual(gotTarget, tt.wantTarget) {
				t.Errorf("findNum() = %v, want %v", gotTarget, tt.wantTarget)
			}
		})
	}
}

// 找出通过车辆最多颜色
func Test_maxColor(t *testing.T) {
	type args struct {
		n    int
		cars []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "maxColor", args: args{n: 3, cars: []int{0, 1, 2, 1}}, wantAns: 2},
		{name: "maxColor", args: args{n: 2, cars: []int{0, 1, 2, 1}}, wantAns: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxColor(tt.args.n, tt.args.cars); gotAns != tt.wantAns {
				t.Errorf("maxColor() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 工作安排
func Test_maxProfit(t *testing.T) {
	type args struct {
		T     int
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "workArrange", args: args{T: 10, tasks: [][]int{{5, 10}, {4, 40}, {6, 30}, {3, 50}}}, want: 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitV1(tt.args.T, tt.args.tasks); got != tt.want {
				//if got := maxProfitWithDP(tt.args.T, tt.args.tasks); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 日志采集系统
func Test_reportLog(t *testing.T) {
	type args struct {
		logs []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "report", args: args{logs: []int{3, 4, 0, 5, 2, 6, 12, 32}}, wantAns: 4},
		{name: "report", args: args{logs: []int{3, 4, 5, 6, 7, 8, 9, 10}}, wantAns: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans1 := reportLog(tt.args.logs)
			ans2 := reportLogWithGPT(tt.args.logs)
			if ans1 != ans2 {
				t.Errorf("reportLog() = %v, want %v", ans1, ans2)
			}
		})
	}
}

// 获得完美走位
func Test_perfectStep(t *testing.T) {
	type args struct {
		step string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "perfect", args: args{step: "ASDA"}, wantAns: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := perfectStep(tt.args.step); gotAns != tt.wantAns {
				t.Errorf("perfectStep() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 任务混部
func Test_serverNum(t *testing.T) {
	type args struct {
		taskNum int
		task    [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// [[1,3,3],[2,4,1],[4,7,1],[5,6,2]]
		{name: "serverNum", args: args{taskNum: 4, task: [][]int{{1, 3, 3}, {2, 4, 1}, {4, 7, 1}, {5, 6, 2}}}, wantAns: 4},
		{name: "serverNum", args: args{taskNum: 4, task: [][]int{{1, 3, 2}, {2, 5, 3}, {4, 6, 4}}}, wantAns: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := serverNum(tt.args.taskNum, tt.args.task); gotAns != tt.wantAns {
				t.Errorf("serverNum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 单词倒叙
func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "reverse", args: args{s: "the sky is blue"}, want: "blue is sky the"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWith2Ptr(tt.args.s); got != tt.want {
				//if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 查找重复代码
func Test_longestCommonSubString(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "longestCommonSubString", args: args{s: "abcde", t: "ace"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubString(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("longestCommonSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 新员工座位安排系统
func Test_maxFriendship(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxFriend", args: args{seats: []int{0, 1, 2, 0, 2, 1, 1, 0, 1, 2, 0}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFriendship(tt.args.seats); got != tt.want {
				t.Errorf("maxFriendship() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 简单的自动曝光
func Test_autoExposure(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "auto exposure", args: args{nums: []int{0, 0}}, want: 128},
		{name: "auto exposure", args: args{nums: []int{3, 128}}, want: 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := autoExposure(tt.args.nums); got != tt.want {
				t.Errorf("autoExposure() = %v, want %v", got, tt.want)
			}
		})
	}
}
