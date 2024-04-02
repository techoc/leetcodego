package t1108

import "strings"

// 1108. IP 地址无效化
// https://leetcode.cn/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	var res string
	for _, s := range address {
		if s == '.' {
			res += "[.]"
		} else {
			res += string(s)
		}
	}
	return res
}

func defangIPaddr2(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
