package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readFile()
	totalScore := calcScore(input)
	fmt.Println(totalScore)
}

func readFile() [][]rune {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Fel vid öppning av fil: ", err)
		return nil
	}

	var input [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 {

			line := []rune{[]rune(fields[0])[0], []rune(fields[1])[0]}
			input = append(input, line)
		}

	}
	return input
}

func calcScore(input [][]rune) int {
	totalScore := 0

	for i := 0; i < len(input); i++ {

		r1 := input[i][0]
		r2 := input[i][1]

		switch {
		case r1 == 'A' && r2 == 'X':
			totalScore += 4
		case r1 == 'A' && r2 == 'Y':
			totalScore += 8
		case r1 == 'A' && r2 == 'Z':
			totalScore += 3
		case r1 == 'B' && r2 == 'X':
			totalScore += 1
		case r1 == 'B' && r2 == 'Y':
			totalScore += 5
		case r1 == 'B' && r2 == 'Z':
			totalScore += 9
		case r1 == 'C' && r2 == 'X':
			totalScore += 7
		case r1 == 'C' && r2 == 'Y':
			totalScore += 2
		case r1 == 'C' && r2 == 'Z':
			totalScore += 6
		}
	}
	return totalScore
}
