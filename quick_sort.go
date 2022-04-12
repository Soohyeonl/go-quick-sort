package quick_sort

import (
	"fmt"
	"math/rand"
)

// Partition 划分函数，将数组分为两部分，左边小于指定数字，右边大于指定数字
func Partition(data *[]int, start int, end int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 特殊情况判断
	if data == nil || len(*data) <= 0 || start < 0 || end >= len(*data) {
		panic("data error")
	}

	// 随机选择一个位置的数字为标志数字
	index := rand.Intn(end-start+1) + start
	(*data)[index], (*data)[end] = (*data)[end], (*data)[index] // 将标志数字移到最后

	small := start - 1 // small指向最右的小于标志数字的位置，因为不知道第一个数字跟标志数字的大小关系，所以初始时small = start - 1

	// 循环遍历end之前的数字
	for index = start; index < end; index++ {
		// index指向的数字小于标志数字
		if (*data)[index] < (*data)[end] {
			// small++，因为现在小于标志数字的个数增加了一个
			small++
			// 只有当左侧存在大于标志数字的数，才会交换数字的位置
			if small != index {
				(*data)[index], (*data)[small] = (*data)[small], (*data)[index]
			}
		}
	}

	// 最后small指向第一个大于标志数的位置
	small++
	// 将标志数字从end换到small的位置，所以small之前的数字都小于标志数，small之后的数字都大于标志数
	(*data)[end], (*data)[small] = (*data)[small], (*data)[end]

	// 返回标志数的位置
	return small
}

// QuickSort 接下来的就是递归的去排序
func QuickSort(data *[]int, start int, end int) {
	if start == end {
		return
	}

	index := Partition(data, start, end)
	if index > start {
		QuickSort(data, start, index-1)
	}
	if index < end {
		QuickSort(data, index+1, end)
	}
}
