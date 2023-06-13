package main

import "fmt"

// User в системе.
type User struct {
    FirstName string
    LastName  string
}

// FullName возвращает фамилию и имя человека.
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}

func main() {
    u := User{
        FirstName: "MIsha",
        LastName:  "Popov",
    }

    fmt.Println(u.FullName())
}