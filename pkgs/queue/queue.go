package queue

import (
	"errors"
	"fmt"
	"strings"
)

var EmptyQueueError = errors.New("Queue is empty")

type Node[T any] struct {
	next  *Node[T]
	value T
}

type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func Init[T any]() *Queue[T] {
	q := &Queue[T]{length: 0, head: nil, tail: nil}

	return q
}

func (q *Queue[T]) String() string {
	if q.length == 0 {
		return "[]"
	}

	var builder strings.Builder
	currentNode := q.head
	builder.WriteByte('[')
	for currentNode != nil {
		fmt.Fprintf(&builder, "%v", currentNode.value)
		if currentNode.next != nil {
			builder.WriteByte(' ')
		}
		currentNode = currentNode.next
	}
	builder.WriteByte(']')

	return builder.String()
}

func (q *Queue[T]) Enqueue(item T) {
	newNode := &Node[T]{value: item}
	q.length++
	if q.length == 1 {
		q.tail, q.head = newNode, newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.head == nil {
		return *new(T), false
	}

	q.length--
	head := q.head
	q.head = q.head.next
	head.next = nil

	return head.value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		return *new(T), false
	}

	return q.head.value, true
}

func (q *Queue[T]) Len() int {
	return q.length
}
