package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	Min    int
	Max    int
	Letter string
}

func (p *PasswordPolicy) New(policy string) {
	split := strings.Split(policy, " ")
	p.Letter = split[1]

	minMax := strings.Split(split[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	p.Min = min
	p.Max = max

	return
}

func main() {
	input := readInput()

	p1 := partOne(input)
	log.Print("PartOne: ", p1)

	p2 := partTwo(input)
	log.Print("PartTwo: ", p2)
}

func partOne(input [][]string) int {
	count := 0
	for _, v := range input {
		a := PartOnePasswordValid(v[0], v[1])
		if a {
			count++
		}
	}

	return count
}

func partTwo(input [][]string) int {
	count := 0
	for _, v := range input {
		a := PartTwoPasswordValid(v[0], v[1])
		if a {
			count++
		}
	}

	return count
}

func PartOnePasswordValid(policy, password string) bool {
	p := PasswordPolicy{}
	p.New(policy)

	c := strings.Count(password, p.Letter)
	if c >= p.Min && c <= p.Max {
		return true
	}

	return false
}

func PartTwoPasswordValid(policy, password string) bool {

	p := PasswordPolicy{}
	p.New(policy)

	count := 0

	if string(password[p.Min]) == p.Letter {
		count++
	}

	if string(password[p.Max]) == p.Letter {
		count++
	}

	if count == 1 {
		return true
	}

	return false
}

func readInput() [][]string {
	file, err := os.Open("passwords.txt")
	if err != nil {
		log.Fatal("Unable to read passwords.txt")
	}
	defer file.Close()

	input := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ":"))
	}

	return input
}
