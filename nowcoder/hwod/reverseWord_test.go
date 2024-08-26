package hwod

import "testing"

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
