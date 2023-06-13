package t2475

func unequalTriplets(nums []int) int {
	res, n := 0, len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					res++
				}
			}
		}
	}
	return res
}
