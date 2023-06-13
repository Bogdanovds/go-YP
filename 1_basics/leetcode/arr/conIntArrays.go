package arr

import "fmt"

func ConIntArray() {
	nums := [...]int{1, 3, 2, 1}
	ans := make([]int, 2 * len(nums))
	temp := 0

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		temp = i
	}

	for j := 0; j < len(nums); j++ {
		temp += 1
		ans[temp] = nums[j]
	}

	fmt.Println(ans)
}

func ConIntSlice() {
	nums := [...]int{1, 3, 2, 1}
	var ans = nums[:]

	for i, _ := range nums {
		ans = append(ans, nums[i])
	}

	fmt.Println(ans)
}
