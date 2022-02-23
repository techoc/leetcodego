package leetcode

//给你一个字符串 s ，根据下述规则反转字符串：
//所有非英文字母保留在原有位置。
//所有英文字母（小写或大写）位置反转。
//返回反转后的 s 。
//示例 1：
//输入：s = "ab-cd"
//输出："dc-ba"
//示例 2：
//输入：s = "a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"

func reverseOnlyLetters(s string) string {
	//初始化头尾指针
	head, tail := 0, len(s)-1
	var isLetter = func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
	}
	//Go语言的字符串无法直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现
	//将字符串转换为字节数组
	res := []byte(s)
	//遍历字符串
	for head < tail {
		//如果是字母，则交换
		if isLetter(res[head]) && isLetter(res[tail]) {
			res[head], res[tail] = res[tail], res[head]
			head++
			tail--
		} else if !isLetter(s[head]) {
			head++
		} else {
			tail--
		}
	}
	return string(res)
}
