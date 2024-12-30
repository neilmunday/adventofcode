package day12

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day int = 12

type GridPoint struct {
	x       int
	y       int
	crop    string
	edges   int // number of edges around this grid point
	visited bool
}

func (gp GridPoint) String() string {
	return fmt.Sprintf("{(%d, %d) = %s, edges = %d, visited = %t}", gp.x, gp.y, gp.crop, gp.edges, gp.visited)
}

type Grid struct {
	grid   [][]GridPoint
	width  int
	height int
}

func NewGrid(data string) *Grid {
	lines := strings.Split(data, "\n")

	gridWidth := len(lines[0])
	gridHeight := len(lines)

	grid := make([][]GridPoint, gridWidth)

	for x := 0; x < gridWidth; x++ {
		grid[x] = make([]GridPoint, gridHeight)
	}

	for y, line := range lines {
		for x, char := range line {
			grid[x][y].edges = 0
			grid[x][y].x = x
			grid[x][y].y = y
			grid[x][y].crop = string(char)
			grid[x][y].visited = false
		}
	}

	return &Grid{
		grid:   grid,
		width:  gridWidth,
		height: gridHeight,
	}
}

type Region struct {
	crop       string
	grid       *Grid
	gridPoints []*GridPoint
	sides      int
}

func NewRegion(crop string, grid *Grid) *Region {
	return &Region{
		crop:       crop,
		grid:       grid,
		gridPoints: make([]*GridPoint, 0),
		sides:      0,
	}
}

func (r *Region) AddGridPoint(gp *GridPoint) {
	r.gridPoints = append(r.gridPoints, gp)
}

func (r *Region) Area() int {
	return len(r.gridPoints)
}

func (r *Region) DiscountPrice() int {
	return r.Area() * r.Sides()
}

func (r *Region) Perimeter() int {
	total := 0
	for _, gp := range r.gridPoints {
		total += gp.edges
	}
	return total
}

func (r *Region) Price() int {
	return r.Area() * r.Perimeter()
}

func (r *Region) Sides() int {
	// no. of sides = no. of corners
	sides := 0
	grid := r.grid

	for _, gp := range r.gridPoints {
		x := gp.x
		y := gp.y

		// top left?
		if (x-1 >= 0 && grid.grid[x-1][y].crop != r.crop || x-1 == -1) && (y-1 >= 0 && grid.grid[x][y-1].crop != r.crop || y-1 == -1) {
			sides += 1
		}
		// top right?
		if (x+1 < grid.width && grid.grid[x+1][y].crop != r.crop || x+1 == grid.width) && (y-1 >= 0 && grid.grid[x][y-1].crop != r.crop || y-1 == -1) {
			sides += 1
		}
		// bottom left?
		if (x-1 >= 0 && grid.grid[x-1][y].crop != r.crop || x-1 == -1) && (y+1 < grid.height && grid.grid[x][y+1].crop != r.crop || y+1 == grid.height) {
			sides += 1
		}
		// bottom right?
		if (x+1 < grid.width && grid.grid[x+1][y].crop != r.crop || x+1 == grid.width) && (y+1 < grid.height && grid.grid[x][y+1].crop != r.crop || y+1 == grid.height) {
			sides += 1
		}
		// inside tl
		if x-1 >= 0 && y-1 >= 0 && grid.grid[x][y-1].crop == r.crop && grid.grid[x-1][y].crop == r.crop && grid.grid[x-1][y-1].crop != r.crop {
			sides += 1
		}
		// inside tr
		if x+1 < grid.width && y-1 >= 0 && grid.grid[x][y-1].crop == r.crop && grid.grid[x+1][y].crop == r.crop && grid.grid[x+1][y-1].crop != r.crop {
			sides += 1
		}
		// inside bl
		if x-1 >= 0 && y+1 < grid.height && grid.grid[x][y+1].crop == r.crop && grid.grid[x-1][y].crop == r.crop && grid.grid[x-1][y+1].crop != r.crop {
			sides += 1
		}
		// inside br
		if x+1 < grid.width && y+1 < grid.height && grid.grid[x][y+1].crop == r.crop && grid.grid[x+1][y].crop == r.crop && grid.grid[x+1][y+1].crop != r.crop {
			sides += 1
		}
	}

	return sides
}

func findRegion(crop string, grid *Grid, gp *GridPoint, region *Region) {
	if gp.crop != crop {
		return
	}

	gp.visited = true

	region.AddGridPoint(gp)

	if gp.x-1 >= 0 && !grid.grid[gp.x-1][gp.y].visited {
		findRegion(crop, grid, &grid.grid[gp.x-1][gp.y], region)
	}

	if gp.x+1 < grid.width && !grid.grid[gp.x+1][gp.y].visited {
		findRegion(crop, grid, &grid.grid[gp.x+1][gp.y], region)
	}

	if gp.y-1 >= 0 && !grid.grid[gp.x][gp.y-1].visited {
		findRegion(crop, grid, &grid.grid[gp.x][gp.y-1], region)
	}

	if gp.y+1 < grid.height && !grid.grid[gp.x][gp.y+1].visited {
		findRegion(crop, grid, &grid.grid[gp.x][gp.y+1], region)
	}
}

func Part1and2() {
	grid := NewGrid(input.GetFileContents(day))
	regions := make([]*Region, 0)

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {

			gp := &grid.grid[x][y]

			// work out no. of edges
			if x == 0 || x == grid.width-1 {
				gp.edges += 1
			}

			if y == 0 || y == grid.height-1 {
				gp.edges += 1
			}

			if x-1 >= 0 && grid.grid[x-1][y].crop != gp.crop {
				gp.edges += 1
			}

			if x+1 < grid.width && grid.grid[x+1][y].crop != gp.crop {
				gp.edges += 1
			}

			if y-1 >= 0 && grid.grid[x][y-1].crop != gp.crop {
				gp.edges += 1
			}

			if y+1 < grid.height && grid.grid[x][y+1].crop != gp.crop {
				gp.edges += 1
			}

			if !grid.grid[x][y].visited {
				region := NewRegion(grid.grid[x][y].crop, grid)
				findRegion(region.crop, grid, &grid.grid[x][y], region)
				regions = append(regions, region)
			}
		}
	}

	totalPrice := 0
	totalDiscountPrice := 0

	for _, region := range regions {
		totalPrice += region.Price()
		totalDiscountPrice += region.DiscountPrice()
	}

	fmt.Printf("Part 1 answer: %d\n", totalPrice)
	fmt.Printf("Part 2 answer: %d\n", totalDiscountPrice)
}
