package day

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
)

type Day7 struct{}

func init() {
	Days.RegisterDay(7, &Day7{})
}

const (
	START    = "S"
	SPLITTER = "^"
	BEAM     = "|"
	EMPTY    = "."
)

func (d *Day7) SolvePart1(input []byte) (string, error) {
	count := 0
	tree := buildTree(input)
	previousRow := []string{}
	for idx, row := range tree {
		if len(previousRow) == 0 {
			previousRow = row
			continue
		}
		if idx < len(tree) {
			newRow, c := printTree(previousRow, row)
			count += c
			for _, t := range newRow {
				tree[idx][t] = BEAM
			}
		}
		previousRow = row
	}

	return fmt.Sprintf("%d", count), nil
}

func (d *Day7) SolvePart2(input []byte) (string, error) {
	initialTree := buildTree(input)

	// Convert initial tree to initial state
	initialBeams := []int{}
	for idx, cell := range initialTree[0] {
		if cell == START || cell == BEAM {
			initialBeams = append(initialBeams, idx)
		}
	}

	// Memoization: maps (row, stateHash) -> count of timelines reaching that state
	memo := make(map[string]int64)

	// DFS with memoization - count how many timelines reach each final state
	count := dfsCount(initialTree, 0, initialBeams, memo)

	return fmt.Sprintf("%d", count), nil
}

func dfsCount(tree [][]string, row int, beamPositions []int, memo map[string]int64) int64 {
	// Base case: if we've processed all rows, this represents 1 timeline
	if row >= len(tree)-1 {
		return 1
	}

	// Create state hash for memoization
	stateHash := hashState(row, beamPositions)
	memoKey := fmt.Sprintf("%d:%s", row, stateHash)

	// Check memoization cache
	if cached, exists := memo[memoKey]; exists {
		return cached
	}

	var count int64
	nextRow := tree[row+1]

	// Find unique splitter positions that are hit by beams
	splitterPositions := make(map[int]bool)
	for _, pos := range beamPositions {
		if pos < len(nextRow) && nextRow[pos] == SPLITTER {
			splitterPositions[pos] = true
		}
	}

	// Convert to sorted slice
	splitterPosList := make([]int, 0, len(splitterPositions))
	for pos := range splitterPositions {
		splitterPosList = append(splitterPosList, pos)
	}
	sort.Ints(splitterPosList)

	if len(splitterPosList) > 0 {
		// Generate all combinations: 2^n timelines for n unique splitters
		numCombinations := 1 << len(splitterPosList) // 2^n
		for combo := 0; combo < numCombinations; combo++ {
			newBeams := make(map[int]bool)

			// Process each beam position
			for _, pos := range beamPositions {
				if pos < len(nextRow) {
					if nextRow[pos] == SPLITTER {
						// Find which splitter this is
						splitterIdx := -1
						for i, splitPos := range splitterPosList {
							if splitPos == pos {
								splitterIdx = i
								break
							}
						}
						if splitterIdx >= 0 {
							// Use bit in combo to decide left (0) or right (1)
							goLeft := (combo>>splitterIdx)&1 == 0
							if goLeft && pos > 0 {
								newBeams[pos-1] = true
							} else if !goLeft && pos < len(nextRow)-1 {
								newBeams[pos+1] = true
							}
						}
					} else if nextRow[pos] == EMPTY || nextRow[pos] == BEAM {
						// Beam continues through empty space or existing beam in all timelines
						newBeams[pos] = true
					}
				}
			}

			// Convert map to sorted slice
			beamSlice := make([]int, 0, len(newBeams))
			for pos := range newBeams {
				beamSlice = append(beamSlice, pos)
			}
			sort.Ints(beamSlice)

			// Recursively explore this path and add to count
			count += dfsCount(tree, row+1, beamSlice, memo)
		}
	} else {
		// No splitter, continue normally
		newBeams := []int{}
		for _, pos := range beamPositions {
			if pos < len(nextRow) {
				if nextRow[pos] == EMPTY || nextRow[pos] == BEAM {
					newBeams = append(newBeams, pos)
				}
			}
		}

		// Recursively explore this path
		count = dfsCount(tree, row+1, newBeams, memo)
	}

	// Cache the result
	memo[memoKey] = count
	return count
}

func buildTree(input []byte) [][]string {
	tree := [][]string{}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "")
		tree = append(tree, parts)
	}
	return tree
}

func printTree(previousRow, tree []string) ([]int, int) {
	beams := []int{}
	count := 0
	for idx, t := range previousRow {
		switch t {
		case START:
			beams = append(beams, idx)
		case BEAM:
			if tree[idx] == SPLITTER {
				count++
				tree[idx-1] = BEAM
				tree[idx+1] = BEAM
			}
			if tree[idx] == EMPTY {
				tree[idx] = BEAM
			}
		}
	}

	return beams, count
}

func hashState(row int, beamPositions []int) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%d:", row))
	for _, pos := range beamPositions {
		builder.WriteString(fmt.Sprintf("%d,", pos))
	}

	h := fnv.New64a()
	h.Write([]byte(builder.String()))
	return fmt.Sprintf("%x", h.Sum64())
}
