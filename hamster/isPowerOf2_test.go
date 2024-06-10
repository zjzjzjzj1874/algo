package hamster

import "testing"

func Test_isPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{n: 12}, want: false},
		{name: "", args: args{n: 8}, want: true},
		{name: "", args: args{n: 6}, want: false},
		{name: "", args: args{n: 4}, want: true},
		{name: "", args: args{n: 2}, want: true},
		{name: "", args: args{n: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("isPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
