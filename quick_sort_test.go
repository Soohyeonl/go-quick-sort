package quick_sort

import (
	"fmt"
	"testing"
)

// 跑一个测试用例
func TestQuickSort(t *testing.T) {
	arr := &[]int{1}
	QuickSort(arr, 0, len(*arr)-1)
	fmt.Println(*arr)
}
