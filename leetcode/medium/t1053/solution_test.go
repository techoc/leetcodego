package t1053

import (
	"log"
	"testing"
)

func TestPrevPermOpt1(t *testing.T) {
	arr := []int{3, 2, 1}
	opt1 := prevPermOpt1(arr)
	for _, i2 := range opt1 {
		log.Println(i2)
	}
}
