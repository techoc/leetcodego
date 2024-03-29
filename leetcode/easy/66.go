package easy

//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//示例1：
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
//示例2：
//输入：digits = [4,3,2,1]
//输出：[4,3,2,2]

func plusOne(digits []int) []int {
	//将最后一位加1
	digits[len(digits)-1]++
	//判断是否需要进位
	if digits[len(digits)-1] >= 10 {
		digits[len(digits)-1] = 0
		//判断是否需要进位
		for i := len(digits) - 2; i >= 0; i-- {
			digits[i]++
			if digits[i] < 10 {
				break
			}
			digits[i] = 0
		}
		//判断是否需要添加一位
		if digits[0] == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
