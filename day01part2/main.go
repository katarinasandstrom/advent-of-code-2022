package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := readFile()
	totalHighest := calculate(input)
	fmt.Println(totalHighest)
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
	highest1 := 0
	highest2 := 0
	highest3 := 0
	partSum := 0
	for i := 0; i < len(input); i++ {
		if input[i] != "" {

			num, err := strconv.Atoi(input[i])
			if err != nil {
				fmt.Println("error", err)
			}
			partSum += num

		} else {
			if partSum > highest1 {
				highest3 = highest2
				highest2 = highest1
				highest1 = partSum
			} else if partSum > highest2 {
				highest3 = highest2
				highest2 = partSum
			} else if partSum > highest3 {
				highest3 = partSum
			}

			partSum = 0
			continue
		}
	}
	return highest1 + highest2 + highest3

}
