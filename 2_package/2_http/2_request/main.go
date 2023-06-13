package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "https://research.swtch.com/interfaces", nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Accept", "application/json")

	// Запрос в представлении HTTP
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestDump))
}
