package middle

import "testing"

func Test_singleNumber137(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "singleNumber137", args: args{nums: []int{2, 2, 3, 2}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber137(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber137() = %v, want %v", got, tt.want)
			}
		})
	}
}
