package main

import (
	"encoding/json"
	"net/http"
)

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	subj := Subj{ // собираем данные
		Product: "Milk",
		Price:   90,
	}
	resp, err := json.Marshal(subj) // кодируем JSON
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json") // устанавливаем заголовок Content-Typec для передачи клиенту информации, кодированной в JSON
	w.WriteHeader(http.StatusOK)                       // устанавливаем статус-код 200
	w.Write(resp)                                      // пишем тело ответа
}

func main() {
	http.HandleFunc("/", JSONHandler) // маршрутизация запросов обработчику
	http.ListenAndServe(":8080", nil) // запуск сервера с адресом localhost, порт 8080

}
