package easy

func busyStudent(startTime []int, endTime []int, queryTime int) (res int) {
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			res++
		}
	}
	return res
}
