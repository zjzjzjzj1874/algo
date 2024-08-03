package dp

import "testing"

func Test_twoEggDrop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "twoEggDrop", args: args{n: 100}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEggDrop(tt.args.n); got != tt.want {
				t.Errorf("twoEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superEggDrop(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "superEggDrop", args: args{k: 3, n: 14}, want: 4},
		{name: "superEggDrop", args: args{k: 5, n: 100}, want: 7},
		{name: "superEggDrop", args: args{k: 5, n: 200}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
