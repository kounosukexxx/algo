package stack

import "testing"

func TestQueue(t *testing.T) {
	q := newQueue(0)
	t.Log(q.len())
	t.Log(q.isEmpty())

	q.push(1)
	t.Log(q)
	t.Log(q.front())
	t.Log(q.len())
	t.Log(q.isEmpty())

	q.push(2)
	t.Log(q.front())
	t.Log(q.len())
	t.Log(q.isEmpty())

	q.pop()
	t.Log(q.front())
	t.Log(q.len())
	t.Log(q.isEmpty())

	q.pop()
	t.Log(q.len())
	t.Log(q.isEmpty())
}
