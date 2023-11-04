package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("请输入第一个数")
	fmt.Scan(&a)
	fmt.Print("请输入第二个数")
	fmt.Scan(&b)
	var sum int
	sum = a + b

	fmt.Print(sum)

}
