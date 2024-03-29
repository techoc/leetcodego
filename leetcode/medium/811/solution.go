package _811

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domains := make(map[string]int)
	for _, str := range cpdomains {
		//分割 str 访问次数和域名
		split := strings.Split(str, " ")
		count, _ := strconv.Atoi(split[0])
		domain := split[1]

		//将整域名填入HashMap
		domains[domain] += count

		//分割子域名 将子域名的访问次数填入HashMap
		for {
			i := strings.IndexByte(domain, '.')
			if i < 0 {
				break
			}
			domain = domain[i+1:]
			domains[domain] += count
		}
	}

	//构造输出结果
	ans := make([]string, 0, len(domains))
	for str, count := range domains {
		ans = append(ans, strconv.Itoa(count)+" "+str)
	}
	return ans
}
