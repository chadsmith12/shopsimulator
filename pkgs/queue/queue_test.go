package queue

import (
	"testing"
)

func ensureEmpty[T any](t *testing.T, q *Queue[T]) {
	if q.Len() != 0 {
		t.Errorf("q.Len() = %d, want %d", q.Len(), 0)
	}

	_, hasValue := q.Peek()
	if hasValue {
		t.Errorf("q.Peek() has value, want false")
	}
}

func ensureSingleValue[T any](t *testing.T, q *Queue[T]) {
	if q.Len() != 1 {
		t.Errorf("q.Len() = %d, want %d", q.Len(), 1)
	}
}

func TestNew(t *testing.T) {
	q := Init[int]()
	ensureEmpty[int](t, q)
}

func TestEnqueue(t *testing.T) {
	q := Init[int]()
	q.Enqueue(42)

	ensureSingleValue[int](t, q)
}

func TestDeque(t *testing.T) {
	q := Init[int]()
	ensureEmpty[int](t, q)
	q.Enqueue(42)
	ensureSingleValue[int](t, q)

	q.Enqueue(69)
	if q.Len() != 2 {
		t.Errorf("q.Len() = %d, want %d", q.Len(), 2)
	}
}

func TestPeekTrueWhenHasValue(t *testing.T) {
	q := Init[int]()
	ensureEmpty[int](t, q)
	q.Enqueue(42)
	value, hasValue := q.Peek()

	if value != 42 {
		t.Errorf("q.Peek() = %d, want %d", value, 42)
	}
	if !hasValue {
		t.Errorf("q.Peek() = %v, want %v", hasValue, true)
	}
}

func TestPeekFalseWhenEmpty(t *testing.T) {
	q := Init[int]()

	_, hasValue := q.Peek()

	if hasValue {
		t.Errorf("q.Peek() = %v, want %v", hasValue, false)
	}
}

func TestStringWithValues(t *testing.T) {
	q := Init[int]()
	q.Enqueue(42)
	q.Enqueue(69)
	q.Enqueue(420)

	const want = "[42 69 420]"
	value := q.String()

	if want != value {
		t.Errorf("q.String() = %v, want %v", value, want)
	}
}
