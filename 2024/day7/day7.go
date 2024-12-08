package day7

import (
	"aoc/input"
	"errors"
	"fmt"
	"strings"
)

const day = 7

func parseLine(line string) (int64, []int64) {
	splitByColon := strings.Split(line, ":")
	solution := input.GetInt64(strings.TrimSpace(splitByColon[0]))
	numberStrs := strings.Split(strings.TrimSpace(splitByColon[1]), " ")
	numbers := make([]int64, len(numberStrs))

	for i, n := range numberStrs {
		numbers[i] = input.GetInt64(strings.TrimSpace(n))
	}

	return solution, numbers
}

func concatOp(n int64, m int64) int64 {
	return input.GetInt64(fmt.Sprintf("%d%d", n, m))
}

func insertAtFront(num int64, numbers []int64) []int64 {
	newNumbers := make([]int64, len(numbers)-1)
	newNumbers[0] = num
	for i := 1; i < len(newNumbers); i++ {
		newNumbers[i] = numbers[i+1]
	}
	return newNumbers
}

/*
Evaluate the input right to left in search of a viable solution.
Requires less ops than left to right solution.
*/
func reverse_check(solution int64, numbers []int64) error {
	// get last number in array
	n := numbers[len(numbers)-1]

	if len(numbers) == 2 {
		if solution-n == numbers[0] {
			return nil
		}
		if solution/n == numbers[0] && solution%n == 0 {
			return nil
		}

		return fmt.Errorf("two values: %d, %d left that do not subtract or divide for %d", numbers[0], numbers[1], solution)
	}

	errDivide := true
	errMinus := true

	if solution%n == 0 {
		err := reverse_check(solution/n, numbers[:len(numbers)-1])
		if err == nil {
			errDivide = false
		}
	}

	if solution-n > 0 {
		err := reverse_check(solution-n, numbers[:len(numbers)-1])
		if err == nil {
			errMinus = false
		}
	}

	if errDivide && errMinus {
		return errors.New("all paths failed")
	}

	return nil
}

/*
Evaluate numbers left to right with option concat operation enabled.
*/
func forwardCheck(solution int64, numbers []int64, enableConcat bool) error {

	if len(numbers) == 2 {

		if numbers[0]+numbers[1] == solution {
			return nil
		}

		if numbers[0]*numbers[1] == solution {
			return nil
		}

		if enableConcat && concatOp(numbers[0], numbers[1]) == solution {
			return nil
		}

		return fmt.Errorf("no solution for %v", numbers)
	}

	// add
	newNum := numbers[0] + numbers[1]
	newNumbers := insertAtFront(newNum, numbers)
	addErr := forwardCheck(solution, newNumbers, enableConcat)

	// multiply
	newNum = numbers[0] * numbers[1]
	newNumbers = insertAtFront(newNum, numbers)
	mulErr := forwardCheck(solution, newNumbers, enableConcat)

	if enableConcat {
		// concat
		newNum = concatOp(numbers[0], numbers[1])
		newNumbers = make([]int64, len(numbers)-1)
		newNumbers[0] = newNum
		for i := 1; i < len(newNumbers); i++ {
			newNumbers[i] = numbers[i+1]
		}

		concatErr := forwardCheck(solution, newNumbers, enableConcat)

		if addErr != nil && mulErr != nil && concatErr != nil {
			return fmt.Errorf("all paths failed")
		}
	} else {
		if addErr != nil && mulErr != nil {
			return fmt.Errorf("all paths failed")
		}
	}

	return nil
}

func Part1and2() (int64, int64) {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	var part1 int64 = 0
	var part2 int64 = 0

	for _, line := range lines {
		solution, numbers := parseLine(line)

		err := reverse_check(solution, numbers)
		if err == nil {
			part1 += solution
		}

		err = forwardCheck(solution, numbers, true)
		if err == nil {
			part2 += solution
		}
	}

	fmt.Printf("Part 1 answer: %d\n", part1)
	fmt.Printf("Part 2 answer: %d\n", part2)

	return part1, part2
}
