package day

import (
	"fmt"
	"strings"
)

type Day4 struct{}

func init() {
	Days.RegisterDay(4, &Day4{})
}

func (d *Day4) SolvePart1(input []byte) (string, error) {
	coords := make([][]string, 0)

	for line := range strings.SplitSeq(string(input), "\n") {
		x := make([]string, len(line))
		parts := strings.Split(line, "")
		copy(x, parts)

		coords = append(coords, x)
	}

	iters := 0
	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords[i]); j++ {
			if coords[i][j] == "." {
				continue
			}
			if validPaperAccess(coords, i, j) {
				iters++
			}
		}
	}

	return fmt.Sprintf("%+v", iters), nil
}

func (d *Day4) SolvePart2(input []byte) (string, error) {
	coords := make([][]string, 0)

	for line := range strings.SplitSeq(string(input), "\n") {
		x := make([]string, len(line))
		parts := strings.Split(line, "")
		copy(x, parts)

		coords = append(coords, x)
	}

	iters := 0

	for {
		prevIters := iters

		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords[i]); j++ {
				if coords[i][j] == "." {
					continue
				}
				if validPaperAccess(coords, i, j) {
					coords[i][j] = "."
					iters++
				}
			}
		}

		if iters == prevIters {
			break
		}
	}

	return fmt.Sprintf("%+v", iters), nil
}

func validPaperAccess(grid [][]string, row int, col int) bool {
	count := 0

	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}

			currentRow := row + r
			currentCol := col + c

			if currentRow >= 0 && currentRow < len(grid) && currentCol >= 0 && currentCol < len(grid[currentRow]) && grid[currentRow][currentCol] == "@" {
				count++
			}

			if count == 4 {
				return false
			}
		}
	}

	return true
}
