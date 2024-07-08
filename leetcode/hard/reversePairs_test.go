package hard

import "testing"

func Test_reversePairs(t *testing.T) {
	type args struct {
		record []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "reversePairs", args: args{record: []int{9, 7, 5, 4, 6}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.record); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
