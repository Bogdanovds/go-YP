package basicoperation

import (
	"fmt"
	"reflect"
)

func BasicOperation() {
	// Удаление последнего элемента слайса
	s1 := []int{1, 2, 3}
	if len(s1) != 0 { // защищаемся от паники
		s1 = s1[:len(s1)-1]
	}
	fmt.Println(s1) // [1 2]

	// Удаление первого элемента слайса
	s2 := []int{1, 2, 3}
	if len(s2) != 0 { // защищаемся от паники
		s2 = s2[1:]
	}
	fmt.Println(s2) // [2 3]

	// Удаление элемента с индексом i
	s3 := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s3) != 0 && i < len(s3) { // защищаемся от паники
		s3 = append(s3[:i], s3[i+1:]...)
	}
	fmt.Println(s3) // [1 2 4 5]

	// Сравнение двух слайсов
	s4 := []int{1, 2, 3}
	s5 := []int{1, 2, 4}
	s6 := []string{"1", "2", "3"}
	s7 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s4, s5)) // false
	fmt.Println(reflect.DeepEqual(s4, s6)) // false
	fmt.Println(reflect.DeepEqual(s4, s7)) // true
}
