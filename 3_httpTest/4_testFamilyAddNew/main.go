package main

import (
    "errors"
    "fmt"
)

type Relationship string // Relationship определяет положение в семье.

const (
    Father      = Relationship("father") // Возможные роли в семье.
    Mother      = Relationship("mother")
    Child       = Relationship("child")
    GrandMother = Relationship("grandMother")
    GrandFather = Relationship("grandFather")
)

type Family struct { // Family описывает семью.
    Members map[Relationship]Person
}

type Person struct { // Person описывает конкретного человека в семье.
    FirstName string
    LastName  string
    Age       int
}

var (
    ErrRelationshipAlreadyExists = errors.New("relationship already exists") // ErrRelationshipAlreadyExists возвращает ошибку, если роль уже занята.
)

// AddNew добавляет нового члена семьи. Если в семье ещё нет людей, создаётся пустой map. Если роль уже занята, метод выдаёт ошибку.
func (f *Family) AddNew(r Relationship, p Person) error {
    if f.Members == nil {
        f.Members = map[Relationship]Person{}
    }
    if _, ok := f.Members[r]; ok {
        return ErrRelationshipAlreadyExists
    }
    f.Members[r] = p
    return nil
}

func main() {
    f := Family{}
    err := f.AddNew(Father, Person{
        FirstName: "Misha",
        LastName:  "Popov",
        Age:       56,
    })
    fmt.Println(f, err)

    err = f.AddNew(Father, Person{
        FirstName: "Drug",
        LastName:  "Mishi",
        Age:       57,
    })
    fmt.Println(f, err)
}