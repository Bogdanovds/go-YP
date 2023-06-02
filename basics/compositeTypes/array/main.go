package main

import (
	rangevalues "github.com/bogdanovds/go-YP/basics/compositeTypes/"
	"fmt"
)

func main() {
	a := rangevalues.RangeWeekTemp()
	fmt.Println("Среднее значение температуры воздуха за неделю:", a)
}
