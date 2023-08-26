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
	A := make([]int, 0, N)
	for i := 0; i < N; i++ {
		a, _ := io.nextInt() // TODO:
		A = append(A, a)
	}

	M, _ := io.nextInt() // TODO:
	B := make(map[int]struct{}, M)
	for i := 0; i < M; i++ {
		b, _ := io.nextInt() // TODO:
		B[b] = struct{}{}
	}

	X, _ := io.nextInt() // TODO:
	input := input{
		x: X,
		A: A,
		B: B,
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

// --------------------impl run---------------------------

type input struct {
	x int
	A []int
	B map[int]struct{}
}

func run(input *input) (string, error) {
	dp := make([]bool, input.x+1) // true: can't reach index, false: unknown
	if canReach(input.x, dp, input) {
		return "Yes", nil
	}
	return "No", nil
}

func canReach(dst int, dp []bool, input *input) bool {
	if dst == 0 {
		return true
	}
	if dst < 0 {
		return false
	}
	if _, ok := input.B[dst]; ok {
		return false
	}
	if dp[dst] {
		return false
	}

	for i := 0; i < len(input.A); i++ {
		if canReach(dst-input.A[i], dp, input) {
			return true
		}
	}

	dp[dst] = true
	return false
}
