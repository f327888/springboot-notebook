package main

import (
	"container/list"
	"fmt"
)

func init2() {
	list1 := list.List{}

	fmt.Println(list1.Len())
	back := list1.Back()
	list1.InsertAfter(1, back) // nul
}

func main() {
	init2()
	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)
	lst.PushFront(200) // 放到最前面
	lst.PushFront(201) // 放到最前面
	// fmt.Println("Here is the double linked list:\n", lst)
	for e := lst.Front(); e != nil; e = e.Next() {
		// fmt.Println(e)
		fmt.Println(e.Value)
	}

}
