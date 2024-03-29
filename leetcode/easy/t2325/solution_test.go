package t2325

import "testing"

func TestDecodeMessage(t *testing.T) {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"

	str := decodeMessage(key, message)
	t.Log(str)
}

func TestDecodeMessageOffice(t *testing.T) {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"

	str := decodeMessageOffice(key, message)
	t.Log(str)
}
