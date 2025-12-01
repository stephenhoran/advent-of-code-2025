package day

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day1 struct{}

func init() {
	Days.RegisterDay(1, &Day1{})
}

func (d *Day1) SolvePart1(input []byte) (string, error) {
	dialStart := 50
	zeros := 0

	for line := range strings.SplitSeq(string(input), "\n") {
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("invalid distance: %v", err)
		}

		// ignore safe dial roll overs
		if distance > 100 {
			distance = distance % 100
		}

		switch direction {
		case 'L':
			dialStart -= distance
		case 'R':
			dialStart += distance
		}

		if dialStart < 0 {
			dialStart += 100
		}

		if dialStart >= 100 {
			dialStart -= 100
		}

		if dialStart == 0 {
			zeros++
		}
	}

	return fmt.Sprintf("%d", zeros), nil
}

func (d *Day1) SolvePart2(input []byte) (string, error) {
	dialStart := 50
	prevDial := 50
	zeros := 0

	for _, line := range strings.Split(string(input), "\n") {
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("invalid distance: %v", err)
		}

		if distance > 100 {
			c := int(math.Floor(float64(distance) / 100))
			zeros += c
		}

		// ignore safe dial roll overs
		if distance > 100 {
			distance = distance % 100
		}

		switch direction {
		case 'L':
			dialStart -= distance
		case 'R':
			dialStart += distance
		}

		if dialStart < 0 {
			dialStart += 100

			if prevDial != 0 && dialStart != 0 {
				zeros++
			}
		}

		if dialStart > 99 {
			dialStart -= 100

			if prevDial != 0 && dialStart != 0 {
				zeros++
			}
		}

		if dialStart == 0 {
			zeros++
		}

		prevDial = dialStart
	}

	return fmt.Sprintf("%d", zeros), nil
}
