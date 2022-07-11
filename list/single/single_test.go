package single

import (
	"testing"

	"github.com/hekatx/dsa/utils"
	"github.com/stretchr/testify/assert"
)

func getMockValues() []any {
	v := []any{}

	vlen := int(utils.RandomInt(4, 10))

	for i := 0; i < vlen; i++ {
		s := utils.RandomString()
		v = append(v, s)
	}

	return v
}

func TestEmptyList(t *testing.T) {
	list := New([]any{})

	assert.Nil(t, list.Next, "List contains elements")

	assert.True(t, list.Empty(), "List is not empty, contains something")

}

func TestNewList(t *testing.T) {
	list := New([]string{"Second list, first value"})

	assert.Equal(t, list.Next.value, "Second list, first value", "A different element was added: %s")
}

func TestSequentialNodes(t *testing.T) {
	m := getMockValues()
	list := New(m)

	currentNode := list.Next

	for i := len(m) - 1; i >= 0; i-- {
		currentMockValue := m[i]
		assert.Equal(t, currentMockValue, currentNode.value, "Expected next node to be %v but instead got %v")
		currentNode = currentNode.Next
	}
}

func TestValueAt(t *testing.T) {
	m := getMockValues()
	list := New(m)

	v, err := list.ValueAt(2)

	assert.Nil(t, err)
	assert.Equal(t, v, m[len(m)-3], "Expected %v to equal %v in %v", v, m[len(m)-3], m)
}

func TestPopFront(t *testing.T) {
	m := getMockValues()

	list := New([]any{})
	_, err := list.PopFront()

	assert.NotNil(t, err, "Calling PopFront on an empty list should have thrown an error")

	list2 := New(m)

	v, _ := list2.PopFront()

	assert.Equal(t, v, m[len(m)-1], "Popped the wrong element")
}

func TestPushBack(t *testing.T) {
	m := getMockValues()
	list := New(m)

	want := m[len(m)-1]
	list.PushBack(want)
	got := list.Back().value

	assert.Equal(t, got, want, "Either pushed or retrieved wrong element, expected %v but received %v", want, got)
}

func TestReverse(t *testing.T) {
	m := getMockValues()
	list := New(m)

	list.Reverse(list.Next)

	currentNode := list.Next

	for i := 0; i <= 0; i++ {
		currentMockValue := m[i]
		assert.Equal(t, currentMockValue, currentNode.value, "Expected next node to be %v but instead got %v")
		currentNode = currentNode.Next
	}
}
