package main

import "fmt"

/*希尔排序
步长收缩的算法
通过步长每次收缩一不断处理进行排序
通常用在并发场合
*/

//处理步长
func shellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { //插入排序的变种
		backup := arr[i] //备份插入的数据
		j := i - gap     //上一个位置循环找到位置插入
		for j > 0 && backup < arr[j] {
			arr[j+gap] = arr[j] //从前往后移动
			j -= gap
		}
		arr[j+gap] = backup //插入
	}
}
func SheelSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组,直接返回
	} else {
		gap := len(arr) / 2
		for gap > 0 {
			for i := 0; i < gap; i++ { //处理每个元素的步长
				shellSortStep(arr, i, gap)
			}
			gap /= 2 //gap--每次循环万改变步长
		}

	}
	return arr
}
func main() {
	arr := []int{11, 2, 9, 3, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(SheelSort(arr))

}
