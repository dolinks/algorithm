package main

/*字符串选择排序*/
import "fmt"

func SelectSortMaxString(arr []string) string {
	max := "sd"
	return max
}

func main() {
	fmt.Println("a" > "b")
	fmt.Println("a" < "b")
	fmt.Println("a" == "a")
	pb := "b"
	pa := "a"
	fmt.Println(&pa, &pb) //go 1.1,1.比较地址,go1.10优化,都可以比较大小
	fmt.Println((pa < pb))
	arr := []string{"c", "a", "b", "x", "z", "m", "n", "d", "f"}
	fmt.Println(arr)

}
