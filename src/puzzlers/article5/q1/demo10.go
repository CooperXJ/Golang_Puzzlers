package main

import "fmt"

var block = "package"

//可重名变量之间不存在类似的限制，它们的类型可以是任意的
func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
