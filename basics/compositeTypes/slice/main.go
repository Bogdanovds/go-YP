package main

import (
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/addElement"
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/assignment"
	"github.com/bogdanovds/go-YP/basics/compositeTypes/slice/basicOperation"
)

func main() {
	//1. Механика append
	addElement.AddElement()

	// 2. Сортировка значений в слайсе
	assignment.Assignment()

	// 3. Основные операции с элементами в слайсе
	basicoperation.BasicOperation()
}
