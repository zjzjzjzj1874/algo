package bytedance

import "testing"

func Test_change(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{name: "change", args: args{n: 200}, wantAns: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := change(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("change() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
