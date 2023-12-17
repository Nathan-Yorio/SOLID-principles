package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No file given\n\n")
		return
	}

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
	rules := make([][]string, 100)

	for scanner.Scan() {
		readLine = scanner.Text()
		if strings.TrimSpace(readLine) == "" {
			continue
		}

		tokens := strings.Fields(readLine)
		rules[ruleCount] = make([]string, 5)
		copy(rules[ruleCount], tokens)
		ruleCount++
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
