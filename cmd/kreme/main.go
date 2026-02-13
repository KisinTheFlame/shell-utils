package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kreme [pattern] [replacement]",
	Short: "Rename files in directory using regex patterns",
	Args:  cobra.ExactArgs(2),
	Run:   runKreme,
}

var (
	dryRun bool
	dir    string
)

func init() {
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be renamed without doing it")
	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Directory to process")
}

func runKreme(cmd *cobra.Command, args []string) {
	pattern := args[0]
	replacement := args[1]

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "kreme: regex invalid: %v\n", err)
		os.Exit(1)
	}

	dirPath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "kreme: %s: invalid directory\n", dir)
		os.Exit(1)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "kreme: %s: no such directory\n", dir)
		os.Exit(1)
	}

	for _, entry := range entries {
		oldName := entry.Name()
		newName := regex.ReplaceAllString(oldName, replacement)

		if oldName == newName {
			continue
		}

		if dryRun {
			fmt.Printf("%s ==> %s\n", oldName, newName)
		} else {
			oldPath := filepath.Join(dirPath, oldName)
			newPath := filepath.Join(dirPath, newName)
			if err := os.Rename(oldPath, newPath); err != nil {
				fmt.Fprintf(os.Stderr, "kreme: cannot apply on file %s: %v\n", oldName, err)
			}
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
