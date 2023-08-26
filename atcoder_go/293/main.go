package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fmt.Scan: スペースor改行ごとにとる
// bufio Scan: 一行ごとにとる
// ざっくりO(N+M)

// 問題解く時、途中or不安なやつはTODOってかいとけ！
// 失敗せずに書けばちゃんと終わる。

var io = newIO()

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	var (
		n int
		m int
	)
	if _, err := fmt.Scan(&n); err != nil {
		return err
	}
	if _, err := fmt.Scan(&m); err != nil {
		return err
	}

	operations := make([]*Operation, 0, m)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < m; i++ {
		scanner.Scan()
		operations = append(operations, NewOperation(scanner.Text()))
	}

	game := NewGame(n)
	for _, operation := range operations {
		game.Apply(operation)
	}

	fmt.Printf("%d %d\n", game.cycleCount, len(game.gm)/2)
	return nil
}

type Operation struct {
	edgeOne Edge
	edgeTwo Edge
}

// TODO: error
func NewOperation(str string) *Operation {
	splitted := strings.Split(str, " ")
	ropeNum1, _ := strconv.Atoi(splitted[0])
	ropeNum2, _ := strconv.Atoi(splitted[2])
	return &Operation{
		edgeOne: *NewEdge(ropeNum1, splitted[1] == "R"),
		edgeTwo: *NewEdge(ropeNum2, splitted[3] == "R"),
	}
}

type Game struct {
	gm         GroupMap // it has 2 edge&Group pair per Group
	cycleCount int
}

// エイリアスはる必要ない
type GroupMap map[EdgeKey]*Group

func NewGame(n int) *Game {
	m := make(GroupMap, 2*n)
	for i := 1; i <= n; i++ {
		edgeOne := NewEdge(i, true)
		edgeTwo := NewEdge(i, false)

		m[edgeOne.ToKey()] = NewGroup(edgeOne, edgeTwo)
		m[edgeTwo.ToKey()] = NewGroup(edgeOne, edgeTwo)
	}
	return &Game{
		gm:         m,
		cycleCount: 0,
	}
}

// TODO: error
func (g *Game) Apply(operation *Operation) {
	group1 := g.gm[operation.edgeOne.ToKey()]
	group2 := g.gm[operation.edgeTwo.ToKey()]
	delete(g.gm, group1.edgeOne.ToKey())
	delete(g.gm, group2.edgeOne.ToKey())
	delete(g.gm, group1.edgeTwo.ToKey())
	delete(g.gm, group2.edgeTwo.ToKey())
	if isGroupEqual(group1, group2) {
		g.cycleCount++ // cycle created
		return
	}

	// merge group1 and group2
	newGroup := Merge(group1, group2, operation)
	g.gm[newGroup.edgeOne.ToKey()] = newGroup
	g.gm[newGroup.edgeTwo.ToKey()] = newGroup
}

type Group struct {
	edgeOne *Edge
	edgeTwo *Edge
}

func NewGroup(edgeOne, edgeTwo *Edge) *Group {
	return &Group{
		edgeOne: edgeOne,
		edgeTwo: edgeTwo,
	}
}

func isGroupEqual(group1, group2 *Group) bool {
	return group1.edgeOne == group2.edgeOne &&
		group1.edgeTwo == group2.edgeTwo
}

func Merge(group1 *Group, group2 *Group, operation *Operation) *Group {
	var (
		edgeOne *Edge
		edgeTwo *Edge
	)
	switch operation.edgeOne.ToKey() {
	case group1.edgeOne.ToKey():
		edgeOne = group1.edgeTwo
	case group1.edgeTwo.ToKey():
		edgeOne = group1.edgeOne
	case group2.edgeOne.ToKey():
		edgeOne = group2.edgeTwo
	case group2.edgeTwo.ToKey():
		edgeOne = group2.edgeOne
	default:
		// TODO: error
	}

	switch operation.edgeTwo.ToKey() {
	case group1.edgeOne.ToKey():
		edgeTwo = group1.edgeTwo
	case group1.edgeTwo.ToKey():
		edgeTwo = group1.edgeOne
	case group2.edgeOne.ToKey():
		edgeTwo = group2.edgeTwo
	case group2.edgeTwo.ToKey():
		edgeTwo = group2.edgeOne
	default:
		// TODO: error
	}

	// TODO: validate

	g := NewGroup(edgeOne, edgeTwo)
	return g
}

type EdgeKey string

type Edge struct {
	ropeNum int
	isRed   bool
}

func NewEdge(ropeNum int, isRed bool) *Edge {
	return &Edge{ropeNum: ropeNum, isRed: isRed}
}

func (e *Edge) ToKey() EdgeKey {
	return EdgeKey(fmt.Sprintf("%d,%v", e.ropeNum, e.isRed))
}

// -----------------------------------------------------------------------------------
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

func (io *IO) nextLine() string {
	io.scanner.Scan()
	return io.scanner.Text()
}

func (io *IO) printLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
