package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	maxSum := 0
	sum := 0
	var value string

	for scanner.Scan() {
		if scanner.Text() == "" {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0

		} else {
			value = scanner.Text()
			intval, _ := strconv.Atoi(value)
			sum += intval
		}
	}

	fmt.Println("Maxsum:", maxSum)
}
