package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// -------------------prepare IO-----------------------

type IO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func newIO() *IO {
	return &IO{
		scanner: bufio.NewScanner(os.Stdin),
		writer:  bufio.NewWriter(os.Stdout),
	}
}

func (io *IO) nextString() string {
	io.scanner.Scan()
	return io.scanner.Text()
}

func (io *IO) nextChar() (rune, error) {
	io.scanner.Scan()
	str := io.scanner.Text()
	runes := []rune(str)
	if len(runes) != 1 {
		return ' ', errors.New("invalid character length")
	}
	return runes[0], nil
}

func (io *IO) nextInt() (int, error) {
	return strconv.Atoi(io.nextString())
}

func (io *IO) printLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

var io = newIO()

// -----------------------main-----------------------------

func main() {
	// prepare io
	io.scanner.Split(bufio.ScanWords)      // switch to separating by space
	io.scanner.Buffer([]byte{}, 100000007) // switch to read large size input

	// prepare input
	N, _ := io.nextInt() // TODO:
	s2t := make(map[string]string, N)
	t2s := make(map[string]string, N)
	for i := 0; i < N; i++ {
		s := io.nextString()
		t := io.nextString()
		s2t[s] = t
		t2s[t] = s
	}
	input := input{
		s2t: s2t,
		t2s: t2s,
	}

	// run
	ans, err := run(&input)

	// output
	if err != nil {
		io.printLn(err)
	} else {
		io.printLn(ans)
	}
	io.writer.Flush()

	// exit
	// if err != nil {
	// 	os.Exit(1)
	// }
	// os.Exit(0)
}

type input struct {
	s2t map[string]string
	t2s map[string]string
}

// --------------------impl run---------------------------

func run(input *input) (string, error) {
	tQueueToCheckIfExistsInT := newQueue(len(input.s2t))
	for _, t := range input.s2t {
		if _, existInS := input.s2t[t]; !existInS {
			tQueueToCheckIfExistsInT.push(t)
		}
	}

	for !tQueueToCheckIfExistsInT.isEmpty() {
		tToCheckIfExistsInT := tQueueToCheckIfExistsInT.front()
		tQueueToCheckIfExistsInT.pop()

		s, existInT := input.t2s[tToCheckIfExistsInT]
		if existInT {
			tQueueToCheckIfExistsInT.push(s)
			delete(input.s2t, s)
			delete(input.t2s, tToCheckIfExistsInT)
		}
	}

	if len(input.s2t) == 0 {
		return "Yes", nil
	}

	return "No", nil
}

// --------------------impl queue-------------------------

type strQueue struct {
	data []string
}

func newQueue(cap int) *strQueue {
	return &strQueue{
		data: make([]string, 0, cap),
	}
}

func (q strQueue) len() int {
	return len(q.data)
}

func (q strQueue) isEmpty() bool {
	return q.len() == 0
}

func (q strQueue) front() string {
	return q.data[0]
}

func (q *strQueue) pop() {
	q.data = q.data[1:]
}

func (q *strQueue) push(i string) {
	q.data = append(q.data, i)
}
