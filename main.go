package main

import (
	"fmt"
	"github.com/techoc/leetcodego/leetcode"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{-2}
	fmt.Println(leetcode.MaxSubArray1(nums))
}
