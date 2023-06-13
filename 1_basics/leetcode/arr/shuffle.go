// На вход функции поступает массив с количеством 2n элементов вида: [X1, X2, ..., Xn, Y1, Y2,...Yn].
// Требуется преобразовать этот массив к виду: [X1, Y1, X2, Y2, ..., Xn, Yn].

package arr

import "fmt"

func ShuffleArray() {
	nums, n := [...]int{1, 2, 3, 4, 4, 3, 2, 1}, 4
	var ans [len(nums)]int

	x, y := 0, 1

	for i := 0; i < n; i++ {
		ans[x], ans[y] = nums[i], nums[i+n]
		x += 2
		y += 2
	}

	fmt.Println(ans)
}

func ShuffleSlice() {
	nums, n := [...]int{1, 2, 3, 4, 4, 3, 2, 1}, 4
	ans := make([]int, 0)

	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}

	fmt.Println(ans)
}
