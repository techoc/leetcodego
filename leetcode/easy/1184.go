package easy

func distanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	sum1, sum2 := 0, 0
	for i, d := range distance {
		if start <= i && i < destination {
			sum1 += d
		} else {
			sum2 += d
		}
	}
	return min(sum1, sum2)
}
