package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct{}

func init() {
	Days.RegisterDay(2, &Day2{})
}

func (d *Day2) SolvePart1(input []byte) (string, error) {
	ranges := strings.Split(string(input), ",")
	result := interator(ranges, isInvalidID)

	return fmt.Sprintf("%d", result), nil
}

func (d *Day2) SolvePart2(input []byte) (string, error) {
	ranges := strings.Split(string(input), ",")
	result := interator(ranges, isInvalidIDPart2)

	return fmt.Sprintf("%d", result), nil
}

func interator(ranges []string, f func(int) bool) int {
	invalidIDsCount := 0
	for _, r := range ranges {
		// split the range into two parts and convert to integers
		parts := strings.Split(r, "-")
		part1, _ := strconv.Atoi(parts[0])
		part2, _ := strconv.Atoi(parts[1])

		for i := part1; i <= part2; i++ {
			if f(i) {
				invalidIDsCount = invalidIDsCount + i
			}
		}
	}

	return invalidIDsCount
}

func isInvalidID(id int) bool {
	idString := strconv.Itoa(id)
	idLength := len(idString)

	// if the id is a single digit, it is not invalid
	if idLength == 1 {
		return false
	}

	part1 := idString[:idLength/2]
	part2 := idString[idLength/2:]

	// check for leading zeros
	if part1[0] == 0 || part2[0] == 0 {
		return false
	}

	return part1 == part2
}

// I learned the hard way that Go Regex does not support backreferences, so I have to implement segment matching manually.
// This feels hacky but not sure of a better way to do this.
func isInvalidIDPart2(id int) bool {
	idString := strconv.Itoa(id)
	idLength := len(idString)

	if idLength == 1 {
		return false
	}

	// Try different segment lengths
	for segmentLen := 1; segmentLen <= idLength/2; segmentLen++ {
		// Check if the length is divisible by segment length
		if idLength%segmentLen != 0 {
			continue
		}

		// Get the first segment
		firstSegment := idString[:segmentLen]

		allMatch := true
		numRepetitions := idLength / segmentLen

		// Check if all segments match the first segment
		for i := 1; i < numRepetitions; i++ {
			segment := idString[i*segmentLen : (i+1)*segmentLen]
			if segment != firstSegment {
				allMatch = false
				break
			}
		}

		// If all segments match and we have at least 2 repetitions
		if allMatch && numRepetitions >= 2 {
			return true
		}
	}

	return false
}
