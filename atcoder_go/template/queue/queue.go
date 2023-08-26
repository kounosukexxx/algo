package stack

type typeQueue struct {
	data []int
}

func newQueue(cap int) *typeQueue {
	return &typeQueue{
		data: make([]type, 0, cap),
	}
}

func (q typeQueue) len() int {
	return len(q.data)
}

func (q typeQueue) isEmpty() bool {
	return q.len() == 0
}

func (q typeQueue) front() type {
	return q.data[0]
}

func (q *typeQueue) pop() {
	q.data = q.data[1:]
}

func (q *typeQueue) push(v type) {
	q.data = append(q.data, v)
}
