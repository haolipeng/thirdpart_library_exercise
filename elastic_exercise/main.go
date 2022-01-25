package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name string
	age  int
}

func main() {
	u := &user{
		"haolipeng",
		20,
	}
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "zhangsan"

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 29

	fmt.Println(u)
}
