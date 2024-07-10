package easy

import "testing"

func Test_isPalindromeWithNum(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isPalindromeWith", args: args{x: 12321}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeWithNum(tt.args.x); got != tt.want {
				t.Errorf("isPalindromeWithNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
