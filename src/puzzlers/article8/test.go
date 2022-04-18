package main

import (
	"container/list"
	"fmt"
)

func main() {
	//我们自己不能生成Element元素，只能让list自己生成
	list := list.New()
	a := list.PushBack(1)
	list.PushBack(2)
	list.InsertAfter("123", a)
	fmt.Printf("len: %v\n", list.Len())
	fmt.Printf("first: %#v\n", list.Front())
	fmt.Printf("second: %#v\n", list.Front().Next())
}
