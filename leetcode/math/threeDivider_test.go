package math

import "testing"

// 1952. 三除数
func Test_isThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "isThree", args: args{n: 9}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isThree(tt.args.n); got != tt.want {
				t.Errorf("isThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
