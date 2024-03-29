package easy

import "testing"

func Test191(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	res := hammingWeight(num)
	t.Log(res)
}
