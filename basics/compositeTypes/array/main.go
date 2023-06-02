package main

import (
	enumFor "github.com/bogdanovds/go-YP/basics/compositeTypes/array/enum"
	"fmt"
)

func main() {
	a := enumFor.RangeWeekTemp()
	fmt.Println("Среднее значение температуры воздуха за неделю:", a)
}
