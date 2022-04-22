package main

import "fmt"

/*
堆排序特点,能最快找到最大数,
数组结构,逻辑结构当作二叉树来处理
二叉树公式:
a[n]=第一个子节点a[2*n+1] 第二个子节点a[2*n+2]
深度:length/2-1
*/

//找到堆最大值
func HeapSortMax(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1        //二叉树深度
		for i := depth; i > 0; i-- { //循环所有三节点
			topmax := i                                                //假定最大的在i的位置
			leftchild := 2*i + 1                                       //左孩子结点
			rigthchild := 2*i + 2                                      //右孩子结点
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //leftchild<=length-1 防止数组越界
				topmax = leftchild
			}
			if rigthchild <= length-1 && arr[rigthchild] > arr[topmax] {
				topmax = rigthchild
			}
			if topmax != i { // 确保i的值是最大
				arr[i], arr[topmax] = arr[topmax], arr[i]
			}
		}
		return arr
	}
}

func main() {

	var arr = []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(HeapSortMax(arr))
}
