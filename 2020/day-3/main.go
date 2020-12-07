package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	input := readInput()

	answerP1 := calcTrees(input, 3, 1)
	log.Print("Part One: ", answerP1)

	answerP2 := calcTrees(input, 1, 1) * calcTrees(input, 3, 1) * calcTrees(input, 5, 1) * calcTrees(input, 7, 1) * calcTrees(input, 1, 2)
	log.Print("Part Two: ", answerP2)
}

func calcTrees(input [][]string, across, down int) int {
	var index int

	count := 0
	place := 0
	input = input[down:]

	for i, line := range input {
		if i%down != 0 {
			continue
		}

		place = place + across

		if place >= len(line) {
			index = (place % len(line))
		} else {
			index = place
		}

		if line[index] == "#" {
			count++
		}
	}

	return count
}

func readInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to read input.txt")
	}
	defer file.Close()

	input := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}

	return input
}
