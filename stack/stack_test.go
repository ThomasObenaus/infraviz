package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type myKVP struct {
	key   string
	value int
}

func TestNewStack(t *testing.T) {

	stack := NewStack()
	require.NotNil(t, stack)

}

func TestPeek(t *testing.T) {

	stack := NewStack()
	require.NotNil(t, stack)

	stack.Push("Hallo")
	assert.Equal(t, 1, stack.Size())
	assert.Equal(t, "Hallo", stack.Peek())
}

func TestPush(t *testing.T) {

	stack := NewStack()
	require.NotNil(t, stack)

	stack.Push("Hallo")
	assert.Equal(t, 1, stack.Size())
}

func TestPop(t *testing.T) {

	stack := NewStack()
	require.NotNil(t, stack)

	elem := stack.Pop()
	assert.Nil(t, elem)

	stack.Push("Hallo")
	elem = stack.Pop()
	assert.Equal(t, "Hallo", elem)
}

func TestMultiWithStruct(t *testing.T) {

	stack := NewStack()
	require.NotNil(t, stack)

	kvp := myKVP{key: "height", value: 10}
	stack.Push(kvp)

	kvp = myKVP{key: "height", value: 20}
	stack.Push(kvp)

	kvp = myKVP{key: "height", value: 30}
	stack.Push(kvp)

	elem := stack.Pop().(myKVP)
	assert.Equal(t, "height", elem.key)
	assert.Equal(t, 30, elem.value)

	elem = stack.Pop().(myKVP)
	assert.Equal(t, "height", elem.key)
	assert.Equal(t, 20, elem.value)

	elem = stack.Pop().(myKVP)
	assert.Equal(t, "height", elem.key)
	assert.Equal(t, 10, elem.value)

	elemNew := stack.Pop()
	assert.Nil(t, elemNew)
}

func Example() {
	stack := NewStack()

	stack.Push("world")
	stack.Push("golang")
	stack.Push("hello")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	// Output: hello
	// golang
	// world
}
