package hard

import "sort"

func zeta(n int) (res int) {
	//符合的x个数有多少个？
	//显然，x每+5，阶乘就会至少多乘一个5，末尾就会至少多一个0，所以如果上面的x有解，那就是5个，如果无解就是0
	for n > 0 {
		n /= 5
		res += n
	}
	return
}

func nx(k int) int {
	return sort.Search(5*k, func(x int) bool { return zeta(x) >= k })
}

func preimageSizeFZF(k int) int {
	return nx(k+1) - nx(k)
}
