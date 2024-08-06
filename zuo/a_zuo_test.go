package zuo

import "testing"

// 最小染色
func Test_minPaint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "minPaint", args: args{s: "RGRGR" /*RRRGG*/}, wantAns: 2},
		{name: "minPaint", args: args{s: "RGGGRGGRGRRGRRGR" /*RRRGG*/}, wantAns: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minPaintWithON(tt.args.s); gotAns != tt.wantAns {
				//if gotAns := minPaint(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minPaint() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 最小带子数量
func Test_minBagNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "minBagNumber", args: args{n: 64}, want: 8},
		//{name: "minBagNumber", args: args{n: 62}, want: 8},
		//{name: "minBagNumber", args: args{n: 8}, want: 1},
		//{name: "minBagNumber", args: args{n: 6}, want: 1},
		//{name: "minBagNumber", args: args{n: 12}, want: 2},
		//{name: "minBagNumber", args: args{n: 14}, want: 2},
		//{name: "minBagNumber", args: args{n: 16}, want: 2},
		{name: "minBagNumber", args: args{n: 1000}, want: 125},
		{name: "minBagNumber", args: args{n: 1002}, want: 126},
		{name: "minBagNumber", args: args{n: 1001}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBagNumber(tt.args.n); got != tt.want {
				t.Errorf("minBagNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 计算最小公倍数
func Test_lcm(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "lcm", args: args{x: 6, y: 8}, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcm(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 求0-40亿之间没有出现过的数字,假设数组中的元素不重复
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
