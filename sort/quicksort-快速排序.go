package main

import "fmt"

/*快速排序
先把第一个数提出来,所有小于的放左边,大于的放右边
*/
func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		splitdata := arr[0]       //以第一个为基准元素
		low := make([]int, 0, 0)  //存储比我小的
		high := make([]int, 0, 0) //存储比我大的
		mid := make([]int, 0, 0)  //存储与我相等的
		mid = append(mid, splitdata)
		for i := 1; i < length; i++ { //重点!!!!!!!!!!!!循环时跳过所选的基准元素
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)   //切割递归处理
		myarr := append(append(low, mid...), high...) //...将slice切片结构打散后传递
		return myarr
	}
}

//改良版
func QuickSort2(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		n := 0                    //n必须 n>=0&& n<length
		splitdata := arr[n]       //以第一个为基准元素
		low := make([]int, 0, 0)  //存储比我小的
		high := make([]int, 0, 0) //存储比我大的
		mid := make([]int, 0, 0)  //存储与我相等的
		mid = append(mid, splitdata)
		for i := 0; i < length; i++ { //重点!!!!!!!!!!!!循环时跳过所选的基准元素
			if i == n {
				continue //到自己直接退出 不用参与比较
			}
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)   //切割递归处理
		myarr := append(append(low, mid...), high...) //...将slice切片结构打散后传递
		return myarr
	}
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(arr)
	fmt.Println(QuickSort2(arr))
}
