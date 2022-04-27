package main

import (
	"fmt"

	"github.com/hekatx/dsa/list/single"
)

func main() {
	list := single.New()
	list.PushFront("hello", 1, 5, 7)
	fmt.Println(list)
}
