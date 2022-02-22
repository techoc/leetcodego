package leetcode

//给你一个整数数组nums。如果nums的一个子集中，所有元素的乘积可以表示为一个或多个 互不相同的质数 的乘积，那么我们称它为好子集。
//
//比方说，如果nums = [1, 2, 3, 4]：
//[2, 3]，[1, 2, 3]和[1, 3]是 好子集，乘积分别为6 = 2*3，6 = 2*3和3 = 3。
//[1, 4] 和[4]不是 好子集，因为乘积分别为4 = 2*2 和4 = 2*2。
//请你返回 nums中不同的好子集的数目对109 + 7取余的结果。
//nums中的 子集是通过删除 nums中一些（可能一个都不删除，也可能全部都删除）元素后剩余元素组成的数组。如果两个子集删除的下标不同，那么它们被视为不同的子集。
//示例 1：
//输入：nums = [1,2,3,4]
//输出：6
//解释：好子集为：
//- [1,2]：乘积为 2 ，可以表示为质数 2 的乘积。
//- [1,2,3]：乘积为 6 ，可以表示为互不相同的质数 2 和 3 的乘积。
//- [1,3]：乘积为 3 ，可以表示为质数 3 的乘积。
//- [2]：乘积为 2 ，可以表示为质数 2 的乘积。
//- [2,3]：乘积为 6 ，可以表示为互不相同的质数 2 和 3 的乘积。
//- [3]：乘积为 3 ，可以表示为质数 3 的乘积。

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func numberOfGoodSubsets(nums []int) (ans int) {
	const mod int = 1e9 + 7
	freq := [31]int{}
	//找到数组中1-30出现的次数
	for _, num := range nums {
		freq[num]++
	}

	f := make([]int, 1<<len(primes))
	f[0] = 1
	for i := 0; i < freq[1]; i++ {
		f[0] = f[0] * 2 % mod
	}
next:
	for i := 2; i < 31; i++ {
		//跳过数组中不存在的数字
		if freq[i] == 0 {
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个
		subset := 0
		for j, prime := range primes {
			//如果i对质因子的平方取余等于0，则表示有相同的质因子相乘，直接跳过
			if i%(prime*prime) == 0 {
				continue next
			}
			//如果i对质因子的取余等于0，则将质因子所处位置的bitmask改为1
			if i%prime == 0 {
				subset |= 1 << j
			}
		}

		// 动态规划
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subset == subset {
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return
}
