package leetcode

//给你一个非负整数 x ，计算并返回x的 算术平方根 。
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//示例 1：
//输入：x = 4
//输出：2
//示例 2：
//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

func mySqrt(x int) int {
	res := 1
	for {
		if res*res > x {
			break
		}
		res++
	}
	return res - 1
}

//二分查找
func mySqrtQuick(x int) int {
	// 二分
	left, right := 0, x
	res := 0
	for left <= right {
		mid := (right-left)>>1 + left
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
