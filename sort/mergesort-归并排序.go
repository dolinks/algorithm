package main

import "fmt"

/*归并排序
分段排序,最后归并
*/

func Merge(leftarr []int, rightarr []int) []int {
	leftindex := 0  //左边索引
	rightindex := 0 //右边的索引
	lastarr := []int{}
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else {
			lastarr = append(lastarr, leftarr[leftindex])
			lastarr = append(lastarr, rightarr[rightindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(leftarr) { //把没有结束的归并过来
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr) {
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr //小于10改用插入排序
	} else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])
		return Merge(leftarr, rightarr)
	}
}
func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(MergeSort(arr))
}
