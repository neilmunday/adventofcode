package day5

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day = 5

// Object to store a range
type Range struct {
	Start int
	End   int
}

// Return true if i exist in the range
func (r *Range) Exists(i int) bool {
	return i >= r.Start && i <= r.End
}

// Return the intersection of this range with another, otherwise return nil
func (r *Range) Intersect(o *Range) *Range {
	intersection := NewRange(0, 0)

	if o.Start >= r.Start && o.Start <= r.End {
		intersection.Start = o.Start
	} else if r.Start >= o.Start && r.Start <= o.End {
		intersection.Start = r.Start
	}

	if o.End >= r.Start && o.End <= r.End {
		intersection.End = o.End
	} else if r.End >= o.Start && r.End <= o.End {
		intersection.End = r.End
	}

	if intersection.Start == 0 || intersection.End == 0 {
		return nil
	}

	return intersection
}

// Return the size of the range
func (r *Range) Size() int {
	return r.End - r.Start + 1
}

// Create a new Range instance
func NewRange(start int, end int) *Range {
	return &Range{
		Start: start,
		End:   end,
	}
}

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	ranges := make([]*Range, 0)
	total := 0
	endOfRanges := false

	for _, line := range lines {
		if len(line) == 0 {
			// end of ranges
			endOfRanges = true
			continue
		}

		if !endOfRanges {
			numbers := strings.Split(line, "-")
			r := NewRange(input.GetInt(numbers[0]), input.GetInt(numbers[1]))
			ranges = append(ranges, r)
		} else {
			ingredient := input.GetInt(line)

			for _, r := range ranges {
				if r.Exists(ingredient) {
					total++
					break
				}
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n", total)
}

func Part2() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	ranges := make([]*Range, 0)

	for _, line := range lines {
		if len(line) == 0 {
			// end of ranges
			break
		}

		numbers := strings.Split(line, "-")
		r := NewRange(input.GetInt(numbers[0]), input.GetInt(numbers[1]))
		ranges = append(ranges, r)
	}

	// condense ranges until no more ranges can be removed
	for {
		currentLength := len(ranges)
		for i := range ranges {
			stop := false
			for j := range ranges {
				if i == j {
					continue
				}

				intersection := ranges[i].Intersect(ranges[j])

				// no intersection
				if intersection == nil {
					continue
				}

				// adjust ranges[i] start and end values
				if ranges[i].Start > ranges[j].Start {
					ranges[i].Start = ranges[j].Start
				}

				if ranges[i].End < ranges[j].End {
					ranges[i].End = ranges[j].End
				}

				// now remove j
				ranges = append(ranges[:j], ranges[j+1:]...)
				stop = true
				break
			}

			if stop {
				break
			}
		}

		if currentLength == len(ranges) {
			break
		}
	}

	total := 0

	// calculate the number of IDs that we are left with
	for _, r := range ranges {
		total += r.Size()
	}

	fmt.Printf("Part 2 answer: %d\n", total)
}
