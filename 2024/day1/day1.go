package day1

import (
	"aoc/input"
	"fmt"
	"sort"
	"strings"
)

func Part1() {
	contents := input.GetFileContents(1)

	lines := strings.Split(contents, "\n")

	var list1 = make([]int, len(lines))
	var list2 = make([]int, len(lines))

	for _, line := range lines {
		stripped := strings.TrimSpace(line)
		numbers := strings.Fields(stripped)

		leftNumber := input.GetInt(numbers[0])
		rightNumber := input.GetInt(numbers[1])

		list1 = append(list1, int(leftNumber))
		list2 = append(list2, int(rightNumber))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var totalDiff = 0

	for i, _ := range list1 {
		if list1[i] > list2[i] {
			totalDiff += list1[i] - list2[i]
		} else {
			totalDiff += list2[i] - list1[i]
		}
	}

	fmt.Printf("Part 1 answer: %d\n", totalDiff)
}

func Part2() {
	contents := input.GetFileContents(1)

	lines := strings.Split(contents, "\n")

	var list1 = make([]int, len(lines))
	var list2 = make([]int, len(lines))

	for _, line := range lines {
		stripped := strings.TrimSpace(line)
		numbers := strings.Fields(stripped)

		leftNumber := input.GetInt(numbers[0])
		rightNumber := input.GetInt(numbers[1])

		list1 = append(list1, int(leftNumber))
		list2 = append(list2, int(rightNumber))
	}

	rightMap := make(map[int]int)

	for _, n := range list2 {
		rightMap[n] += 1
	}

	var total = 0

	for _, leftNumber := range list1 {
		// how many times does leftNumber appear in list2?
		occurences := rightMap[leftNumber]
		total += leftNumber * occurences
	}

	fmt.Printf("Part 2 answer: %d\n", total)
}
