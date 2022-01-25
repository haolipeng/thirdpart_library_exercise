package main

import (
	"fmt"
	"github.com/fatih/set"
)

func main() {
	oldPortSet := set.New(set.ThreadSafe)
	newPortSet := set.New(set.ThreadSafe)

	oldPortSet.Add(0)
	oldPortSet.Add(1)
	oldPortSet.Add(2)
	oldPortSet.Add(3)
	oldPortSet.Add(10)
	oldPortSet.Add(11)

	newPortSet.Add(1)
	newPortSet.Add(2)
	newPortSet.Add(3)
	newPortSet.Add(4)

	dSet := set.Difference(oldPortSet, newPortSet)
	dSet.Each(func(i interface{}) bool {
		if id, ok := i.(int); ok {
			fmt.Println("value:", id)
		}

		return true
	})
}
