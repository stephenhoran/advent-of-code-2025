package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Day6 struct{}

func init() {
	Days.RegisterDay(6, &Day6{})
}

func (d *Day6) SolvePart1(input []byte) (string, error) {
	parsedInput, operations := parseInput(input)
	problems := len(operations)
	answers := make([]int, 0)

	for p := 0; p < problems; p++ {
		result := 0
		for i := 0; i < len(parsedInput); i++ {
			if i == 0 {
				result = parsedInput[i][p]
				continue
			}

			result = applyOperations(result, parsedInput[i][p], operations[p])
		}

		answers = append(answers, result)
	}

	final := 0
	for _, a := range answers {
		final += a
	}

	return fmt.Sprintf("%d", final), nil
}

func (d *Day6) SolvePart2(input []byte) (string, error) {
	parsedInput, operations := parseInputWithSpaces(input)
	gridStart := len(operations) - 1
	gridEnd := 0
	oper := ""
	result := 0

	for {
		if gridStart <= 0 {
			break
		}

		gridStart, gridEnd, oper = findProblemGrid(operations, gridStart)
		result += solveProblemGrid(parsedInput, gridStart, gridEnd, oper)
		gridStart = gridEnd - 1
	}

	return fmt.Sprintf("%d", result), nil
}

func parseInput(input []byte) ([][]int, []string) {
	var operations []string
	i := make([][]int, 0)

	lines := strings.Split(string(input), "\n")
	linesLen := len(lines)
	for idx, line := range lines {
		if idx == linesLen-1 {
			operations = strings.FieldsFunc(line, func(c rune) bool {
				return c == ' ' || c == '\t'
			})
			break
		}

		p := make([]int, 0)

		// maybe just using strings.Fields would be better just didnt feel like changing it.
		parts := strings.FieldsFunc(line, func(c rune) bool {
			return c == ' ' || c == '\t'
		})
		for _, part := range parts {
			j, _ := strconv.Atoi(part)
			p = append(p, j)
		}

		i = append(i, p)
	}
	return i, operations

}

func parseInputWithSpaces(input []byte) ([][]string, []string) {
	var operations []string
	i := make([][]string, 0)

	lines := strings.Split(string(input), "\n")
	linesLen := len(lines)
	for idx, line := range lines {
		if idx == linesLen-1 {
			operations = strings.Split(line, "")
			break
		}

		p := make([]string, 0)
		parts := strings.Split(line, "")
		for _, part := range parts {
			p = append(p, part)
		}

		i = append(i, p)
	}
	return i, operations
}

func applyOperations(x int, y int, operation string) int {
	switch operation {
	case "*":
		return x * y
	case "+":
		return x + y
	default:
		return 0
	}
}

func findProblemGrid(input []string, previousEndIndex int) (start int, end int, oper string) {
	for i := previousEndIndex; i <= len(input); i-- {
		if input[i] != " " {
			return previousEndIndex, i, input[i]
		}
	}
	return 0, 0, ""
}

func solveProblemGrid(input [][]string, start int, end int, oper string) int {
	result := 0
	for i := start; i > end-1; i-- {
		numString := ""
		for j := 0; j < len(input); j++ {
			numString += input[j][i]
		}

		//ignore the blank space line between problems
		if strings.TrimSpace(numString) == "" {
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(numString))
		if err != nil {
			panic(err)
		}

		// if this is the first number, set it as the result
		if result == 0 {
			result = num
			continue
		}
		result = applyOperations(result, num, oper)
	}

	return result
}
