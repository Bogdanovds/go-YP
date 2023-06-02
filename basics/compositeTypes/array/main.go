package main

import (
	enumValuesFor "github.com/bogdanovds/go-YP/basics/compositeTypes/array/enum"
	"fmt"
)

func main() {
	a := enumValuesFor.RangeWeekTemp()
	fmt.Println("Среднее значение температуры воздуха за неделю:", a)
}
