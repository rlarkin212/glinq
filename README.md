# Glinq
[![Go Reference](https://pkg.go.dev/badge/github.com/rlarkin212/glinq.svg)](https://pkg.go.dev/github.com/rlarkin212/glinq)

Go implemention of LINQ using generics

## Examples
---
```go
package main

import (
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
    
    //output: [a b c d]

	userAges := glinq.Select(users, func(x user) int {
		return x.age
	})

    //output: [20 12 5 90]

	maxAge := glinq.Max(userAges)

    //output: 90

    minAge := glinq.Min(userAges)

    //output: 5

	ageAverage := glinq.Average(userAges)

    //output: 31.75

	orderByAgeAscending := glinq.OrderBy(users, func(x user) int {
		return x.age
	})

    //output: [{c 5} {b 12} {a 20} {d 90}]

	orderByAgeDescending := glinq.OrderByDescending(users, func(x user) int {
		return x.age
	})

    //output: [{d 90} {a 20} {b 12} {c 5}]
}
```

