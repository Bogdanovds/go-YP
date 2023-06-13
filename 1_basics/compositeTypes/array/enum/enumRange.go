package enum

func EnumRange() int{
	var weekTemp = [7]int{4, 4, 6, 8, 21, 9, 5}
	
	sumTemp := 0

	for _, temp := range weekTemp {
		sumTemp += temp
	}
	average := sumTemp / len(weekTemp) 
	
	return average
}