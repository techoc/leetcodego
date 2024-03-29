package t2325

import (
	"strings"
)

func decodeMessage(key string, message string) string {
	//得到密码表
	keyTable := make(map[string]string)
	//去除空格
	key = strings.ReplaceAll(key, " ", "")
	count := 0
	for i := 0; i < len(key); i++ {
		//如果是第一次进入密码表
		_, prs := keyTable[string(key[i])]
		if !prs {
			keyTable[string(key[i])] = string(rune(97 + count))
			count++
		}
	}

	//根据密码表解密字符
	var res string
	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			res += " "
		} else {
			res += keyTable[string(message[i])]
		}
	}
	return res
}

func decodeMessageOffice(key string, message string) string {
	cur := byte('a')
	rules := map[rune]byte{}

	for _, c := range key {
		if c != ' ' && rules[c] == 0 {
			rules[c] = cur
			cur++
		}
	}

	m := []byte(message)
	for i, c := range message {
		if c != ' ' {
			m[i] = rules[c]
		}
	}

	return string(m)
}
