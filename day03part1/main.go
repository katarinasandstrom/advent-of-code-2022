package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readFile()
	result := handleInput(input)
	fmt.Println(result)

}

func readFile() []string {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Fel vid öppning av fil: ", err)
		return nil
	}

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)

	}
	return input
}

func handleInput(input []string) int {
	totalSum := 0

	for _, line := range input {

		runes := []rune(line)
		half := len(runes) / 2

		left := runes[:half]
		right := runes[half:]

		leftSet := make(map[rune]bool, len(left))
		for _, r := range left {
			leftSet[r] = true
		}
		for _, r := range right {
			if leftSet[r] {
				totalSum += letterScore(r)
				break
			}
		}
	}

	return totalSum
}

func letterScore(r rune) int {

	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}
	return 0
}
