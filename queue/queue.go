package queue

import (
	"fmt"
	"sort"
)

var (
	ErrQueueIsEmpty = fmt.Errorf("queue is empty")
)

// New ...
func New() *Queue {
	return &Queue{}
}

func NewPriorityQueue() *Queue {
	return &Queue{isPriorityQueue: true}
}

/*
Requirements:
- Queue of integer
- Should be able to check whether a queue is empty or not
- Should has function to return size of the queue
- Push should add the data to the head
- Peek should return the data on the head,
	but doesnt remove it from the queue
- Pop should return the data on the head,
	and remove it from the queue
- When peek and pop from empty queue, return error
- If priority queue is enabled, then all smaller
	number will go straight to the head of queue
*/

// Queue ...
type Queue struct {
	data            []int
	isPriorityQueue bool
}

func (q Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Push(num int) {
	q.data = append(q.data, num)
	if q.isPriorityQueue {
		sort.Ints(q.data)
	}
}

func (q Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return -1, ErrQueueIsEmpty
	}
	return q.data[0], nil
}

func (q *Queue) Pop() (int, error) {
	val, err := q.Peek()
	if err != nil {
		return val, ErrQueueIsEmpty
	}
	q.data = q.data[1:]
	return val, nil
}
