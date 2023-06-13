package tasks

import "fmt"

func ChangeBaseArr() {
	s := []int{1, 2, 3}
	Append(&s)
	fmt.Println(s)
}

func Append(s *[]int) {
	*s = append(*s, 4)
}
