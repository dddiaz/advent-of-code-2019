package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello Advent of Code 2019. 🌲🌲🌲")
	day1()
}

func day1() int64 {
	// read file line by line
	file, err := os.Open("/Users/ddiaz/go/src/github.com/dddiaz/advent-of-code-2019/day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var weights []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		if s, err := strconv.ParseInt(line, 10, 64); err == nil {
			weights = append(weights, s)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total int64 = 0

	for n := 0; n < len(weights); n++ {
		total = total + ModuleFuel(weights[n])
	}

	fmt.Println("Day 1 Part 1 Result:", total)
	return total

}
