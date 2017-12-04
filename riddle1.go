package main

import "flag"
import "fmt"
import "strconv"

type solver struct {
}

func new() solver {
	return solver{}
}

func (s *solver) checkCapcha(input string) int {

	var sum int

	for index := range input {
		var nextElement byte

		if index < len(input)-1 {
			nextElement = input[index+1]
		} else {
			nextElement = input[0]
		}

		if input[index] == nextElement {
			converted, _ := strconv.Atoi(string(nextElement))
			sum += converted
		}
	}

	return sum
}

func main() {
	var inputCapcha = flag.String("capcha", "", "capcha to check")
	flag.Parse()

	solver := new()

	solution := solver.checkCapcha(*inputCapcha)

	fmt.Println("Found capcha solution: ", solution)
}
