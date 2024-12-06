package day5

import (
	"aoc/input"
	"fmt"
	"slices"
	"sort"
	"strings"
)

const day = 5

/*
Helper function to remove an element from an array
where ordering is not important.
*/
func remove(s []string, i int) []string {
	// set element to remove to the last item in the array
	s[i] = s[len(s)-1]
	// return all items minus the last element, thus removing it
	return s[:len(s)-1]
}

func Part1and2() (int, int) {
	// we can solve both parts in one go
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	ruleMap := make(map[int]map[int]bool)
	ruleParsingDone := false

	part1Total := 0
	part2Total := 0

	for _, line := range lines {
		stripped := strings.TrimSpace(line)

		if stripped == "" {
			ruleParsingDone = true
			continue
		}

		if !ruleParsingDone {
			pages := strings.Split(stripped, "|")

			num1 := input.GetInt(pages[0])
			num2 := input.GetInt(pages[1])

			if ruleMap[num1] == nil {
				ruleMap[num1] = make(map[int]bool)
			}

			if !ruleMap[num1][num2] {
				ruleMap[num1][num2] = true
			}

		} else {
			updatesOk := true

			pages := strings.Split(stripped, ",")

			for index, pageStr := range pages {
				page := input.GetInt(pageStr)

				for i := index + 1; i+1 <= len(pages); i++ {
					comparePage := input.GetInt(pages[i])
					if !ruleMap[page][comparePage] {
						updatesOk = false
						break
					}
				}
			}

			if updatesOk {
				// get middle number of the array
				middle := input.GetInt(pages[len(pages)/2])
				part1Total += int(middle)
			} else {
				// need to fix ordering
				newPagesMap := make(map[int]int)
				countArray := make([]int, len(pages))

				for index, pageStr := range pages {
					page := input.GetInt(pageStr)

					// if page is not in map, it has to go at the end
					rule, ok := ruleMap[page]
					if !ok {
						newPagesMap[page] = 0
					} else {
						// how many of the other numbers are in the rule?
						count := 0
						searchPages := remove(slices.Clone(pages), index)
						for _, searchPageStr := range searchPages {
							searchPage := input.GetInt(searchPageStr)
							if rule[searchPage] {
								count += 1
							}
						}
						newPagesMap[count] = page
						countArray[index] = count
					}
				}
				// sort by page hits
				sort.Sort(sort.Reverse(sort.IntSlice(countArray)))
				// now we just the page that is in the middle
				middle := len(countArray) / 2
				part2Total += newPagesMap[middle]
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n", part1Total)
	fmt.Printf("Part 2 answer: %d\n", part2Total)

	return part1Total, part2Total
}
