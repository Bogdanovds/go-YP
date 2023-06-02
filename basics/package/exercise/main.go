package main

import (
	"fmt"

	"gitlab.com/bogdanovds/YP/package/exercise/mathslice"
)

func main()  {
	s := mathslice.Slice{1,2,3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	mathslice.MapSlice(s, func(i mathslice.Element) mathslice.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))
	fmt.Println("Свёртка слайса умножением: ",
		mathslice.FolbSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x * y
			},
			1))
}