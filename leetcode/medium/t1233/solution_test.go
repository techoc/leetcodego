package t1233

import "testing"

func TestRemoveSubfolders(t *testing.T) {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	ans := removeSubfolders(folder)
	t.Log(ans)
}
