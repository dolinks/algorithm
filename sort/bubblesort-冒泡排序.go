package main

import "fmt"

/*
冒泡排序
*/
//遍历找出最大值,最大值沉底
func BubbleFindMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] { //两两比较,大的放后边
				arr[i], arr[i+1] = arr[i+1], arr[i] //交换变量
			}

		}
		return arr[length-1] //最后一个元素即为最大值
	}
}

//冒泡排序
func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //*每次排序沉底最大一个,只剩一个,不需要冒泡了
			isneedExchange := false           //不需要交换
			for j := 0; j < length-1-i; j++ { //-i 算法精华 当i等于0的时候 只需要虚幻到倒数一位,当i等于1时,只需要循环到倒数二位
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isneedExchange = true
				}
				//fmt.Println(arr)
			}
			if !isneedExchange {
				break
			}
		}
		return arr
	}
}

//冒泡排序
//func BubbleSort(arr []int) []int {
//	length := len(arr)
//	if length <= 1 {
//		return arr
//	} else {
//		for i := 0; i < length; i++ {
//			for j := length; j < i; j-- {
//				if arr[j] > arr[i] {
//					arr[i] = arr[j]
//				}
//			}
//		}
//		return arr
//	}
//}

func main() {
	arr := []int{11, 9, 2, 1, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(BubbleFindMax(arr))
	fmt.Println(BubbleSort(arr))
}
