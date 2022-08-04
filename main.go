package main

import (
	"fmt"

	s "github.com/hekatx/dsa/list/single"
)

func main() {
	list := s.New([]any{})
	list.PushFront([]any{"hello", 1, 5, 7})
	list.Reverse(list.Next)
	cn := list.Next
	for i := 0; i < 4; i++ {
		fmt.Println(cn)
		cn = cn.Next
	}
}
