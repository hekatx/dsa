package main

import (
	"fmt"

	s "github.com/hekatx/dsa/list/single"
)

func main() {
	list := s.New([]any{})
	list.PushFront([]any{"hello", 1, 5, 7})
	fmt.Println(list.Next)
}
