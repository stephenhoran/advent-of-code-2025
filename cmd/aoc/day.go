package aoc

import (
	daypkg "advent-of-code-2025/internal/day"
	"advent-of-code-2025/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	day  int
	part int
	file string
)

var DayCmd = &cobra.Command{
	Use:   "day",
	Short: "Day to solve",
	Long:  `Day to solve`,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate day range
		if day < 1 || day > 25 {
			log.Fatalf("Day must be between 1 and 25, got %d", day)
		}

		// Validate part range
		if part < 1 || part > 2 {
			log.Fatalf("Part must be 1 or 2, got %d", part)
		}

		// Get day implementation
		dayImpl := daypkg.Days.GetDay(day)
		if dayImpl == nil {
			log.Fatalf("Day %d is not implemented yet", day)
		}

		// Read file as bytes (Day interface expects []byte)
		input, err := utils.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		t := time.Now()

		// Call the appropriate part
		var result string
		switch part {
		case 1:
			result, err = dayImpl.SolvePart1(input)
		case 2:
			result, err = dayImpl.SolvePart2(input)
		}

		if err != nil {
			log.Fatalf("Error solving day %d part %d: %v", day, part, err)
		}

		fmt.Printf("Day %d, Part %d: %s (took %s)\n", day, part, result, time.Since(t))
	},
}

func init() {
	DayCmd.Flags().IntVarP(&day, "day", "d", 0, "Day number to solve")
	DayCmd.Flags().IntVarP(&part, "part", "p", 0, "Part number to solve")
	DayCmd.Flags().StringVar(&file, "file", "", "input file to solve")
	DayCmd.MarkFlagRequired("day")
	DayCmd.MarkFlagRequired("part")
	DayCmd.MarkFlagRequired("file")
}
