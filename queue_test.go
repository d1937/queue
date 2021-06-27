package queue

import (
	"fmt"
	"testing"
)

func TestQueue_pop(t *testing.T) {
	q := New()
	q.Push("Hello")
	fmt.Println(q.Pop())
	fmt.Println(q.Qsize())
	q.Push("Hello")
	q.Push("Hello")
	q.Push("Hello")
	fmt.Println(q.Qsize())

}
