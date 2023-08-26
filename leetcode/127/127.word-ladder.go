package l127

import (
	"fmt"
)

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start

// TODO: Frontはいらん
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var graph graph
	found := false
	for _, word := range wordList {
		if word == beginWord {
			graph = initGraph(wordList)
			found = true
			break
		}
	}
	if !found {
		graph = initGraph(append(wordList, beginWord))
	}

	// BFS (should be graph's method)
	ans := 0
	q := NewQueue(len(wordList))
	pushed := make(map[string]struct{}, len(wordList)+1)

	// push beginWord
	q.Push(beginWord)
	pushed[beginWord] = struct{}{}

	// search endWord
	for !q.IsEmpty() {
		ans++

		// pop all words
		words := []string{}
		for !q.IsEmpty() {
			words = append(words, q.Front())
			q.Pop()
		}

		// check if popped words contains endWord, and push connected nodes into the queue
		for _, word := range words {
			if word == endWord {
				return ans // endWord found
			}

			// push connected nodes into the queue
			for _, node := range graph[word] {
				if _, ok := pushed[node]; ok {
					continue
				}
				q.Push(node)
				pushed[node] = struct{}{}
			}
		}
	}

	// endWord not found
	return 0
}

// queue
type Queue struct {
	data []string
	size int
}

// NewQueue instantiates a new queue
func NewQueue(cap int) *Queue {
	return &Queue{data: make([]string, 0, cap), size: 0}
}

// Push adds a new element at the end of the queue
func (q *Queue) Push(n string) {
	q.data = append(q.data, n)
	q.size++
}

// Pop removes the first element from queue
func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

// Front returns the first element of queue
func (q *Queue) Front() string {
	return q.data[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// String implements Stringer interface
func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}

// undirected string graph (list)
type graph map[string][]string

func initGraph(wordList []string) graph {
	graph := make(graph, len(wordList))
	for i := range wordList {
		for j := i + 1; j < len(wordList); j++ {
			if differsByOneChar(wordList[i], wordList[j]) {
				graph.addEdge(wordList[i], wordList[j])
			}
		}
	}
	return graph
}

func (g *graph) addEdge(w, v string) {
	(*g)[w] = append((*g)[w], v)
	(*g)[v] = append((*g)[v], w)
}

func differsByOneChar(w, v string) bool {
	count := 0
	for i := 0; i < len(w); i++ {
		if w[i] != v[i] {
			count++
		}
	}
	return count == 1
}

// @lc code=end
