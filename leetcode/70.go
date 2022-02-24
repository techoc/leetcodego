package leetcode

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var a, b, c int
	//a 表示第n-2个台阶的走法，b表示第n-1个台阶的走法,传统迭代方法
	a, b = 1, 2
	for i := 3; i <= n; i++ {
		//累加结果
		c = a + b
		//向下迭代
		a = b //下次迭代的第n-2个台阶的走法等于上次迭代n-1个台阶的走法
		b = c //下次迭代的第n-1个台阶的走法等于上次迭代的第n个台阶走法
	}
	return c
}
