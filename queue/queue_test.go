package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/imrenagi/pp-challenge/queue"
)

func TestNewQueue(t *testing.T) {
	q := New()
	assert.NotNil(t, q)
}

func TestIsEmpty_AfterInit(t *testing.T) {
	q := New()
	assert.True(t, q.IsEmpty())
}

func TestSize_ZeroAfterInit(t *testing.T) {
	q := New()
	assert.Equal(t, q.Size(), 0)
}

func TestPush_AfterInit(t *testing.T) {
	q := New()
	q.Push(1)

	assert.False(t, q.IsEmpty())
	assert.Equal(t, q.Size(), 1)

	q.Push(2)
	assert.Equal(t, q.Size(), 2)
}

func TestPeek_ReturnFront(t *testing.T) {
	q := New()
	q.Push(1)

	assert.Equal(t, 1, q.Size())
	val, _ := q.Peek()
	assert.Equal(t, 1, val)
	assert.Equal(t, 1, q.Size())

	q.Push(2)
	q.Push(3)
	assert.Equal(t, 3, q.Size())

	val, _ = q.Peek()
	assert.Equal(t, 1, val)
	assert.Equal(t, 3, q.Size())
}

func TestPop_ReturnRemoveFront(t *testing.T) {
	q := New()
	q.Push(1)

	assert.False(t, q.IsEmpty())
	assert.Equal(t, 1, q.Size())

	pop1, _ := q.Pop()
	assert.Equal(t, 1, pop1)
	assert.True(t, q.IsEmpty())
}

func TestPushPop(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(5)

	assert.Equal(t, 2, q.Size())
	pop1, _ := q.Pop()
	assert.Equal(t, 1, pop1)
	pop2, _ := q.Pop()
	assert.Equal(t, 5, pop2)
	assert.True(t, q.IsEmpty())
}

func TestPeek_FromEmpty(t *testing.T) {
	q := New()
	_, err := q.Peek()
	assert.NotNil(t, err)
	assert.Equal(t, ErrQueueIsEmpty, err)
}

func TestPop_FromEmpty(t *testing.T) {
	q := New()
	_, err := q.Pop()
	assert.NotNil(t, err)
	assert.Equal(t, ErrQueueIsEmpty, err)
}

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue()
	assert.True(t, q.IsEmpty())

	q.Push(5)
	q.Push(1)
	q.Push(7)
	q.Push(2)
	q.Push(3)

	peek1, _ := q.Peek()
	assert.Equal(t, 1, peek1)
	assert.Equal(t, 5, q.Size())

	pop1, _ := q.Pop()
	assert.Equal(t, 1, pop1)
	assert.Equal(t, 4, q.Size())

	pop2, _ := q.Pop()
	assert.Equal(t, 2, pop2)

	pop3, _ := q.Pop()
	assert.Equal(t, 3, pop3)

	pop4, _ := q.Pop()
	assert.Equal(t, 5, pop4)

	pop5, _ := q.Pop()
	assert.Equal(t, 7, pop5)
}
