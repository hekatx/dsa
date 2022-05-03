# Data structrures and algorithms in Go

My first implementation of common DSA for learning purposes.

## Linked lists

```go
// Current implementation only admits slices with values
// so to create an empty linked list you have to call New() with an empty slice. Slices can be any type.
l := list.New([]any{})

v := []string{"hello", "world"}

// Since PushFront is used to insert nodes, list order will be the reverse of the slice used to create it

// You can add elements to the start of the list with:
l.PushFront(v)

v2 := []int{1}

// Similarly, you can add elements to the end of the list with:
l.PushBack(v2)

// Get a particular value by index
l.ValueAt(2) // returns "hello"

// Remove the first element of the list
v, err = l.PopFront() // returns the "world" and an error

// Get first element of the list
l.Front() // returns "hello" since we popped "world"

// Get last element of the list
l.Back() // returns 1
```
