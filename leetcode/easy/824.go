package easy

import "strings"

func toGoatLatin(sentence string) string {
	var res string
	// 分割单词
	for index, word := range strings.Split(sentence, " ") {
		//元音字母则在单词后添加 "ma"
		if word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u' || word[0] == 'A' || word[0] == 'E' || word[0] == 'I' || word[0] == 'O' || word[0] == 'U' {
			res += word + "ma"
		} else {
			//非元音字母则移除第一个字符并将它放到末尾，之后再添加 "ma"
			res += word[1:] + word[0:1] + "ma"
		}
		//每个单词都添加相同索引数的 "a"
		for i := 0; i < index+1; i++ {
			res += "a"
		}
		//如果是最后一个单词，不加空格
		if index != len(strings.Split(sentence, " "))-1 {
			res += " "
		}
	}
	return res
}
