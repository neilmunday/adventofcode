package day13

import (
	"aoc/input"
	"fmt"
	"regexp"
	"strings"
)

const day int = 13

type Equation struct {
	x1   int
	x2   int
	y1   int
	y2   int
	ansX int
	ansY int
}

func (equation *Equation) IterativeSolve(iterations int) (int, int) {
	// we're told a and b will not be pressed more than 100 times
	// so we can loop through all combinations until we find a hit
	for a := 0; a < iterations; a++ {
		for b := 0; b < iterations; b++ {
			if (equation.x1*a)+(equation.y1*b) == equation.ansX && (equation.x2*a)+(equation.y2*b) == equation.ansY {
				return a, b
			}
		}
	}

	return 0, 0
}

func (equation *Equation) CramersSolve() (int, int) {
	D := (equation.x1 * equation.y2) - (equation.y1 * equation.x2)
	Dx := (equation.ansX * equation.y2) - (equation.y1 * equation.ansY)
	Dy := (equation.x1 * equation.ansY) - (equation.x2 * equation.ansX)

	// check for remainder - if there is one, no solution found
	if Dx%D != 0 || Dy%D != 0 {
		return 0, 0
	}

	a := Dx / D
	b := Dy / D

	return a, b
}

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	buttonRe := regexp.MustCompile(`^Button ([A|B]): X\+([\d]+), Y\+([\d]+)$`)
	prizeRe := regexp.MustCompile(`^Prize: X=([\d]+), Y=([\d]+)$`)

	equation := Equation{}

	tokens1 := 0
	tokens2 := 0

	for _, line := range lines {
		buttonMatches := buttonRe.FindStringSubmatch(line)
		if len(buttonMatches) > 0 {
			button := buttonMatches[1]
			x := input.GetInt(buttonMatches[2])
			y := input.GetInt(buttonMatches[3])

			if button == "A" {
				equation.x1 = x
				equation.x2 = y
			} else if button == "B" {
				equation.y1 = x
				equation.y2 = y
			} else {
				panic(fmt.Sprintf("Unrecognised button: %s\n", button))
			}
		} else {
			prizeMatches := prizeRe.FindStringSubmatch(line)
			if len(prizeMatches) > 0 {
				equation.ansX = input.GetInt(prizeMatches[1])
				equation.ansY = input.GetInt(prizeMatches[2])

				// part 1
				ansA, ansB := equation.IterativeSolve(100)

				if ansA != 0 && ansB != 0 {
					tokens1 += int((ansA * 3) + ansB)
				}

				// part 2
				equation.ansX += 10000000000000
				equation.ansY += 10000000000000

				ansA, ansB = equation.CramersSolve()

				if ansA != 0 && ansB != 0 {
					tokens2 += int((ansA * 3) + ansB)
				}
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n", tokens1)
	fmt.Printf("Part 2 answer: %d\n", tokens2)
}
