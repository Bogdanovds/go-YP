package main

import (
	"log"
	"net/http"
)

type MyHandler struct {
	Templ []byte
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(h.Templ)
}

func main() {
	handler1 := MyHandler{
		Templ: []byte("Hola, Mundo"),
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler1,
	}
	server.ListenAndServe()

	// вызов ListenAndServe — блокирующий, последний в программе
	// возникающие ошибки на серверных машинах пишут в системный лог,
	// а не в стандартную консоль ошибок,
	// поэтому обычно вызывают вот так
	log.Fatal(server.ListenAndServe())
}
