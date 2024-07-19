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

// 面试题 16.10. 生存人数
func Test_maxAliveYear(t *testing.T) {
	type args struct {
		birth []int
		death []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxAlive", args: args{birth: []int{1900, 1901, 1950}, death: []int{1948, 1951, 2000}}, want: 1901},
		{name: "maxAlive", args: args{
			birth: []int{1972, 1908, 1915, 1957, 1960, 1948, 1912, 1903, 1949, 1977, 1900, 1957, 1934, 1929, 1913, 1902, 1903, 1901},
			death: []int{1997, 1932, 1963, 1997, 1983, 2000, 1926, 1962, 1955, 1997, 1998, 1989, 1992, 1975, 1940, 1903, 1983, 1969}},
			want: 1960},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAliveYear(tt.args.birth, tt.args.death); got != tt.want {
				t.Errorf("maxAliveYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
