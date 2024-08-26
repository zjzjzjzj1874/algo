package hwod

import "testing"

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
