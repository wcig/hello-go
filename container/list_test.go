package test

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()

	// 尾部添加
	l.PushBack("canon")
	printList(l)

	// 头部添加
	l.PushFront(67)
	printList(l)

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	printList(l)
	fmt.Println(element.Value)

	// 在fist之后添加high
	l.InsertAfter("high", element)
	printList(l)

	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	printList(l)

	// 使用
	l.Remove(element)
	printList(l)
}

func printList(list *list.List) {
	for item := list.Front(); item != nil; item = item.Next() {
		fmt.Print(item.Value)
		if item.Next() != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func TestAAA(t *testing.T) {
	s := new(int)
	*s = 98
	fmt.Printf("%p %v\n", s, *s) // 0xc00008a160 98

	type user struct {
		name string
		age  int
	}
	u := new(user)
	u.name = "Tom"
	u.age = 10
	fmt.Printf("%p %v\n", u, *u) // 0xc000098080 {Tom 10}

	var n *int
	n = new(int) // 分配空间 注释掉将报错：panic: runtime error: invalid memory address or nil pointer dereference
	*n = 98
	fmt.Printf("%p %v\n", n, *n) // 0xc00008a170 98
}
