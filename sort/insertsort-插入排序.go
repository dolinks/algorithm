package main

import "fmt"

/*插入排序*/

func InsertTest(arr []int) []int {
	backup := arr[2]
	j := 2 - 1 // 上一个位置循环找到位置插入
	for j > 0 && backup < arr[j] {
		arr[j+1] = arr[j] //从前往后移动
		j--
	}
	arr[j+1] = backup //插入
	return arr
}

func InserSort(arr []int) []int {
	length := len(arr)
	if length <= 1 { //一个元素不处理直接返回原数组
		return arr
	} else {
		for i := 1; i < length; i++ {
			backup := arr[i] //备份插入的数据
			j := i - 1       //上一个位置循环找到位置插入
			for j > 0 && backup < arr[j] {
				arr[j+1] = arr[j] //从前向后移动
				j--
			}
			arr[j+1] = backup //插入
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int{39, 1, 19, 29, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(InserSort(arr))
}
