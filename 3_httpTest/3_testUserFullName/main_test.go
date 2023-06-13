package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestUser_FullName(t *testing.T) {
    type fields struct {
        FirstName string
        LastName  string
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "simple test",
            fields: fields{
                FirstName: "Misha",
                LastName:  "Popov",
            },
            want: "Misha Popov",
        },
        {
            name: "long name",
            fields: fields{
                FirstName: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz",
                LastName:  "Picasso",
            },
            want: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz Picasso",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            u := User{
                FirstName: tt.fields.FirstName,
                LastName:  tt.fields.LastName,
            }
            v := u.FullName()
            // как и в предыдущем тесте сроки сравниваются с помощью функции Equal
            assert.Equal(t, tt.want, v)
        })
    }
}