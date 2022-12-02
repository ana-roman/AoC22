package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Top3 struct {
	first  int
	second int
	third  int
}

func Day1_2() {
	f, err := os.Open("day1/input_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	var value string
	var top Top3

	for scanner.Scan() {
		if scanner.Text() == "" {
			findTopThree(&top, sum)
			sum = 0
		} else {
			value = scanner.Text()
			intval, _ := strconv.Atoi(value)
			sum += intval
		}
	}

	// Handle the last sum, when there is no more empty line at the end
	findTopThree(&top, sum)

	fmt.Println("Max 3 sum:", top.first+top.second+top.third)
}

func findTopThree(top *Top3, sum int) {
	if sum != 0 {
		if sum > top.first {
			top.third = top.second
			top.second = top.first
			top.first = sum
		} else if sum > top.second {
			top.third = top.second
			top.second = sum
		} else if sum > top.third {
			top.third = sum
		}
	}
}
