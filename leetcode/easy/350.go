package easy

// 给定两个数组，编写一个函数来计算它们的交集。
// 示例 1：
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]

//哈希表实现
func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return Intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}
