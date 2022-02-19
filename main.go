package main

import "fmt"

type PQueue struct {
	list      []int
	min_index int
}

func NewPQ() *PQueue {
	value := PQueue{make([]int, 0, 8), 0}
	q := &value
	return q
}

func (q *PQueue) Push(ele int) {
	q.list = append(q.list, ele)
	if ele < q.list[q.min_index] {
		q.min_index = q.Size() - 1
	}
}

func (q *PQueue) Pop() (int, bool) {
	if q.Size() == 0 {
		return 0, false
	}

	poped := q.list[q.min_index]

	// remove the popped from the list
	q.list = append(q.list[:q.min_index], q.list[q.min_index+1:]...)

	// find the index of the new minnest one
	curr_min := int(^uint(0) >> 1) // max of int
	for i := range q.list {
		if q.list[i] < curr_min {
			curr_min = q.list[i]
			q.min_index = i
		}
	}

	return poped, true
}

func (q *PQueue) Size() int {
	return len(q.list)
}

func main() {
	pq := NewPQ()
	fmt.Println(pq)
	pq.Push(10)
	fmt.Println(pq)
	pq.Push(2)
	fmt.Println(pq)
	pq.Push(1)
	fmt.Println(pq)
	pq.Push(3)
	fmt.Println(pq)
	pq.Push(0)
	fmt.Println(pq)

	for pq.Size() > 0 {
		if ele, ok := pq.Pop(); ok {
			fmt.Println("pop:", ele, "q:", pq)
		}
	}
}
