package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TuringMachine represents a simple Turing machine.
type TuringMachine struct {
	headPosition int
	tape         []byte
	state        string
	rules        [][]string
}

// NewTuringMachine initializes a new TuringMachine with the given file path.
func NewTuringMachine(filePath string) (*TuringMachine, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	return parseTuringMachine(file)
}

// parseTuringMachine extracts necessary information from the file and returns a TuringMachine.
func parseTuringMachine(file *os.File) (*TuringMachine, error) {
	scanner := bufio.NewScanner(file)

	var tm TuringMachine

	if scanner.Scan() {
		headPosition, _ := strconv.Atoi(scanner.Text())
		tm.headPosition = headPosition
	}

	if scanner.Scan() {
		tm.tape = []byte(scanner.Text())
	}

	for scanner.Scan() {
		readLine := scanner.Text()
		if strings.TrimSpace(readLine) == "" {
			continue
		}
		tokens := strings.Fields(readLine)
		tm.rules = append(tm.rules, tokens)
	}

	return &tm, nil
}

// Run executes the Turing machine based on its current state and rules.
func (tm *TuringMachine) Run() {

	state := "0"
	for {
		for i := 0; i < len(tm.rules); i++ {
			if tm.rules[i][0] == state && tm.rules[i][1][0] == tm.tape[tm.headPosition] {
				tm.tape[tm.headPosition] = tm.rules[i][2][0]
				if tm.rules[i][3][0] == 'L' {
					tm.headPosition--
				} else {
					tm.headPosition++
				}
				state = tm.rules[i][4]

				tm.printTape()
			}
		}
	}
}

// printTape prints the current state of the tape.
func (tm *TuringMachine) printTape() {
	fmt.Println("")
	for j := 0; j < len(tm.tape); j++ {
		fmt.Print(string(tm.tape[j]))
	}
}

func main() {
	// Check for input args
	if len(os.Args) == 1 {
		fmt.Println("Argument not provided as input")
		fmt.Println("Ex. ./main some_file.txt")
		return
	}

	turingMachine, err := NewTuringMachine(os.Args[1])
	if err != nil {
		fmt.Println("Error creating Turing machine:", err)
		return
	}

	turingMachine.Run()
}
