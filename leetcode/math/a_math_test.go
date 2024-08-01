package math

import "testing"

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
