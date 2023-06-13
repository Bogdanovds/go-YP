package main

import "net/http"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
	http.HandleFunc("/", HelloWorld)// маршрутизация запросов обработчику
	server := &http.Server{ // конструируем свой сервер
		Addr: "mydomain.com:80",
	}
	server.ListenAndServe()
}
