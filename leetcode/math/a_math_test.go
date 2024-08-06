package math

import "testing"

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
