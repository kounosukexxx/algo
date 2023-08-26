package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// NOTE: 無理だった！解説も数学的な感じでワカラン。
// もっと早く答えみていい。

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
	T, _ := io.nextInt()
	input := make([]*test, 0, T)
	for i := 0; i < T; i++ {
		t := &test{}
		t.N, _ = io.nextInt()
		t.D, _ = io.nextInt()
		t.K, _ = io.nextInt()
		input = append(input, t)
	}

	// run
	// before := time.Now()
	ans, err := run(input)
	// log.Println(time.Since(before))

	// output
	if err != nil {
		io.printLn(err)
	} else {
		for _, v := range ans {
			io.printLn(v)
		}
	}
	io.writer.Flush()

	// exit
	// if err != nil {
	// 	os.Exit(1)
	// }
	// os.Exit(0)
}

// --------------------impl run---------------------------
type test struct {
	N int
	D int
	K int

	SortKey string
}

func sortTests(tests []*test) {
	for _, t := range tests {
		t.SortKey = strconv.Itoa(t.N) + strconv.Itoa(t.D) + strconv.Itoa(t.K)
	}
	sort.Slice(tests, func(i, j int) bool {
		return tests[i].SortKey < tests[j].SortKey
	})
}

func run(tests []*test) ([]int, error) {
	sortTests(tests)

	ans := make([]int, 0, len(tests))
	var sd savedData
	for _, t := range tests {
		ans = append(ans, doTest(t, &sd))
	}
	return ans, nil
}

type savedData struct {
	N             int
	D             int
	lastA         int
	lastX         int
	lastIsMarked  []bool
	lastMarkCount int
}

func (sd *savedData) upsertSaveData(
	t *test,
	lastA,
	lastX,
	lastMarkCount int,
	lastIsMarked []bool,
) {
	sd.N = t.N
	sd.D = t.D
	sd.lastA = lastA
	sd.lastX = lastX
	sd.lastMarkCount = lastMarkCount
	sd.lastIsMarked = lastIsMarked
}

func (sd savedData) needsRestore(t *test) bool {
	return t.N == sd.N && t.D == sd.D
}

func doTest(t *test, sd *savedData) int {
	// prepare
	var (
		x         int
		A         int // may be unnecessary
		isMarked  []bool
		markCount int
	)
	if sd.needsRestore(t) {
		x = sd.lastX
		A = sd.lastA
		isMarked = sd.lastIsMarked
		markCount = sd.lastMarkCount
	} else {
		x = 0
		A = 0
		isMarked = make([]bool, t.N)
		isMarked[0] = true
		markCount = 1
	}

	for {
		// save
		if markCount == t.K {
			sd.upsertSaveData(t, A, x, markCount, isMarked)
			break
		}

		// decide where to mark
		x = (A + t.D) % t.N
		for isMarked[x] {
			x = (x + 1) % t.N
		}

		// mark
		isMarked[x] = true
		markCount++
		A = x
	}

	return x
}
