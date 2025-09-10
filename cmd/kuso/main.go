package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "kuso",
	Short:   "Text processing utilities for stdin",
	Version: "0.1.0",
}

var trimCmd = &cobra.Command{
	Use:   "trim",
	Short: "Trim whitespace from each line",
	Run:   runTrim,
}

var trim2Cmd = &cobra.Command{
	Use:   "trim2",
	Short: "Trim whitespace from entire input",
	Run:   runTrim2,
}

var splitCmd = &cobra.Command{
	Use:   "split [separator]",
	Short: "Split lines by regex separator",
	Args:  cobra.ExactArgs(1),
	Run:   runSplit,
}

var substCmd = &cobra.Command{
	Use:   "subst [regex] [replacement]",
	Short: "Substitute text using regex",
	Args:  cobra.ExactArgs(2),
	Run:   runSubst,
}

var revCmd = &cobra.Command{
	Use:   "rev",
	Short: "Reverse each line",
	Run:   runRev,
}

var headCmd = &cobra.Command{
	Use:   "head [num]",
	Short: "Take first N characters from each line",
	Args:  cobra.ExactArgs(1),
	Run:   runHead,
}

var tailCmd = &cobra.Command{
	Use:   "tail [num]",
	Short: "Take last N characters from each line",
	Args:  cobra.ExactArgs(1),
	Run:   runTail,
}

func init() {
	rootCmd.AddCommand(trimCmd, trim2Cmd, splitCmd, substCmd, revCmd, headCmd, tailCmd)
}

func readStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func runTrim(cmd *cobra.Command, args []string) {
	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func runTrim2(cmd *cobra.Command, args []string) {
	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	output := strings.TrimSpace(strings.Join(lines, "\n"))
	fmt.Print(output)
}

func runSplit(cmd *cobra.Command, args []string) {
	separator := args[0]
	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	regex, err := regexp.Compile(separator)
	if err != nil {
		fmt.Fprintf(os.Stderr, "regex invalid: %v\n", err)
		os.Exit(1)
	}

	var result []string
	for _, line := range lines {
		parts := regex.Split(line, -1)
		result = append(result, parts...)
	}
	fmt.Println(strings.Join(result, "\n"))
}

func runSubst(cmd *cobra.Command, args []string) {
	pattern := args[0]
	replacement := args[1]
	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "regex invalid: %v\n", err)
		os.Exit(1)
	}

	for i, line := range lines {
		lines[i] = regex.ReplaceAllString(line, replacement)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func runRev(cmd *cobra.Command, args []string) {
	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	for i, line := range lines {
		runes := []rune(line)
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		lines[i] = string(runes)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func runHead(cmd *cobra.Command, args []string) {
	var num int
	if _, err := fmt.Sscanf(args[0], "%d", &num); err != nil {
		fmt.Fprintf(os.Stderr, "invalid number: %s\n", args[0])
		os.Exit(1)
	}

	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	for i, line := range lines {
		runes := []rune(line)
		if len(runes) > num {
			runes = runes[:num]
		}
		lines[i] = string(runes)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func runTail(cmd *cobra.Command, args []string) {
	var num int
	if _, err := fmt.Sscanf(args[0], "%d", &num); err != nil {
		fmt.Fprintf(os.Stderr, "invalid number: %s\n", args[0])
		os.Exit(1)
	}

	lines, err := readStdin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}

	for i, line := range lines {
		runes := []rune(line)
		if len(runes) > num {
			runes = runes[len(runes)-num:]
		}
		lines[i] = string(runes)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}