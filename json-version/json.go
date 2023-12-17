package main

import (
	// "bufio"   // Also not needed with the JSON version
	"fmt"
	"os"
	"strconv"
	// "strings" // Not needed with the JSON version
	"encoding/json"
)

// Special struct just for holding TuringMachine state values
type TuringMachine struct {
	headPosition int
	tape         []byte
	state        string
	rules        [][]string
}

// Creates a new Turing Machine state instance with a given filePath
func NewTuringMachine(filePath string) (*TuringMachine, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open: encountered error opening file:", filePath, err)
		return nil, err
	}
	defer file.Close()

	return parseTuringMachine(file)
}

// rewritten to accomodate parsing a JSON file instead of the original txt format
// uses the encoding/json library instead of the strings library
func parseTuringMachine(file *os.File) (*TuringMachine, error) {
	var tm TuringMachine

	decoder := json.NewDecoder(file)

	// Struct the mirrors the format of the JSON version with the keys
	// which the token values will map to 
	var jsonStruct struct {
		HeadStartPosition string `json:"head-start-position"`
		Tape              string `json:"tape"`
		Rules             []struct {
			State     string `json:"state"`
			Read      string `json:"read"`
			Write     string `json:"write"`
			Move      string `json:"move"`
			NextState string `json:"next-state"`
		} `json:"rules"` //Top level
	}

	err := decoder.Decode(&jsonStruct)
	if err != nil {
		return nil, err
	}

	// First take in the first level of the structured JSON 
	// since the text file based version essentially just has it on a single line
	headPosition, err := strconv.Atoi(jsonStruct.HeadStartPosition)
	if err != nil {
		return nil, err
	}
	tm.headPosition = headPosition
	tm.tape = []byte(jsonStruct.Tape)

	// Convert and assign values to Rules
	for _, rule := range jsonStruct.Rules {
		tokens := []string{
			rule.State,
			rule.Read,
			rule.Write,
			rule.Move,
			rule.NextState,
		}
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
