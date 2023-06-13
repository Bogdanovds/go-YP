package main

import (
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/addElement"
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/assignment"
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/basicOperation"
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/tasks"
)

func main() {
	//1. Механика append.
	addElement.AddElement()

	// 2. Сортировка значений в слайсе.
	assignment.Assignment()

	// 3. Основные операции с элементами в слайсе.
	basicoperation.BasicOperation()

	// 4. Давить элемент в слайс так, чтобы изменения повлияли на базовый массив.
	tasks.ChangeBaseArr()

	// 5. Удали пользователей возрастом выше 40 лет из слайса, выведи результат в консоль.
	tasks.Filter()
}
