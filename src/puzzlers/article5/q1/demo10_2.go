package main

//当导入一个包的前面是"."时，那么此时使用该包中的公开函数时不需要加上前缀，比如此处的fmt在Printf的时候就不需要写成fmt.Printf了
import . "fmt"

func test() {
	block := "function"
	{
		block := "inner"
		Printf("The block is %s.\n", block)
	}
	Printf("The block is %s.\n", block)
}
