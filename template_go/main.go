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
	os.Exit(0)
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

func nextXAndY() (int, int, error) {
	var x, y int
	_, err := fmt.Scanf("%d.%d", &x, &y)
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func getInput() (*input, error) {
	strs := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		str, err := nextString()
		if err != nil {
			return nil, err
		}
		strs = append(strs, str)
	}
	return &input{
		strs: strs,
	}, nil
}

// --------------------------impl----------------------------

type input struct {
	strs []string
}

func do(input *input) ([]string, error) {
	return input.strs, nil
}
