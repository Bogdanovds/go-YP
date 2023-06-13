package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // конструируем контекст с Timeout
	defer cancel() // функция cancel() позволяет при необходимости остановить операции
	
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil) // собираем запрос с контекстом
	client := &http.Client{} // конструируем клиент
	resp, err := client.Do(req) // отправляем запрос

	if err !=nil {
		return
	}

	fmt.Println(resp)
}
