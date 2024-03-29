package easy

import "strings"

func numUniqueEmails(emails []string) int {
	emailSet := map[string]struct{}{}
	for _, email := range emails {
		i := strings.IndexByte(email, '@')
		local := strings.SplitN(email[:i], "+", 2)[0] // 去掉本地名第一个加号之后的部分
		local = strings.ReplaceAll(local, ".", "")    // 去掉本地名中所有的句点
		emailSet[local+email[i:]] = struct{}{}
	}
	return len(emailSet)
}
