package math

import "testing"

// 2769. 找出最大的可达成数字
func Test_theMaximumAchievableX(t *testing.T) {
	type args struct {
		num int
		t   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "theMaximumAchievableX", args: args{num: 4, t: 1}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theMaximumAchievableX(tt.args.num, tt.args.t); got != tt.want {
				t.Errorf("theMaximumAchievableX() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1201. 丑数 III
func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "uglyNumber", args: args{n: 3, a: 2, b: 3, c: 5}, want: 4},
		{name: "uglyNumber", args: args{n: 7, a: 7, b: 7, c: 7}, want: 49},
		{name: "uglyNumber", args: args{n: 1000000000, a: 2, b: 217983653, c: 336916467}, want: 1999999984},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 470. 用 Rand7() 实现 Rand10()
func Test_rand10(t *testing.T) {
	tests := []struct {
		name    string
		wantAns int
	}{
		{name: "rand10", wantAns: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := rand10(); gotAns > tt.wantAns {
				t.Errorf("rand10() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// 292. Nim 游戏
func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "canWinNim", args: args{n: 3}, want: true},
		{name: "canWinNim", args: args{n: 4}, want: false},
		{name: "canWinNim", args: args{n: 5}, want: true},
		{name: "canWinNim", args: args{n: 6}, want: true},
		{name: "canWinNim", args: args{n: 8}, want: false},
		{name: "canWinNim", args: args{n: 43}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := canWinNim(tt.args.n)
			if got := canWinNimWithRecursion(tt.args.n); got != tt.want || got1 != got1 {
				//if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
