package leetcode

//	最大子序和
//	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//	输出：6
//	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func MaxSubArray(nums []int) int {
	//贪心
	max := -2147483648
	sum := 0
	for _, value := range nums {
		sum = value + sum
		if sum >= max {
			max = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return max
}

func MaxSubArrayDynamic(nums []int) int {
	//动态规划
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func MaxSubArrayDivideConquer(nums []int) int {
	//分治
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
