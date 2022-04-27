package single

import "testing"

func TestEmptyList(t *testing.T) {
	list := New()

	if list.next != nil {
		t.Fatalf("List contains elements")
	}

	if !list.Empty() {
		t.Errorf("List is not empty, contains %v", list.size)
	}
}

func TestNewList(t *testing.T) {
	list := New("First value")

	if list.Empty() {
		t.Errorf("The list is empty, it contains %v nodes", list.size)
	}

	list2 := New()

	list2.PushFront("Second list, first value")

	got := list2.next.value
	want := "Second list, first value"

	if got != want {
		t.Errorf("A different element was added: %s", got)
	}
}

func TestSequentialNodes(t *testing.T) {
	mockValues := []interface{}{1, "Hello", 231, "World"}
	list := New(mockValues...)

	currentNode := list.next

	for i := len(mockValues) - 1; i >= 0; i-- {
		currentMockValue := mockValues[i]
		if currentMockValue != currentNode.value {
			t.Errorf("Expected next node to be %v but instead got %v", currentMockValue, currentNode.value)
		}
		currentNode = currentNode.next
	}
}

func TestValueAt(t *testing.T) {
	mockValues := []interface{}{1, "Hello", 231, "World"}

	list := New(mockValues...)

	got, err := list.ValueAt(2)
	want := "Hello"

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Expected list.ValueAt(2) to equal %v but instead go %v", want, got)
	}
}

func TestPopFront(t *testing.T) {
	mockValues := []interface{}{1, "Hello", 231, "World"}

	list := New()
	_, err := list.PopFront()

	if err == nil {
		t.Error("Calling PopFront on an empty list should have thrown an error")
	}

	list2 := New(mockValues...)

	got, _ := list2.PopFront()
	want := "World"

	if got != want {
		t.Errorf("Popped the wrong element")
	}
}

func TestPushBack(t *testing.T) {
	mockValues := []interface{}{1, "Hello", 231, "World"}
	list := New(mockValues...)

	want := "Last value"
	list.PushBack(want)
	got := list.Back().value

	if got != want {
		t.Errorf("Either pushed or retrieved wrong element, expected %v but received %v", want, got)
	}
}
