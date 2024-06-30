package middle

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "partition", args: args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  2,
									Next: nil,
								},
							},
						},
					},
				},
			},
			x: 3,
		}, want: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
