package t2460

import "testing"

func TestApplyOperations(t *testing.T) {
	nums := []int{1, 2, 2, 1, 1, 0}
	operations := applyOperations(nums)
	operationsOffice := applyOperationsOffice(nums)
	t.Log(operations)
	t.Log(operationsOffice)
}
