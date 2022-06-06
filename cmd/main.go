package main

import (
	"fmt"

	"github.com/rlarkin212/glinq"
)

type user struct {
	name string
	age  int
}

func main() {
	users := []user{
		{
			name: "a",
			age:  20,
		},
		{
			name: "b",
			age:  12,
		},
		{
			name: "c",
			age:  5,
		},
		{
			name: "d",
			age:  90,
		},
	}

	userNames := glinq.Select(users, func(x user) string {
		return x.name
	})
	fmt.Println(userNames)
	//output: [a b c d]

	userAges := glinq.Select(users, func(x user) int {
		return x.age
	})
	fmt.Println(userAges)
	//output: [20 12 5 90]

	maxAge := glinq.Max(userAges)
	fmt.Println(maxAge)
	//output: 90

	minAge := glinq.Min(userAges)
	fmt.Println(minAge)
	//output: 5

	ageAverage := glinq.Average(userAges)
	fmt.Println(ageAverage)
	//output: 31.75

	orderByAgeAscending := glinq.OrderBy(users, func(x user) int {
		return x.age
	})
	fmt.Println(orderByAgeAscending)
	//output: [{c 5} {b 12} {a 20} {d 90}]

	orderByAgeDescending := glinq.OrderByDescending(users, func(x user) int {
		return x.age
	})
	fmt.Println(orderByAgeDescending)
	//output: [{d 90} {a 20} {b 12} {c 5}]
}
