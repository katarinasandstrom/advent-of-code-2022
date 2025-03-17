package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := readFile()
	highestSum := calculate(input)
	fmt.Println(highestSum)
}

func readFile() []string {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Fel vid öppning av fil: ", err)
		return nil
	}

	var text []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func calculate(input []string) int {
	highest := 0
	partSum := 0
	for i := 0; i < len(input); i++ {
		if input[i] != "" {

			num, err := strconv.Atoi(input[i])
			if err != nil {
				fmt.Println("error", err)
			}
			partSum += num

			if partSum > highest {
				highest = partSum
			}

		} else {
			partSum = 0
			continue
		}
	}
	return highest

}
