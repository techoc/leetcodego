package easy

func finalPrices(prices []int) (ans []int) {
	//如果所选择的商品价格在后面有小于该商品的价格
	//则优惠后的商品价格就等于 原商品价格 - 小于原商品的价格
	ans = make([]int, len(prices))
	for index, price := range prices {
		discount := 0
		for _, discountPrice := range prices[index+1:] {
			//获得第一个小于当前商品价格的折扣
			if discountPrice <= price {
				discount = discountPrice
				break
			}
		}
		ans[index] = price - discount
	}
	return ans
}
