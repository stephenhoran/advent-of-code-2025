package day

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day5 struct{}

func init() {
	Days.RegisterDay(5, &Day5{})
}

func (d *Day5) SolvePart1(input []byte) (string, error) {
	parts := strings.Split(string(input), "\n\n")
	fir := freshIngredients(strings.Split(parts[0], "\n"))

	c := 0
	for _, ingredient := range strings.Split(parts[1], "\n") {
		i, _ := strconv.Atoi(ingredient)
		if isValidIngredient(i, fir) {
			c++
		}
	}

	return fmt.Sprintf("%+v\n", c), nil
}

func (d *Day5) SolvePart2(input []byte) (string, error) {
	count := 0

	parts := strings.Split(string(input), "\n\n")
	c := ingredientRange(freshIngredients(strings.Split(parts[0], "\n")))

	//recursively merge until no more merges are needed
	for {
		merged := ingredientRange(c)
		if len(merged) == len(c) {
			break
		}
		c = merged
	}

	for _, m := range c {
		count += m[1] - m[0] + 1
	}

	return fmt.Sprintf("%+v\n", count), nil
}

func freshIngredients(r []string) [][]int {
	ingredients := make([][]int, 0)
	for _, line := range r {
		parts := strings.Split(line, "-")
		f, _ := strconv.Atoi(parts[0])
		l, _ := strconv.Atoi(parts[1])
		ingredients = append(ingredients, []int{f, l})
	}

	return ingredients
}

func isValidIngredient(i int, f [][]int) bool {
	for _, ingredientRange := range f {
		if i >= ingredientRange[0] && i <= ingredientRange[1] {
			return true
		}
	}
	return false
}

func ingredientRange(f [][]int) [][]int {
	i := make([][]int, 0)

	for fidx, ingredientRange := range f {
		if fidx == 0 {
			i = append(i, ingredientRange)
			continue
		}

		compacted := false
		for iidx, ingredient := range i {
			if ingredient[0] <= ingredientRange[1] && ingredientRange[0] <= ingredient[1] {
				i[iidx] = compactRange(ingredient[0], ingredient[1], ingredientRange[0], ingredientRange[1])
				compacted = true
				break
			}
		}
		if !compacted {
			i = append(i, ingredientRange)
		}
	}

	return i
}

func compactRange(i1 int, i2 int, j1 int, j2 int) []int {
	return []int{int(math.Min(float64(i1), float64(j1))), int(math.Max(float64(i2), float64(j2)))}
}
