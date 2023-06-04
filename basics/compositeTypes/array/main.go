package main

import (
	"fmt"
	enumFor "github.com/bogdanovds/go-YP/basics/compositeTypes/array/enum"
	enumRange "github.com/bogdanovds/go-YP/basics/compositeTypes/array/enum"
	enumPointer "github.com/bogdanovds/go-YP/basics/compositeTypes/array/enum"
)

func main() {
	a := enumFor.RangeWeekTemp()
	b := enumRange.EnumRange()
	c := enumPointer.EnumPointer()
	fmt.Println("Среднее значение температуры воздуха за неделю:", a, b, c)
}
