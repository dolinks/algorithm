package main

/*奇偶排序
奇数位置坐标排一次,偶数位置坐标排一次,然后反复...知道完成

1,9,3,8,2,7,5,6,4,0
1	3	2	5	4
7	8	9	6	0
*/
import "fmt"

func OddEvenSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		isSorted := false //技术位,偶数位都不需要排序的有序
		for isSorted == false {
			isSorted = true
			for i := 1; i < length-1; i = i + 2 { //奇数位
				if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
					isSorted = false
				}
			}
			fmt.Println("奇数", arr)
			for i := 0; i < length-1; i = i + 2 { //偶数位
				if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
					isSorted = false
				}
			}
			fmt.Println("偶数", arr)
		}
	}
	return arr
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	OddEvenSort(arr)
	fmt.Println("OddEvenOrder")
}
