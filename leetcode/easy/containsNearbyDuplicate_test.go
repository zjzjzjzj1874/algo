package easy

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "containsNearbyDuplicate", args: args{nums: []int{1, 2, 3, 1}, k: 3}, want: true},
		{name: "containsNearbyDuplicate", args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2}, want: false},
		{name: "containsNearbyDuplicate", args: args{nums: []int{99, 99}, k: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicateHash(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
