package t1450

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0
	for i := 0; i < len(startTime); i++ {
		// 如果开始时间小于查询时间 并且 结束时间大于查询时间
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}
