package leetcode

import "testing"

func Test_manacher(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "manacher", args: args{s: "babad"}, want: "bab"},
		{name: "manacher", args: args{s: "cbbd"}, want: "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manacher(tt.args.s); got != tt.want {
				t.Errorf("manacher() = %v, want %v", got, tt.want)
			}
		})
	}
}
