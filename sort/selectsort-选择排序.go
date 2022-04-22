package main

import (
	"fmt"
	"strings"
)

/*
选择排序
*/

//找到最大值
func SelectSortMax(arr []int) int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 0; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

//数字选择排序
func SelectSort(arr []int) []int {
	length := len(arr) // 数组长度
	if length <= 1 {
		return arr //一个数组元素直接返回
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要挑选
			min := i                          //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				if arr[min] < arr[j] { //如果改成大则从小到大排序
					min = j //保存极小值索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] // 数据交换
			}
		}
	}
	return arr
}

//字符串选择排序
func main() {
	arr := []int{1, 9, 2, 8, 3, 10, 7, 4, 6, 5, 0}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(SelectSort(arr))
	//相等2,不等-1,+1,
	//a<b<c  首先比较第一个字母,左边小于右边-1,左边大于右边+1,第一个字母比较不成功比较第二个
	//fmt.Println("a" < "b")  //字符串存在的地址比较
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.Compare("a2", "a1"))
	fmt.Println(strings.Compare("a2", "a2"))
}
