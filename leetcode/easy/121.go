package easy

// 买卖股票的最佳时机
// 思路: 有两种操作 买/卖 最优化买:对比 上一次买与这次买 选择最省钱的方式
// 最优化卖:对比 上次卖与这次卖 选择最赚钱的方式 4ms/3.6mb
func MaxProfit(prices []int) int {
	// 0 = 持有股票
	// 1 = 不持有股票
	n := len(prices)
	if n < 2 {
		return 0
	}

	f := make([][2]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		// 这个时候购入
		f[i][0] = Max(f[i-1][0], -prices[i])
		// 这个时候卖出
		f[i][1] = Max(f[i-1][0]+prices[i], f[i-1][1])
	}
	return f[n-1][1]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	//fmt.Println(x, y)
	return y
}

// 优化: 内存可以压缩 不用建立二维数组 重复使用买卖变量 4ms/3.1mb
func MaxProfitPro(prices []int) int {
	// 0 = 持有股票
	// 1 = 不持有股票
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy := -prices[0]
	sell := 0

	for i := 1; i < n; i++ {
		// 这个时候购入
		tBuy := Max(buy, -prices[i])
		// 这个时候卖出
		tSell := Max(buy+prices[i], sell)
		buy, sell = tBuy, tSell
	}
	return sell
}

//一次遍历
func MaxProfitOne(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	//最低价格
	minprice := int(^uint(0) >> 1) //定义一个最大整数
	//最大利润
	maxprofit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
