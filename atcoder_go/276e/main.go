package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		if _, err2 := fmt.Fprintln(os.Stderr, err); err2 != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func run() error {
	// get input
	input, err := getInput()
	if err != nil {
		return err
	}

	// do a work
	output, err := do(input)
	if err != nil {
		return err
	}

	// output
	_, err = fmt.Println(output)
	if err != nil {
		return err
	}

	return nil
}

// -------------------------input--------------------------
func nextString() (string, error) {
	var str string
	_, err := fmt.Scan(&str)
	return str, err
}

func nextInt() (int, error) {
	var num int
	_, err := fmt.Scan(&num)
	return num, err
}

func getInput() (*input, error) {
	h, _ := nextInt()
	_, _ = nextInt()
	grid := make([][]rune, h)
	input := &input{}
	for i := 0; i < h; i++ {
		str, _ := nextString()
		// こっちでも可
		// for j := 0; j < w; j++ {
		// 	r := rune(str[j])
		// 	if r == 'S' {
		// 		input.startRow = i
		// 		input.startColumn = j
		// 	}
		// 	grid[i] = append(grid[i], r)
		// }
		for j, r := range str {
			if r == 'S' {
				input.startRow = i
				input.startColumn = j
			}
			grid[i] = append(grid[i], r)
		}
	}

	input.grid = grid
	return input, nil
}

// --------------------------impl----------------------------

type input struct {
	grid        [][]rune
	startRow    int
	startColumn int
}

func do(input *input) (string, error) {
	visited := make([][]bool, len(input.grid))
	for i := 0; i < len(input.grid); i++ {
		visited[i] = make([]bool, len(input.grid[i]))
	}

	if isLongRoundTrip(input.grid, visited, 0, FROM_UNSPECIFIED, input.startRow, input.startColumn) {
		return "Yes", nil
	}
	return "No", nil
}

type from int

const (
	FROM_UNSPECIFIED from = iota
	FROM_ABOVE
	FROM_RIGHT
	FROM_LEFT
	FROM_BELOW
)

func isLongRoundTrip(grid [][]rune, visited [][]bool, currLen int, from from, row, column int) bool {
	if isUnreachable(grid, row, column) {
		return false
	}
	if grid[row][column] == 'S' && from != FROM_UNSPECIFIED {
		return currLen >= 4
	}

	if visited[row][column] {
		return false
	}
	visited[row][column] = true

	if from != FROM_ABOVE && isLongRoundTrip(grid, visited, currLen+1, FROM_BELOW, row-1, column) {
		return true
	}
	if from != FROM_RIGHT && isLongRoundTrip(grid, visited, currLen+1, FROM_LEFT, row, column+1) {
		return true
	}
	if from != FROM_LEFT && isLongRoundTrip(grid, visited, currLen+1, FROM_RIGHT, row, column-1) {
		return true
	}
	if from != FROM_BELOW && isLongRoundTrip(grid, visited, currLen+1, FROM_ABOVE, row+1, column) {
		return true
	}

	return false
}

func isUnreachable(grid [][]rune, row, column int) bool {
	return row >= len(grid) ||
		row < 0 ||
		column >= len(grid[0]) ||
		column < 0 ||
		grid[row][column] == '#'
}
