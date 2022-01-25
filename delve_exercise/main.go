package main

import (
	"fmt"
	"github.com/prashantv/gostub"
)

var counter = 100

func stubGlobalVariable() {
	stubs := gostub.Stub(&counter, 200)
	defer stubs.Reset()
	fmt.Println("Counter:", counter)
}

func main() {
	stubGlobalVariable()
}
