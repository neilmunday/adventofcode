package day10

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day int = 10

type GridPoint struct {
	x       int
	y       int
	height  int
	visited bool
}

func (gp *GridPoint) String() string {
	return fmt.Sprintf("{(%d, %d) height: %d, visited: %t}", gp.x, gp.y, gp.height, gp.visited)
}

type Grid struct {
	grid        [][]GridPoint
	width       int
	height      int
	startPoints []*GridPoint
}

func NewGrid(data string) *Grid {
	lines := strings.Split(data, "\n")

	gridWidth := len(lines[0])
	gridHeight := len(lines)

	grid := make([][]GridPoint, gridWidth)

	startPoints := make([]*GridPoint, 0)

	for x := 0; x < gridWidth; x++ {
		grid[x] = make([]GridPoint, gridHeight)
	}

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			grid[x][y].x = x
			grid[x][y].y = y
			if line[x] == '.' {
				grid[x][y].height = 1000
			} else {
				grid[x][y].height = input.GetInt(string(line[x]))
			}
			grid[x][y].visited = false

			if grid[x][y].height == 0 {
				startPoints = append(startPoints, &grid[x][y])
			}
		}
	}

	return &Grid{
		grid:        grid,
		width:       gridWidth,
		height:      gridHeight,
		startPoints: startPoints,
	}
}

func (grid *Grid) ResetVisited() {
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			grid.grid[x][y].visited = false
		}
	}
}

/*
Perform depth first search from given point.
*/
func dfs(grid *Grid, gp *GridPoint, recordVisit bool) int {

	if recordVisit {
		gp.visited = true
	}

	if gp.height == 9 {
		// end of trail found
		return 1
	}

	score := 0

	// west
	if gp.x-1 >= 0 && grid.grid[gp.x-1][gp.y].height == gp.height+1 && !grid.grid[gp.x-1][gp.y].visited {
		score += dfs(grid, &grid.grid[gp.x-1][gp.y], recordVisit)
	}
	// east
	if gp.x+1 < grid.width && grid.grid[gp.x+1][gp.y].height == gp.height+1 && !grid.grid[gp.x+1][gp.y].visited {
		score += dfs(grid, &grid.grid[gp.x+1][gp.y], recordVisit)
	}
	// north
	if gp.y-1 >= 0 && grid.grid[gp.x][gp.y-1].height == gp.height+1 && !grid.grid[gp.x][gp.y-1].visited {
		score += dfs(grid, &grid.grid[gp.x][gp.y-1], recordVisit)
	}
	// south
	if gp.y+1 < grid.height && grid.grid[gp.x][gp.y+1].height == gp.height+1 && !grid.grid[gp.x][gp.y+1].visited {
		score += dfs(grid, &grid.grid[gp.x][gp.y+1], recordVisit)
	}
	return score
}

func Part1() {
	grid := NewGrid(input.GetFileContents(day))

	total := 0

	for _, startPoint := range grid.startPoints {
		startPoint.visited = true
		total += dfs(grid, startPoint, true)
		grid.ResetVisited()
	}

	fmt.Printf("Part 1 answer: %d\n", total)
}

func Part2() {
	grid := NewGrid(input.GetFileContents(day))

	total := 0

	for _, startPoint := range grid.startPoints {
		startPoint.visited = true
		total += dfs(grid, startPoint, false)
	}

	fmt.Printf("Part 2 answer: %d\n", total)
}
