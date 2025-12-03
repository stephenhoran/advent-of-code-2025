package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Day3 struct{}

func init() {
	Days.RegisterDay(3, &Day3{})
}

func (d *Day3) SolvePart1(input []byte) (string, error) {
	joltages := strings.Split(string(input), "\n")
	total := 0

	for _, joltage := range joltages {
		largestNumbers := selectLargestNDigits(joltage, 2)

		// going to build a string from the varying size of numbers then convert to int
		// this is probably slower but supports both parts.
		var sb strings.Builder
		for _, num := range largestNumbers {
			sb.WriteString(strconv.Itoa(num))
		}

		number, err := strconv.Atoi(sb.String())
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", err)
		}
		total += number
	}

	return fmt.Sprintf("%d", total), nil
}

func (d *Day3) SolvePart2(input []byte) (string, error) {
	joltages := strings.Split(string(input), "\n")
	total := 0

	for _, joltage := range joltages {
		largestNumbers := selectLargestNDigits(joltage, 12)

		// going to build a string from the varying size of numbers then convert to int
		// this is probably slower but supports both parts.
		var sb strings.Builder
		for _, num := range largestNumbers {
			sb.WriteString(strconv.Itoa(num))
		}

		number, err := strconv.Atoi(sb.String())
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", err)
		}
		total += number
	}

	return fmt.Sprintf("%d", total), nil
}

// selectLargestNDigits selects n digits to form the largest number from the given string
// it returns the digits in the order they were found
// it starts at the given position and searches for the largest digit in the allowed range
// it then selects the digit and moves the start position to the next position
// it repeats this process until n digits are selected
// it returns the digits in the order they were found
func selectLargestNDigits(j string, n int) []int {
	result := make([]int, n)
	start := 0

	for pos := 0; pos < n; pos++ {
		// Calculate how many digits we still need after this position
		remainingNeeded := n - pos - 1
		// We can only look up to len(j) - remainingNeeded to ensure we have enough digits left
		maxSearchEnd := len(j) - remainingNeeded

		// Find the largest digit in our new range
		maxDigit := -1
		maxIdx := start
		for i := start; i < maxSearchEnd; i++ {
			digit := int(j[i] - '0')
			if digit > maxDigit {
				maxDigit = digit
				maxIdx = i
			}
		}

		result[pos] = maxDigit
		start = maxIdx + 1 // Next search starts after the selected digit
	}

	return result
}
