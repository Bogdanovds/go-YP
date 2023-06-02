package mathslice

type Slice []Element

type Element int

//SumSlice - возвращает сумму элементов
func SumSlice(slice Slice) (res Element) {
	for _, s := range slice {
		res += s
	}
	return

}

//MapSlice - применяет функцию ор к каждому элементу
func MapSlice(slice Slice, op func(Element) Element) (res Element) {
	for i, s := range slice {
		slice[i] = op(s)
	}
	return
}

// FolbSlice - сворачивает слайс
func FolbSlice(slice Slice, op func(Element, Element) Element, init Element) (res Element) {
	res = op(init, slice[0])
	for i := 0; i < len(slice); i++ {
		res = op(res, slice[i])
	}
	return
}
