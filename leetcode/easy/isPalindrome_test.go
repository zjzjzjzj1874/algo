package easy

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{name: "palindrome", args: args{head: &ListNode{
		//	Val: 1,
		//	Next: &ListNode{
		//		Val: 2,
		//		Next: &ListNode{
		//			Val: 2,
		//			Next: &ListNode{
		//				Val:  1,
		//				Next: nil,
		//			},
		//		},
		//	},
		//}}, want: true},
		{name: "palindrome", args: args{head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
