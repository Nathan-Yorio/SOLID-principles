package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check for input args
	if len(os.Args) == 1 {
		fmt.Println("Argument not provided as input")
		fmt.Println("Ex. ./main some_file.txt")
		return
	}

	// Open File
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var readLine string
	var headPosition int
	var tape []byte
	state := "0"

	if scanner.Scan() {
		readLine = scanner.Text()
		headPosition, _ = strconv.Atoi(readLine)
	}

	if scanner.Scan() {
		readLine = scanner.Text()
		tape = []byte(readLine)
	}

	ruleCount := 0
	rules := make([][]string, 16)

	for scanner.Scan() {
		readLine = scanner.Text()
		if strings.TrimSpace(readLine) == "" { //Check if we reached the end of the file
			continue
		}

		tokens := strings.Fields(readLine)	 // Extract only the number parts, not the space in between
		rules[ruleCount] = make([]string, 5) // Each "rule" is a 5 string long array
		copy(rules[ruleCount], tokens)		 // Copy the 5 tokens on a line into the current 5 string array
		ruleCount++							 // increase the number of rules
	}

	for {
		for i := 0; i < ruleCount; i++ {
			if rules[i][0] == state && rules[i][1][0] == tape[headPosition] {
				tape[headPosition] = rules[i][2][0]
				if rules[i][3][0] == 'L' {
					headPosition--
				} else {
					headPosition++
				}
				state = rules[i][4]

				fmt.Println("")
				for j := 0; j < len(tape); j++ {
					fmt.Print(string(tape[j]))
				}
			}
		}
	}
}

