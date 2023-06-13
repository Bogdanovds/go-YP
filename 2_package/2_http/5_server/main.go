package main

import (
    "net/http"
)

// HelloWorld — обработчик запроса.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
    http.HandleFunc("/", HelloWorld) // маршрутизация запросов обработчику
    http.ListenAndServe(":8080", nil)// запуск сервера с адресом localhost, порт 8080
}