package easy

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{name: "addBinary", args: args{a: "11", b: "1"}, want: "100"},
		//{name: "addBinary", args: args{a: "1011", b: "1010"}, want: "10101"},
		{name: "addBinary", args: args{a: "1", b: "111"}, want: "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
