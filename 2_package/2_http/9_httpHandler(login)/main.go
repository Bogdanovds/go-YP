package main

import (
	"fmt"
	"net/http"
)

var form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/login" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // проверяем, каким методом получили запрос
	case "POST": // если методом POST
		login := r.FormValue("login")
		password := r.FormValue("password")

		if !Auth(login, password) { // проверяем пароль вспомогательной функцией
			w.Header().Set("Content-Type", "text/plain; charset=utf-8") // если пароль не верен, указываем код ошибки в заголовке
			w.WriteHeader(401)                                          // пишем в тело ответа
			fmt.Fprintln(w, "Wrong password")
			return
		}

		//AuthorisedHandler(w, r) // при успешной авторизации обрабатываем запрос например, передаём другому обработчику

	default: // в остальных случаях предлагаем форму авторизации
		fmt.Fprint(w, form)
	}
}

// Auth — вспомогательная функция авторизации
// за пределами урока реализация может выглядеть так
func Auth(l, p string) bool {
	pass, ok := Logins[l]
	return ok && p == pass
}

var Logins = make(map[string]string)

func main() {
	http.HandleFunc("/", Login)       // маршрутизация запросов обработчику
	http.ListenAndServe(":8080", nil) // запуск сервера с адресом localhost, порт 8080

}
