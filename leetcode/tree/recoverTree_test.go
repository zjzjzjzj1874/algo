package tree

import "testing"

func Test_findTwoSwapped(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "findTwoSwapped", args: args{nums: []int{1, 2, 7, 4, 5, 6, 3, 8}}, want1: 3, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findTwoSwapped(tt.args.nums)
			if got != tt.want {
				t.Errorf("findTwoSwapped() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findTwoSwapped() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "rob",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   3,
								Left:  &TreeNode{},
								Right: nil},
							Right: nil},
						Right: nil},
					Right: nil,
				}},
			want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.args.root)
			if got != tt.want {
				t.Errorf("findTwoSwapped() got = %v, want %v", got, tt.want)
			}
		})
	}
}
