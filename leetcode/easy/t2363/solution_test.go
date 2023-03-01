package t2363

import "testing"

func TestMergeSimilarItems(t *testing.T) {
	items1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	items2 := [][]int{{3, 1}, {1, 5}}

	items := mergeSimilarItems(items1, items2)

	t.Logf("%v", items)

}
