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
	strs := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		strs = append(strs, io.nextString())
	}
	input := input{
		strs: strs,
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
	strs []string
}

// --------------------impl run---------------------------

func run(input *input) ([]string, error) {
	return input.strs, nil
}
