package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Три основных варианта content-Type

	// 1. "application/json"
	// Данные передаются в формате JSON. В качестве io.Reader используется bytes.Buffer. В заголовках "Content-Type" устанавливает "application/json"
	func() {
		var body = []byte(`{"message":"Hello"}`)
		url := "golang.org"
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Println(request)
	}()

	// 2. "multipart/form-data"
	// В такой контейнер браузеры запаковывают HTML <form>. Удобно загружать файлы на сервер
	func() {
		url := "golang.org"

		file, _ := os.Open("example.txt")
		defer file.Close()

		body := &bytes.Buffer{}                                         // конструируем буфер
		writer := multipart.NewWriter(body)                             // на основе буфера конструируем multipart.Writer из пакета mime/multipart
		part, err := writer.CreateFormFile("uploadfile", "example.txt") // готовим форму для отправки файла на сервер
		if err != nil {
			return
		}

		_, err = io.Copy(part, file) // копируем файл в форму. multipart.Writer отформатирует данные и запишет в предоставленный буфер
		if err != nil {
			return
		}
		writer.Close()

		request, err := http.NewRequest(http.MethodPost, url, body) // запрос
		if err != nil {
			return
		}

		request.Header.Add("Content-Type", writer.FormDataContentType()) // заголовок запроса
		fmt.Println(request)
	}()

	// 3. "application/x-www-form-urlencoded"
	// Это простой способ добавить к запросу параметры в формате «ключ-значение».
	// Данные кодируются в виде строки key1=value1&key2=value2.
	// При GET-запросе параметры указываются после адреса: example.com/endpoint?key1=value1&key2=value2.
	func() {
		
		data := url.Values{} // готовим контейнер для данных. используем тип url.Values из пакета net/url
		data.Set("key1", "value1") // устанавливаем данные
		data.Set("key2", "value2")
		endpoint := "example.com/endpoint?key1=value1&key2=value2"
		request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode())) // конструируем запрос
		if err != nil {
			return
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")// устанавливаем заголовки
		request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		fmt.Println(request)
	}()
}
