package tasks

import "fmt"

type User struct {
	Name string
	Age  int
}

func Filter() {
	users := []User{
		{
			Name: "Alice",
			Age:  21,
		},
		{
			Name: "John",
			Age:  34,
		},
		{
			Name: "Alexander",
			Age:  45,
		},
		{
			Name: "Ivan",
			Age:  13,
		},
		{
			Name: "Denis",
			Age:  44,
		},
		{
			Name: "Mary",
			Age:  26,
		},
		{
			Name: "Rose",
			Age:  41,
		},
	}

	for i := len(users) - 1; i >= 0; i-- {
		if users[i].Age > 40 {
			users = users[:i+copy(users[i:], users[i+1:])]
		}
	}

	fmt.Println(users)
}
