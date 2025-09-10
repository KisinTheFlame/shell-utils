package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func execute() error {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !isAllWhitespace(line) {
			lines = append(lines, line)
		}
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read stdin: %v", err)
	}

	days := len(lines)
	if days == 0 {
		return fmt.Errorf("no valid input lines")
	}

	var timePairs [][2]time.Time
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return fmt.Errorf("time format not correct")
		}

		startTime, err := time.Parse("15:04", parts[0])
		if err != nil {
			return fmt.Errorf("time format not correct")
		}

		endTime, err := time.Parse("15:04", parts[1])
		if err != nil {
			return fmt.Errorf("time format not correct")
		}

		timePairs = append(timePairs, [2]time.Time{startTime, endTime})
	}

	var totalMinutes int64
	for _, pair := range timePairs {
		startTime := pair[0]
		endTime := pair[1]
		diff := endTime.Sub(startTime)
		totalMinutes += int64(diff.Minutes())
	}

	average := float64(totalMinutes) / float64(days) / 60.0
	expected := int64(days) * 12 * 60
	diff := totalMinutes - expected

	fmt.Printf("days: %d\n", days)
	fmt.Printf("average: %.4f hrs\n", average)
	fmt.Printf("diff: %+d mins\n", diff)

	return nil
}

func isAllWhitespace(s string) bool {
	for _, char := range s {
		if !isWhitespace(char) {
			return false
		}
	}
	return true
}

func isWhitespace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}