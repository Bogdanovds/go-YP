package main

import (
    "context"
    "fmt"
)

func main() {
    fmt.Println(context.Background() == context.Background())
}