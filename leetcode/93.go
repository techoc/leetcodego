package leetcode

import (
	"strconv"
	"strings"
)

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
// 你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]

func RestoreIpAddresses(s string) []string {
	res := []string{}
	var dfs func(subRes []string, start int)

	dfs = func(subRes []string, start int) {
		//如果字符串长度为4，切点开始的位置等于字符串的长度，则说明已经切分完毕
		if len(subRes) == 4 && start == len(s) {
			res = append(res, strings.Join(subRes, "."))
			return
		}
		//如果字符出啊你的长度为4，切点开始的位置小于字符串的长度，则说明还没有切分完毕
		if len(subRes) == 4 && start < len(s) {
			return
		}
		for length := 1; length <= 3; length++ {
			//如果切点开始的位置加上当前的长度大于字符串的长度，则说明当前的长度不合法，直接跳过
			if start+length-1 >= len(s) {
				return
			}
			//如果当前的长度不为0并且切点位置上字符的位置为0，则说明当前的长度不合法，直接跳过
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			//
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			subRes = append(subRes, str)
			dfs(subRes, start+length)
			subRes = subRes[:len(subRes)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}
