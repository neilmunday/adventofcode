package day6

import (
	"aoc/input"
	"fmt"
	"regexp"
	"strings"
)

const day = 6

type Column struct {
	start int
	end   int
}

type Calc struct {
	col       Column
	vals      []string
	operation string
}

func NewCalc() *Calc {
	return &Calc{
		col: Column{
			start: 0,
			end:   0,
		},
		vals: make([]string, 0),
	}
}

func (c *Calc) AddVal(v string) {
	c.vals = append(c.vals, v)
}

func (c *Calc) Calc() int {

	result := input.GetInt(c.vals[0])

	for _, val := range c.vals[1:] {
		if c.operation == "+" {
			result += input.GetInt(val)
		} else {
			result *= input.GetInt(val)
		}
	}

	return result
}

func (c *Calc) CalcCephalopod() int {

	numRegEx, _ := regexp.Compile("[0-9]+")
	result := 0

	if c.operation == "*" {
		result = 1
	}

	for i := range c.col.end - c.col.start {
		numberStr := ""
		for _, valStr := range c.vals {
			if numRegEx.MatchString(string(valStr[i])) {
				numberStr += string(valStr[i])
			}
		}

		if c.operation == "+" {
			result += input.GetInt(numberStr)
		} else {
			result *= input.GetInt(numberStr)
		}
	}

	return result
}

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	calcs := make([]*Calc, 0)

	numRegEx, _ := regexp.Compile("[0-9]+")
	opRegEx, _ := regexp.Compile("[+*]")

	for _, valStr := range numRegEx.FindAllString(lines[0], -1) {
		calc := NewCalc()
		calc.AddVal(valStr)

		calcs = append(calcs, calc)
	}

	for _, line := range lines[1 : len(lines)-1] {
		for i, valStr := range numRegEx.FindAllString(line, -1) {
			calcs[i].AddVal(valStr)
		}
	}

	total := 0

	// process last line
	for i, opStr := range opRegEx.FindAllString(lines[len(lines)-1], -1) {
		calcs[i].operation = opStr
		total += calcs[i].Calc()
	}

	fmt.Printf("Part 1 answer: %d\n", total)
}

func Part2() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	calcs := make([]*Calc, 0)
	calc := NewCalc()
	calc.col = Column{
		start: 0,
		end:   0,
	}
	calc.operation = lines[len(lines)-1][0:1]

	calcs = append(calcs, calc)
	currentIndex := 0

	// look at the last line to find start and end columnns of numbers
	for i, char := range lines[len(lines)-1] {
		if i == 0 {
			continue
		}

		charStr := string(char)

		if charStr == "*" || charStr == "+" {
			calcs[currentIndex].col.end = i - 1
			calc = NewCalc()
			calc.col.start = i
			calc.col.end = 0
			calc.operation = charStr
			calcs = append(calcs, calc)
			currentIndex++
		}
	}

	calcs[currentIndex].col.end = len(lines[0])

	// parse number lines to create calculations
	for _, line := range lines[:len(lines)-1] {
		for _, calc := range calcs {
			calc.AddVal(line[calc.col.start:calc.col.end])
		}
	}

	total := 0

	// perform the Cephalopod maths
	for _, calc := range calcs {
		total += calc.CalcCephalopod()
	}

	fmt.Printf("Part 2 answer: %d\n", total)
}
