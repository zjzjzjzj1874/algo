package hard

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "longest", args: args{s: ")))(()"}, want: 2},
		//{name: "longest", args: args{s: ")()())"}, want: 4},
		//{name: "longest", args: args{s: "()())"}, want: 4},
		//{name: "longest", args: args{s: "((()))())"}, want: 8},
		{name: "longest", args: args{s: "((()()))())"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesStack(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
