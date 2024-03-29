package t1233

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/
func removeSubfolders(folder []string) (ans []string) {

	// 字典序排序
	sort.Strings(folder)
	ans = append(ans, folder[0])

	// 从第二个开始遍历
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		// 如果当前目录和上一个结果有不同的前缀 或者 f 多出来的第一个字符不是 '/' 则不删除该目录
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return
}
