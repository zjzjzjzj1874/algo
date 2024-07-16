package monitonicStack

import "testing"

func Test_nextGreaterElement556(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "nextGreaterElement556", args: args{n: 12}, want: 21},
		{name: "nextGreaterElement556", args: args{n: 21}, want: -1},
		{name: "nextGreaterElement556", args: args{n: 31425}, want: 31452},
		{name: "nextGreaterElement556", args: args{n: 230241}, want: 230412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement556(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement556() = %v, want %v", got, tt.want)
			}
		})
	}
}
