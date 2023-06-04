// Функция sort.Ints сортирует полученный слайс целых чисел по возрастанию.
// Она не меняет размер и ёмкость слайса, поэтому может спокойно работать с ним.
package assignment

import (
	"fmt"
	"sort"
)

func Assignment() {
	s := []int{5, 4, 1, 3, 2}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4 5]
}
