package arr

import "fmt"

func GoodPair() {
	nums := [...]int{1, 2, 3}
	pair := 0

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pair += 1
			}
		}
	}

	fmt.Println(pair)
}
