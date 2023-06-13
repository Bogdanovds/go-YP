package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}
	response, err := client.Get("https://ya.ru")
	client.Timeout = time.Second * 1
	fmt.Println(response, err)
}
