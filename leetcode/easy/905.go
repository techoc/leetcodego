package easy

func sortArrayByParity(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			res[left] = nums[i]
			left++
		} else {
			res[right] = nums[i]
			right--
		}
	}
	return res
}
