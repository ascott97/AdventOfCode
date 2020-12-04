package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	answerP1 := partOne(input)
	log.Print(answerP1)

	answerP2 := partTwo(input)
	log.Print(answerP2)
}

func partOne(input []int) int {
	s := map[int]bool{}
	for _, num := range input {
		s[num] = true
		if s[2020-num] == true {
			return num * (2020 - num)
		}
	}

	return 0
}

func partTwo(input []int) int {
	for _, num := range input {
		a := 2020 - num
		s := map[int]bool{}
		for _, num2 := range input {
			s[num2] = true
			if s[a-num2] == true {
				return num2 * num * (a - num2)
			}
		}
	}

	return 0
}

func readInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to read input.txt")
	}
	defer file.Close()

	input := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if number, err := strconv.Atoi(scanner.Text()); err == nil {
			input = append(input, number)
		} else {
			log.Fatal("Unable to read a number from input.txt")
		}
	}

	return input
}
