package enum

func RangeWeekTemp() int {
	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

	sumTemp := 0

	for i := 0; i < len(weekTemp); i++ {
		sumTemp += weekTemp[i]
	}

	average := sumTemp / len(weekTemp)

	return average
}