package middle

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "accounts", args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"John", "johnnybravo@mail.com"},
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"Mary", "mary@mail.com"},
				}},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			//[["Alex","Alex5@m.co","Alex4@m.co","Alex0@m.co"],
			//["Ethan","Ethan3@m.co","Ethan3@m.co","Ethan0@m.co"],
			//["Kevin","Kevin4@m.co","Kevin2@m.co","Kevin2@m.co"],
			//["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe2@m.co"],
			//["Gabe","Gabe3@m.co","Gabe4@m.co","Gabe2@m.co"]]
			name: "accounts", args: args{
				accounts: [][]string{
					{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
					{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
					{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
					{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
					{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
				}},
			//[["Alex","Alex0@m.co","Alex4@m.co","Alex5@m.co"],
			//["Ethan","Ethan0@m.co","Ethan3@m.co"],
			//["Gabe","Gabe0@m.co","Gabe2@m.co","Gabe3@m.co","Gabe4@m.co"],
			//["Kevin","Kevin2@m.co","Kevin4@m.co"]]...
			want: [][]string{
				{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				fmt.Println("want:", tt.want)
				fmt.Println("got:", got)
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
