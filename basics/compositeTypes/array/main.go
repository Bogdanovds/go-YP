package main

import (
	rangevalues "example/hello/go/src/github.com/bogdanovds/go-YP/basics/compositeTypes/array/rangeValues"
	"fmt"
)

func main() {
	a := rangevalues.RangeWeekTemp()
	fmt.Println("Среднее значение температуры воздуха за неделю:", a)
}
