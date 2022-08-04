package double

import (
	"testing"

	// "github.com/hekatx/dsa/utils"
	"github.com/stretchr/testify/assert"
)

func TestEmptyList(t *testing.T) {
	list := New()

	assert.Nil(t, list.Head, "List contains elements")
	assert.Nil(t, list.Tail, "List contains elements")

	assert.True(t, list.Empty(), "List is not empty, contains not indexed by Head nor List")
}

func TestNewList(t *testing.T) {
	v := "head and tail"
	list := New(v, 2, 3)

	assert.Equal(t, list.Head.value, v, "Head is pointing to the wrong value")
	assert.Equal(t, list.Tail.value, 3, "Tail is pointing to the wrong value")
}

func TestPush(t *testing.T) {
	list := New(1, 2, 3)
	list.PushFront(4)

	assert.Equal(t, list.Head.value, 1, "Head is pointing to the wrong value")
	assert.Equal(t, list.Tail.value, 4, "Tail is pointing to the wrong value")

	list2 := New()
	list2.PushFront(4)

	assert.Equal(t, list2.Head.value, 4, "Head is pointing to the wrong value")
	assert.Equal(t, list2.Tail.value, 4, "Tail is pointing to the wrong value")
}

func TestEmpty(t *testing.T) {
	list := New()

	assert.True(t, list.Empty(), "List is not empty")

	list2 := New(1)

	assert.False(t, list2.Empty(), "List is empty")
}

func TestSize(t *testing.T) {
	v := []any{1, 2, 3, 4}
	list := New(v...)

	assert.Equal(t, list.size, len(v), "List doesn't have the appropiate size")
}

func TestValueAt(t *testing.T) {
	list := New(1, 2, 3, "g")

	v, e := list.ValueAt(1)

	assert.Nil(t, e, "An error ocurred")
	assert.Equal(t, 1, v.value)
}
