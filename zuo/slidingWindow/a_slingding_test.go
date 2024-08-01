package slidingWindow

import "testing"

// 给定一个有序数组arr，代表数轴上从左到右有n个点arr[0]、arr[1]...arr[n-1]，给定一个正数L,代表一根长度为L的绳子,求绳子最多能覆盖其中的几个点。
func Test_ropeMaxPoint(t *testing.T) {
	type args struct {
		nums    []int
		ropeLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ropeMaxPoint", args: args{nums: []int{2, 4, 8, 9, 12, 17}, ropeLen: 5}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ropeMaxPoint(tt.args.nums, tt.args.ropeLen); got != tt.want {
				t.Errorf("ropeMaxPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
